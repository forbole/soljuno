package txs

import (
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
	err := m.db.CreateTxPartition(int(block.Slot / 1000))
	if err != nil {
		return err
	}
	m.buffer <- block
	return nil
}

// Prune implements pruning.PruningService
func (m *Module) Prune(slot uint64) error {
	for {
		partitionName, err := m.db.GetOldestTxPartitionNameBeforeSlot(slot)
		if err != nil {
			return err
		}
		if partitionName == "" {
			return nil
		}

		err = m.db.DropTxPartition(partitionName)
		if err != nil {
			return err
		}
	}
}
