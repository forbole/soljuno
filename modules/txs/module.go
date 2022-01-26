package txs

import (
	"sync"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/modules/pruning"
	"github.com/forbole/soljuno/types"
	"github.com/panjf2000/ants/v2"
)

var _ modules.Module = &Module{}
var _ pruning.PruningService = &Module{}

type Module struct {
	db     db.TxDb
	buffer chan types.Block
	pool   *ants.Pool

	mtx sync.Mutex
}

func NewModule(db db.TxDb, pool *ants.Pool) *Module {
	return &Module{
		db:     db,
		buffer: make(chan types.Block),
		pool:   pool,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "txs"
}

// HandleBlock implements modules.MessageModule
func (m *Module) HandleBlock(block types.Block) error {
	err := m.createPartition(block.Slot)
	if err != nil {
		return err
	}
	m.buffer <- block
	return nil
}

// createPartition creates a new partition for the txs module
func (m *Module) createPartition(slot uint64) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	err := m.db.CreateTxPartition(int(slot / 1000))
	if err != nil {
		return err
	}
	return nil
}

// Prune implements pruning.PruningService
func (m *Module) Prune(slot uint64) error {
	return m.db.PruneTxsBeforeSlot(slot)
}
