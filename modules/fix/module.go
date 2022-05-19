package fix

import (
	"sync"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

type ClientProxy interface {
	GetBlocks(start uint64, end uint64) ([]uint64, error)
}

var (
	_ modules.Module      = &Module{}
	_ modules.BlockModule = &Module{}
)

type Module struct {
	db           db.FixMissingBlockDb
	SlotInterval uint64
	slotQueue    types.SlotQueue
	client       ClientProxy
	mtx          sync.Mutex
}

func NewModule(db db.FixMissingBlockDb, slotQueue types.SlotQueue, client ClientProxy) *Module {
	return &Module{
		db:        db,
		slotQueue: slotQueue,
		client:    client,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "fix"
}

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(block types.Block) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	// set 50 delayed slot
	interval := (block.Slot - 50) / 100
	if interval <= m.SlotInterval {
		return nil
	}
	// the first time to set the current interval
	if m.SlotInterval == 0 {
		m.SlotInterval = interval
		return nil
	}
	m.SlotInterval = interval
	return HandleBlock(block, m.db, m.slotQueue, m.client)
}
