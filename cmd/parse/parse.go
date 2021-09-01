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
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/forbole/soljuno/solana/parser"
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
	parser.Register(vote.ProgramID, vote.VoteParser{})

	// Create workers
	workerCtx := worker.NewContext(ctx.Proxy, ctx.Database, parser, ctx.Logger, exportQueue, ctx.Modules)
	workers := make([]worker.Worker, cfg.GetWorkers(), cfg.GetWorkers())
	for i := range workers {
		workers[i] = worker.NewWorker(i, workerCtx)
	}

	waitGroup.Add(1)

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
		// Add the genesis to the queue if requested
		exportQueue <- 0
	}

	if cfg.ShouldParseOldBlocks() {
		go enqueueMissingSlots(exportQueue, ctx)
	}

	if cfg.ShouldParseNewBlocks() {
		go startNewBlockListener(exportQueue, ctx)
	}

	// Block main process (signal capture will call WaitGroup's Done)
	waitGroup.Wait()
	return nil
}

// enqueueMissingSlots enqueue jobs (block slots) for missed blocks starting
// at the startSlot up until the latest known slot.
func enqueueMissingSlots(exportQueue types.SlotQueue, ctx *Context) {
	// Get the config
	cfg := types.Cfg.GetParsingConfig()

	// Get the latest slot
	latestBlockSlot, err := ctx.Proxy.LatestSlot()
	if err != nil {
		panic(fmt.Errorf("failed to get last block from RPC client: %s", err))
	}

	// TODO solana fastsync module

	ctx.Logger.Info("syncing missing blocks...", "latest_block_slot", latestBlockSlot)
	for i := cfg.GetStartSlot(); i <= latestBlockSlot; i++ {
		ctx.Logger.Debug("enqueueing missing block", "slot", i)
		exportQueue <- i
	}
}

// TODO rebuild for json rpc websocket to subscribe
// startNewBlockListener subscribes to new block events via the Tendermint RPC
// and enqueues each new block height onto the provided queue. It blocks as new
// blocks are incoming.
func startNewBlockListener(exportQueue types.SlotQueue, ctx *Context) {

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
