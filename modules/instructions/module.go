package instructions

import (
	"sync"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/modules/pruning"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/pool"
)

var _ modules.Module = &Module{}
var _ pruning.PruningService = &Module{}

// Module represents the module allowing to store instructions properly inside a dedicated table
type Module struct {
	db     db.InstructionDb
	buffer chan dbtypes.InstructionRow
	pool   pool.Pool

	mtx sync.Mutex
}

func NewModule(db db.InstructionDb, pool pool.Pool) *Module {
	return &Module{
		db:     db,
		buffer: make(chan dbtypes.InstructionRow),
		pool:   pool,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "instructions"
}

func (m *Module) HandleBlock(block types.Block) error {
	return m.createPartition(block.Slot)
}

// HandleInstruction implements modules.InstructionModule
func (m *Module) HandleInstruction(instruction types.Instruction, tx types.Tx) error {
	m.buffer <- dbtypes.NewInstructionRowFromInstruction(instruction)
	return nil
}

// createPartition creates a new partition for the msgs module
func (m *Module) createPartition(slot uint64) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	err := m.db.CreateInstructionsPartition(int(slot / 1000))
	if err != nil {
		return err
	}
	return nil
}

// Prune implements pruning.PruningService
func (m *Module) Prune(slot uint64) error {
	return m.db.PruneInstructionsBeforeSlot(slot)
}
