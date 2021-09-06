package system

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

type Module struct {
	db db.Database
}

func NewModule(db db.Database) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "system"
}

func (m *Module) HandleMsg(msg types.Message, tx types.Tx) error {
	return nil
}
