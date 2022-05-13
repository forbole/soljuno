package worker

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/forbole/soljuno/types/logging"
	"github.com/forbole/soljuno/types/pool"

	"github.com/forbole/soljuno/modules"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/parser/manager"
	"github.com/forbole/soljuno/types"
)

// Worker defines a job consumer that is responsible for getting and
// aggregating block and associated data and exporting it to a database.
type Worker struct {
	queue         types.SlotQueue
	cp            client.ClientProxy
	db            db.Database
	parserManager manager.ParserManager
	logger        logging.Logger

	pool    pool.Pool
	index   int
	modules []modules.Module

	stopChannel chan bool
}

// NewWorker allows to create a new Worker implementation.
func NewWorker(index int, ctx *Context) Worker {
	return Worker{
		index:         index,
		cp:            ctx.ClientProxy,
		queue:         ctx.Queue,
		db:            ctx.Database,
		parserManager: ctx.ParserManager,
		modules:       ctx.Modules,
		logger:        ctx.Logger,
		pool:          ctx.Pool,
	}
}

func (w Worker) WithStopChannel(ch chan bool) Worker {
	w.stopChannel = ch
	return w
}

// Start starts a worker by listening for new jobs (block heights) from the
// given worker queue. Any failed job is logged and re-enqueued.
func (w Worker) Start() {
	logging.WorkerCount.Inc()

	var stopSignal bool
	go func() {
		stopSignal = <-w.stopChannel
	}()

	for {
		// check the stop signal
		if stopSignal {
			w.logger.Debug("closed worker", "number", w.index)
			w.stopChannel <- true
			return
		}

		// wait if the pool is full
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			for !w.pool.IsFree() {
				time.Sleep(time.Second)
			}
			wg.Done()
		}()
		wg.Wait()

		// work on the block
		timeout := time.After(5 * time.Second)
		select {
		case slot := <-w.queue:
			w.work(slot)
		case <-timeout:
			continue
		}
	}
}

func (w Worker) work(slot uint64) {
	start := time.Now()
	if err := w.process(slot); err != nil {
		// re-enqueue any failed job
		// TODO: Implement exponential backoff or max retries for a block slot.
		w.logger.Error("re-enqueueing failed block", "slot", slot, "err", err)
		w.queue <- slot
	}
	logging.WorkerSlot.WithLabelValues(fmt.Sprintf("%d", w.index)).Set(float64(slot))
	w.logger.Debug("processed block time", "slot", slot, "seconds", time.Since(start).Seconds())
}

// process defines the job consumer workflow. It will fetch a block for a given
// slot and associated metadata and export it to a database. It returns an
// error if any export process fails.
func (w Worker) process(slot uint64) error {
	var err error
	defer func() {
		r := recover()
		if r != nil {
			panic(fmt.Errorf("panic on slot %v with error: %v", slot, r))
		}
	}()

	w.logger.Info("processing block", "slot", slot)
	b, err := w.cp.GetBlock(slot)
	if err != nil {
		return fmt.Errorf("failed to get block from rpc server: %s", err)
	}
	block := types.NewBlockFromResult(w.parserManager, slot, b)

	// set block leader
	leaders, err := w.cp.GetSlotLeaders(slot, 1)
	if err != nil && !strings.Contains(err.Error(), "-32602:Invalid slot range: leader schedule for epoch") {
		return err
	}
	if len(leaders) == 1 {
		block.Leader = leaders[0]
	} else {
		block.Leader = ""
	}
	fmt.Println(block)

	err = w.ExportBlock(block)
	return err
}

// ExportBlock accepts a finalized block and a corresponding set of transactions
// and persists them to the database along with attributable metadata. An error
// is returned if the write fails.
func (w Worker) ExportBlock(block types.Block) error {
	return w.handleBlock(block)
}

// handleBlock handles all the events in a block
func (w Worker) handleBlock(block types.Block) error {
	if err := w.handleBlockModules(block); err != nil {
		return err
	}

	// handle txs asynchronously
	errChs := make([]chan error, len(block.Txs))
	for i, tx := range block.Txs {
		tx := tx
		errCh, err := w.pool.DoAsync(func() error { return w.handleTx(tx) })
		if err != nil {
			return err
		}
		errChs[i] = errCh
	}

	// check errors
	for _, errCh := range errChs {
		err := <-errCh
		if err != nil {
			return err
		}
	}
	return nil
}

// handleBlockModules handles the block with modules
func (w Worker) handleBlockModules(block types.Block) error {
	for _, module := range w.modules {
		if blockModule, ok := module.(modules.BlockModule); ok {
			err := blockModule.HandleBlock(block)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// handleTx handles all the events in a transaction
func (w Worker) handleTx(tx types.Tx) error {
	for _, module := range w.modules {
		if transactionModule, ok := module.(modules.TransactionModule); ok {
			err := transactionModule.HandleTx(tx)
			if err != nil {
				return err
			}
		}
	}
	return w.handleInstructions(tx)
}

// handleInstructions handles all the instructions events in a transaction
func (w Worker) handleInstructions(tx types.Tx) error {
	for _, instruction := range tx.Instructions {
		instruction := instruction
		err := w.handleInstruction(tx, instruction)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w Worker) handleInstruction(tx types.Tx, instruction types.Instruction) error {
	for _, module := range w.modules {
		if InstructionModule, ok := module.(modules.InstructionModule); ok {
			err := InstructionModule.HandleInstruction(instruction, tx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
