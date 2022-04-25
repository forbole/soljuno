package bank

import (
	"sync"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

var (
	_ modules.BlockModule              = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	db                  db.BankDb
	tasks               chan func()
	BalanceEntries      []AccountBalanceEntry
	TokenBalanceEntries []TokenAccountBalanceEntry

	HistoryBalanceEntries      []AccountBalanceEntry
	HistoryTokenBalanceEntries []TokenAccountBalanceEntry

	mtx sync.Mutex
}

func NewModule(db db.BankDb) *Module {
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

	balanceEntries := getAccountBalaceEntries(block)
	m.BalanceEntries = MergeAccountBalanceEntries(m.BalanceEntries, getAccountBalaceEntries(block))
	m.HistoryBalanceEntries = MergeAccountBalanceEntries(m.HistoryBalanceEntries, balanceEntries)

	tokenBalanceEntries := getTokenAccountBalaceEntries(block)
	m.TokenBalanceEntries = MergeTokenAccountBalanceEntries(m.TokenBalanceEntries, tokenBalanceEntries)
	m.HistoryTokenBalanceEntries = MergeTokenAccountBalanceEntries(m.HistoryTokenBalanceEntries, tokenBalanceEntries)
	log.Debug().Str("module", m.Name()).Uint64("slot", block.Slot).Msg("handled block")
	return nil
}
