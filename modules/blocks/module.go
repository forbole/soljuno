package blocks

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

var (
	_ modules.Module = &Module{}
)

type Module struct {
	db db.BlockDb
}

func NewModule(db db.BlockDb) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "blocks"
}

// HandleBlock implements modules.InstructionModule
func (m *Module) HandleBlock(block types.Block) error {
	return m.db.SaveBlock(
		dbtypes.NewBlockRow(
			block.Slot, block.Height, block.Hash, block.Leader, block.Timestamp, len(block.Txs),
		),
	)
}
