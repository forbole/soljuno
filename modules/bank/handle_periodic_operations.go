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
		m.mtx.Lock()
		balances := m.balanceEntries
		tokenBalances := m.tokenBalanceEntries
		m.balanceEntries = nil
		m.tokenBalanceEntries = nil
		m.mtx.Unlock()

		utils.WatchMethod(func() error { return m.updateBalances(balances, tokenBalances) })
	}); err != nil {
		return fmt.Errorf("error while setting up vote periodic operation: %s", err)
	}
	return nil
}

func (m *Module) updateBalances(balances []AccountBalanceEntry, tokenBalances []TokenAccountBalanceEntry) error {
	errChan := make(chan error)
	go func() {
		errChan <- m.db.SaveAccountBalances(EntriesToBalances(balances))

	}()
	go func() {
		errChan <- m.db.SaveAccountTokenBalances(EntriesToTokenBalances(tokenBalances))
	}()
	for i := 0; i < 2; i++ {
		err := <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}
