package worker

import (
	"fmt"
	"time"

	"github.com/forbole/soljuno/solana/parser"
	"github.com/forbole/soljuno/types/logging"
	"github.com/panjf2000/ants/v2"

	"github.com/forbole/soljuno/modules"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// Worker defines a job consumer that is responsible for getting and
// aggregating block and associated data and exporting it to a database.
type Worker struct {
	queue  types.SlotQueue
	cp     client.Proxy
	db     db.Database
	parser parser.Parser
	logger logging.Logger

	pool    *ants.Pool
	index   int
	modules []modules.Module
}

// NewWorker allows to create a new Worker implementation.
func NewWorker(index int, ctx *Context) Worker {
	return Worker{
		index:   index,
		cp:      ctx.ClientProxy,
		queue:   ctx.Queue,
		db:      ctx.Database,
		parser:  ctx.Parser,
		modules: ctx.Modules,
		logger:  ctx.Logger,
		pool:    ctx.Pool,
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
			w.pool.Submit(func() {
				w.logger.Error("re-enqueueing failed block", "slot", i, "err", err)
				w.queue <- i
			})
		}
		logging.WorkerSlot.WithLabelValues(fmt.Sprintf("%d", w.index)).Set(float64(i))
		w.logger.Debug("process block time", "seconds", time.Since(start).Seconds())
	}
}

// process defines the job consumer workflow. It will fetch a block for a given
// slot and associated metadata and export it to a database. It returns an
// error if any export process fails.
func (w Worker) process(slot uint64) error {
	exists, err := w.db.HasBlock(slot)
	if err != nil {
		return fmt.Errorf("error while searching for block: %s", err)
	}

	if exists {
		w.logger.Debug("skipping already exported block", "slot", slot)
		return nil
	}

	w.logger.Info("processing block", "slot", slot)

	b, err := w.cp.Block(slot)
	if err != nil {
		return fmt.Errorf("failed to get block from rpc server: %s", err)
	}

	block := types.NewBlockFromResult(w.parser, slot, b)

	return w.ExportBlock(block)
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

	w.handleBlockModules(block)

	// Handle modules
	w.pool.Submit(func() {
		w.HandleModules(block)
	})
	return nil
}

// ExportTxs accepts a slice of transactions and persists then inside the database.
// An error is returned if the write fails.
func (w Worker) HandleModules(block types.Block) {
	// Handle all the transactions inside the block
	w.handleBlockModules(block)
	for _, tx := range block.Txs {
		w.handleTx(tx)
	}
	w.logger.Info("block indexed", "slot", block.Slot)
}

// handleBlockModules handles the block with modules
func (w Worker) handleBlockModules(block types.Block) {
	for _, module := range w.modules {
		if blockModule, ok := module.(modules.BlockModule); ok {
			err := blockModule.HandleBlock(block)
			if err != nil {
				w.logger.BlockError(module, block, err)
			}
		}
	}
}

// handleTx handles all the the element contained inside the transaction
func (w Worker) handleTx(tx types.Tx) {
	for _, module := range w.modules {
		if transactionModule, ok := module.(modules.TransactionModule); ok {
			err := transactionModule.HandleTx(tx)
			if err != nil {
				w.logger.TxError(module, tx, err)
			}
		}
	}
	w.handleMessageModules(tx)
}

// handleMessageModules handles all the messages contained inside the transaction with modules
func (w Worker) handleMessageModules(tx types.Tx) {
	for _, msg := range tx.Messages {
		for _, module := range w.modules {
			if messageModule, ok := module.(modules.MessageModule); ok {
				err := messageModule.HandleMsg(msg, tx)
				if err != nil {
					w.logger.MsgError(module, tx, msg, err)
				}
			}
		}
	}
}
