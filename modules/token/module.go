package token

import (
	"fmt"

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
func (m *Module) HandleMsg(msg types.Message, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	tokenDb, ok := m.db.(db.TokenDb)
	if !ok {
		return fmt.Errorf("pruning is enabled, but your database does not implement PruningDb")
	}
	return HandleMsg(msg, tx, tokenDb)
}
