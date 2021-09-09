package bank

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
	return "bank"
}

func (m *Module) HandleTx(tx types.Tx) error {
	bankDb, ok := m.db.(db.BankDb)
	if !ok {
		return fmt.Errorf("bank is enabled, but your database does not implement BankDb")
	}
	return HandleTx(tx, bankDb)
}
