package bank

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
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

// HandleBank implements modules.BankModule
func (m *Module) HandleBank(block types.Block) error {
	bankDb, ok := m.db.(db.BankDb)
	if !ok {
		return fmt.Errorf("bank is enabled, but your database does not implement BankDb")
	}
	err := HandleBank(block, bankDb)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Uint64("slot", block.Slot).Msg("handled block")
	return nil
}
