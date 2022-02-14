package worker

import (
	"fmt"
	"time"

	"github.com/forbole/soljuno/types/logging"
	"github.com/forbole/soljuno/types/pool"

	"github.com/forbole/soljuno/modules"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/parser/manager"
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

// Start starts a worker by listening for new jobs (block heights) from the
// given worker queue. Any failed job is logged and re-enqueued.
func (w Worker) Start() {
	logging.WorkerCount.Inc()
	for i := range w.queue {
		start := time.Now()
		if err := w.process(i); err != nil {
			// re-enqueue any failed job
			// TODO: Implement exponential backoff or max retries for a block slot.
			w.logger.Error("re-enqueueing failed block", "slot", i, "err", err)
			w.queue <- i
		}
		logging.WorkerSlot.WithLabelValues(fmt.Sprintf("%d", w.index)).Set(float64(i))
		w.logger.Debug("processed block time", "slot", i, "seconds", time.Since(start).Seconds())
		wait := make(chan bool)
		go func() {
			for !w.pool.IsFree() {
				time.Sleep(time.Second)
			}
			wait <- true
		}()
		<-wait
	}
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
	exists, err := w.db.HasBlock(slot)
	if err != nil {
		return fmt.Errorf("error while searching for block: %s", err)
	}

	if exists {
		w.logger.Debug("skipping already exported block", "slot", slot)
		return nil
	}

	w.logger.Info("processing block", "slot", slot)
	b, err := w.cp.GetBlock(slot)
	if err != nil {
		return fmt.Errorf("failed to get block from rpc server: %s", err)
	}
	block := types.NewBlockFromResult(w.parserManager, slot, b)

	// set block proposer
	proposers, err := w.cp.GetSlotLeaders(slot, 1)
	if err != nil {
		return err
	}
	block.Proposer = proposers[0]

	err = w.ExportBlock(block)
	return err
}

// ExportBlock accepts a finalized block and a corresponding set of transactions
// and persists them to the database along with attributable metadata. An error
// is returned if the write fails.
func (w Worker) ExportBlock(block types.Block) error {
	// Save the block
	err := w.db.SaveBlock(block)
	if err != nil {
		return fmt.Errorf("failed to persist block: %s", err)
	}
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
	return w.handleMessages(tx)
}

// handleMessages handles all the messages events in a transaction
func (w Worker) handleMessages(tx types.Tx) error {
	for _, msg := range tx.Messages {
		msg := msg
		err := w.handleMessage(tx, msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w Worker) handleMessage(tx types.Tx, msg types.Message) error {
	for _, module := range w.modules {
		if messageModule, ok := module.(modules.MessageModule); ok {
			err := messageModule.HandleMsg(msg, tx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
