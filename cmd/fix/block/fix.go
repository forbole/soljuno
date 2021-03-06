package block

import (
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	cmdtypes "github.com/forbole/soljuno/cmd/types"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/modules/fix"
	"github.com/forbole/soljuno/solana/program/parser/manager"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/worker"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
)

var (
	waitGroup sync.WaitGroup
)

func FixMissingBlocksCmd(cmdCfg *cmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "missing-blocks [start] [end]",
		Short:   "Fix missing blocks from specific start slot",
		PreRunE: cmdtypes.ReadConfig(cmdCfg),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := GetFixContext(cmdCfg)
			if err != nil {
				return err
			}

			start, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			end, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			return StartFixing(ctx, start, end)
		},
	}
	return cmd
}

func StartFixing(ctx *Context, start uint64, end uint64) error {
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

	// Create and register solana instruction parserManager
	parserManager := manager.NewDefaultManager()

	cfg := types.Cfg.GetParsingConfig()
	workerCtx := worker.NewContext(ctx.Proxy, ctx.Database, parserManager, ctx.Logger, ctx.Pool, ctx.SlotQueue, ctx.Modules)
	workers := make([]worker.Worker, cfg.GetWorkers())
	workerStopChs := make([]chan bool, cfg.GetWorkers())
	for i := range workers {
		stopCh := make(chan bool, 1)
		workers[i] = worker.NewWorker(i, workerCtx).WithStopChannel(stopCh)
		workerStopChs[i] = stopCh
	}

	// Start each blocking worker in a go-routine where the worker consumes jobs
	// off of the export queue.
	for i, w := range workers {
		ctx.Logger.Debug("starting worker...", "number", i+1)
		go w.Start()
	}

	waitGroup.Add(1)

	// Run all the async operations
	for _, module := range ctx.Modules {
		if module, ok := module.(modules.AsyncOperationsModule); ok {
			go module.RunAsyncOperations()
		}
	}

	// Listen for and trap any OS signal to gracefully shutdown and exit
	trapSignal(ctx, workerStopChs)

	go enqueueMissingSlots(ctx, start, end)

	// Block main process (signal capture will call WaitGroup's Done)
	waitGroup.Wait()
	return nil
}

func enqueueMissingSlots(ctx *Context, start uint64, end uint64) {
	ctx.Logger.Info("fixing missing blocks...", "latest_block_slot", end)
	for i := start; i < end; {
		next := i + 1000
		if next >= end {
			next = end - 1
		}
		fix.EnqueueMissingSlots(ctx.Database, ctx.SlotQueue, ctx.Proxy, i, next)
		i = next + 1
	}
	ctx.Logger.Info("enqueueing missing block finished")

}

// trapSignal will listen for any OS signal and invoke Done on the main
// WaitGroup allowing the main process to gracefully exit.
func trapSignal(ctx *Context, workerStopChs []chan bool) {
	var sigCh = make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		sig := <-sigCh
		ctx.Logger.Info("caught signal; shutting down...", "signal", sig.String())
		close(ctx, workerStopChs)
	}()
}

// close stops the program properly
func close(ctx *Context, workerStopChs []chan bool) {
	defer ctx.Database.Close()
	defer waitGroup.Done()
	defer ctx.Logger.Info("stopped the program...")

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

	for _, module := range ctx.Modules {
		if periodicModule, ok := module.(modules.PeriodicOperationsModule); ok {
			if err := periodicModule.RunPeriodicOperations(); err != nil {
				ctx.Logger.Error("module", module.Name(), "err", err)
			}
		}
	}
}
