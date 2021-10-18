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

// HandleTx implements modules.TransactionModule
func (m *Module) HandleTx(tx types.Tx) error {
	bankDb, ok := m.db.(db.BankDb)
	if !ok {
		return fmt.Errorf("bank is enabled, but your database does not implement BankDb")
	}

	err := HandleTx(tx, bankDb)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Str("tx", tx.Hash).Uint64("slot", tx.Slot).
		Msg("handled tx")
	return nil
}

// HandleTx implements modules.BlockModule
func (m *Module) HandleBlock(block types.Block) error {
	bankDb, ok := m.db.(db.BankDb)
	if !ok {
		return fmt.Errorf("bank is enabled, but your database does not implement BankDb")
	}

	err := HandleBlock(block, bankDb)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Uint64("slot", block.Slot).Msg("handled block")
	return nil
}
