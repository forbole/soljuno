package parse

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/panjf2000/ants/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/forbole/soljuno/solana/parser"
	"github.com/forbole/soljuno/solana/program/bpfloader"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/solana/program/token"
	"github.com/forbole/soljuno/solana/program/vote"

	"github.com/forbole/soljuno/types/logging"

	"github.com/go-co-op/gocron"

	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/worker"

	"github.com/spf13/cobra"
)

var (
	waitGroup sync.WaitGroup
)

// ParseCmd returns the command that should be run when we want to start parsing a chain state.
func ParseCmd(cmdCfg *Config) *cobra.Command {
	return &cobra.Command{
		Use:     "parse",
		Short:   "Start parsing the blockchain data",
		PreRunE: ReadConfig(cmdCfg),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := GetParsingContext(cmdCfg)
			if err != nil {
				return err
			}

			go StartPrometheus()

			return StartParsing(context)
		},
	}
}

// StartPrometheus allows to start a Telemetry server used to expose useful metrics
func StartPrometheus() {
	cfg := types.Cfg.GetTelemetryConfig()
	if !cfg.IsEnabled() {
		return
	}

	router := mux.NewRouter()
	router.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.GetPort()), router)
	if err != nil {
		panic(err)
	}
}

// StartParsing represents the function that should be called when the parse command is executed
func StartParsing(ctx *Context) error {
	// Get the config
	cfg := types.Cfg.GetParsingConfig()
	logging.StartSlot.Add(float64(cfg.GetStartSlot()))

	// Start periodic operations
	scheduler := gocron.NewScheduler(time.UTC)
	for _, module := range ctx.Modules {
		if module, ok := module.(modules.PeriodicOperationsModule); ok {
			err := module.RegisterPeriodicOperations(scheduler)
			if err != nil {
				return err
			}
		}
	}
	scheduler.StartAsync()

	// Create a queue that will collect, aggregate, and export blocks and metadata
	exportQueue := types.NewQueue(25)

	// Create and register solana message parser
	parser := parser.NewParser()
	parser.Register(vote.ProgramID, vote.Parser{})
	parser.Register(stake.ProgramID, stake.Parser{})
	parser.Register(system.ProgramID, system.Parser{})
	parser.Register(token.ProgramID, token.Parser{})
	parser.Register(bpfloader.ProgramID, bpfloader.Parser{})
	parser.Register(upgradableLoader.ProgramID, upgradableLoader.Parser{})

	workerCfg := types.Cfg.GetWorkerConfig()
	// Create workers
	pool, err := ants.NewPool(workerCfg.GetPoolSize())
	if err != nil {
		return err
	}

	bankTasks := types.NewBankTaskQueue(25)
	workerCtx := worker.NewContext(ctx.Proxy, ctx.Database, parser, ctx.Logger, pool, exportQueue, ctx.Modules, bankTasks)
	workers := make([]worker.Worker, cfg.GetWorkers())
	for i := range workers {
		workers[i] = worker.NewWorker(i, workerCtx)
	}

	waitGroup.Add(1)

	// Hanle bank module tasks synchronously to prevent dead locks
	go func() {
		for {
			task := <-bankTasks
			task()
		}
	}()

	// Run all the async operations
	for _, module := range ctx.Modules {
		if module, ok := module.(modules.AsyncOperationsModule); ok {
			go module.RunAsyncOperations()
		}
	}

	// Start each blocking worker in a go-routine where the worker consumes jobs
	// off of the export queue.
	for i, w := range workers {
		ctx.Logger.Debug("starting worker...", "number", i+1)
		go w.Start()
	}

	// Listen for and trap any OS signal to gracefully shutdown and exit
	trapSignal(ctx)

	if cfg.ShouldParseGenesis() {
		// TODO add genesis block
	}

	latestSlot, err := ctx.Proxy.LatestSlot()
	if err != nil {
		panic(fmt.Errorf("failed to get last block from RPC client: %s", err))
	}
	if cfg.ShouldParseOldBlocks() {
		cfg := types.Cfg.GetParsingConfig()
		go enqueueMissingSlots(ctx, exportQueue, cfg.GetStartSlot(), latestSlot)
	}

	if cfg.ShouldParseNewBlocks() {
		go startNewBlockListener(ctx, exportQueue, latestSlot)
	}

	// Block main process (signal capture will call WaitGroup's Done)
	waitGroup.Wait()
	return nil
}

// enqueueMissingSlots enqueue jobs (block slots) for missed blocks starting
// at the startSlot up until the latest known slot.
func enqueueMissingSlots(ctx *Context, exportQueue types.SlotQueue, start uint64, end uint64) {
	ctx.Logger.Info("syncing missing blocks...", "latest_block_slot", end)
	for i := start; i < end; {
		next := i + 25
		if next > end {
			next = end
		}
		slots, err := ctx.Proxy.Slots(i, next)
		if err != nil {
			continue
		}
		for _, slot := range slots {
			ctx.Logger.Debug("enqueueing missing block", "slot", slot)
			exportQueue <- slot
		}
		i = next + 1
	}
}

// startNewBlockListener subscribes to new block events via the RPC
// and enqueues each new block slot onto the provided queue. It blocks as new
// blocks are incoming.
func startNewBlockListener(ctx *Context, exportQueue types.SlotQueue, start uint64) {
	for {
		end, err := ctx.Proxy.LatestSlot()
		if err != nil {
			continue
		}
		enqueueMissingSlots(ctx, exportQueue, start, end)
		start = end + 1
		time.Sleep(time.Second)
	}
}

// trapSignal will listen for any OS signal and invoke Done on the main
// WaitGroup allowing the main process to gracefully exit.
func trapSignal(ctx *Context) {
	var sigCh = make(chan os.Signal)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		sig := <-sigCh
		ctx.Logger.Info("caught signal; shutting down...", "signal", sig.String())
		defer ctx.Database.Close()
		defer waitGroup.Done()
	}()
}
