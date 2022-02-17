package bank

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", m.Name()).Msg("setting up periodic tasks")
	if _, err := scheduler.Every(10).Second().Do(func() {
		m.mtx.Lock()
		defer m.mtx.Unlock()
		balances := m.balanceEntries
		tokenBalances := m.tokenBalanceEntries
		m.balanceEntries = nil
		m.tokenBalanceEntries = nil
		err := m.updateBalances(balances, tokenBalances)
		if err != nil {
			log.Error().Str("module", m.Name()).Err(err).Send()
		}
	}); err != nil {
		return fmt.Errorf("error while setting up bank periodic operation: %s", err)
	}
	return nil
}

func (m *Module) updateBalances(balances []AccountBalanceEntry, tokenBalances []TokenAccountBalanceEntry) error {
	errChan := make(chan error, 2)
	go func() {
		errChan <- m.db.SaveAccountBalances(EntriesToBalances(balances))
	}()
	go func() {
		errChan <- m.db.SaveAccountTokenBalances(EntriesToTokenBalances(tokenBalances))
	}()
	for i := 0; i < len(errChan); i++ {
		err := <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}
