package bank

import (
	"fmt"

	"github.com/forbole/soljuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", m.Name()).Msg("setting up periodic tasks")

	if _, err := scheduler.Every(10).Second().Do(func() {
		utils.WatchMethod(m.updateBalances)
	}); err != nil {
		return fmt.Errorf("error while setting up vote periodic operation: %s", err)
	}

	return nil
}
func (m *Module) updateBalances() error {
	m.mtx.Lock()
	balances := m.balanceEntries
	tokenBalances := m.tokenBalanceEntries
	m.balanceEntries = nil
	m.tokenBalanceEntries = nil
	m.mtx.Unlock()
	err := m.db.SaveAccountBalances(EntriesToBalances(balances))
	if err != nil {
		return err
	}
	err = m.db.SaveAccountTokenBalances(EntriesToTokenBalances(tokenBalances))
	return err
}
