package token

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
	return "token"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg types.Message, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	return m.db.SaveMessage(msg)
}
