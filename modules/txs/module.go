package txs

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

var _ modules.Module = &Module{}

type Module struct {
	db     db.Database
	buffer chan types.Block
}

func NewModule(db db.Database) *Module {
	return &Module{
		db:     db,
		buffer: make(chan types.Block),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "txs"
}

// HandleBlock implements modules.MessageModule
func (m *Module) HandleBlock(block types.Block) error {
	m.buffer <- block
	return nil
}
