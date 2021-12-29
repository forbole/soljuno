package bank

import (
	"sync"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

type Module struct {
	db                  db.Database
	tasks               chan func()
	balanceEntries      []AccountBalanceEntry
	tokenBalanceEntries []TokenAccountBalanceEntry
	mtx                 sync.Mutex
}

func NewModule(db db.Database) *Module {
	return &Module{
		db:    db,
		tasks: make(chan func()),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bank"
}

// HandleBank implements modules.BankModule
func (m *Module) HandleBlock(block types.Block) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.balanceEntries = MergeAccountBalanceEntries(m.balanceEntries, GetAccountBalaceEntries(block))
	m.tokenBalanceEntries = MergeTokenAccountBalanceEntries(m.tokenBalanceEntries, GetTokenAccountBalaceEntries(block))
	log.Debug().Str("module", m.Name()).Uint64("slot", block.Slot).Msg("handled block")
	return nil
}
