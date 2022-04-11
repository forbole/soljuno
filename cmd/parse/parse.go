package parse

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	cmdtypes "github.com/forbole/soljuno/cmd/types"

	"github.com/forbole/soljuno/types/logging"

	"github.com/go-co-op/gocron"

	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/solana/parser/manager"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/worker"

	"github.com/spf13/cobra"
)

var (
	waitGroup sync.WaitGroup
)

// ParseCmd returns the command that should be run when we want to start parsing a chain state.
func ParseCmd(cmdCfg *cmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "parse",
		Short:   "Start parsing the blockchain data",
		PreRunE: cmdtypes.ReadConfig(cmdCfg),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := GetParsingContext(cmdCfg)
			if err != nil {
				return err
			}
			return StartParsing(context)
		},
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

	// Create and register solana instruction parserManager
	parserManager := manager.NewDefaultManager()

	workerCtx := worker.NewContext(ctx.Proxy, ctx.Database, parserManager, ctx.Logger, ctx.Pool, exportQueue, ctx.Modules)
	workers := make([]worker.Worker, cfg.GetWorkers())
	workerStopChs := make([]chan bool, cfg.GetWorkers())
	for i := range workers {
		stopCh := make(chan bool, 1)
		workers[i] = worker.NewWorker(i, workerCtx).WithStopChannel(stopCh)
		workerStopChs[i] = stopCh
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
	trapSignal(ctx, exportQueue, workerStopChs)

	latestSlot, err := ctx.Proxy.GetLatestSlot()
	if err != nil {
		panic(fmt.Errorf("failed to get last block from RPC client: %s", err))
	}

	var oldBlockListenerWg sync.WaitGroup
	if cfg.ShouldParseOldBlocks() {
		cfg := types.Cfg.GetParsingConfig()
		oldBlockListenerWg.Add(1)

		go func() {
			enqueueMissingSlots(ctx, exportQueue, cfg.GetStartSlot(), latestSlot)
			oldBlockListenerWg.Done()
		}()
	}

	if cfg.ShouldParseNewBlocks() {
		go func() {
			oldBlockListenerWg.Wait()
			startNewBlockListener(ctx, exportQueue, latestSlot+1)
		}()
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
		slots, err := ctx.Proxy.GetBlocks(i, next)
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
		end, err := ctx.Proxy.GetLatestSlot()
		if err != nil {
			continue
		}
		if end > start {
			enqueueMissingSlots(ctx, exportQueue, start, end)
			start = end + 1
		}
		time.Sleep(time.Second)
	}
}

// trapSignal will listen for any OS signal and invoke Done on the main
// WaitGroup allowing the main process to gracefully exit.
func trapSignal(ctx *Context, queue types.SlotQueue, workerStopChs []chan bool) {
	var sigCh = make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		defer ctx.Database.Close()
		defer waitGroup.Done()
		defer ctx.Logger.Info("stopped the program...")

		sig := <-sigCh
		ctx.Logger.Info("caught signal; shutting down...", "signal", sig.String())

		ctx.Logger.Info("closing workers...")
		// close workers
		for _, ch := range workerStopChs {
			ch <- true
		}

		// wait stopped signal from workers
		for i := 0; i < len(workerStopChs); i++ {
			<-workerStopChs[i]
		}

		// wait if the pool is not empty
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			for !ctx.Pool.IsEmpty() {
				time.Sleep(time.Second)
			}
			wg.Done()
		}()
		wg.Wait()

		next := <-queue
		err := updateStartSlot(ctx.GlobalCfg, next)
		if err != nil {
			ctx.Logger.Info("failed to update start slot")
		}

	}()
}

func updateStartSlot(cfg types.Config, slot uint64) error {
	parsingCfg := cfg.GetParsingConfig()
	parsingCfg.SetStartSlot(slot)
	cfg.SetParsingConfig(parsingCfg)
	return types.Write(cfg, types.GetConfigFilePath())
}
