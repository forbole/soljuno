package bank

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", m.Name()).Msg("setting up periodic tasks")
	if _, err := scheduler.Every(10).Second().Do(func() {
		utils.WatchMethod(m, m.RunPeriodicOperations)
	}); err != nil {
		return fmt.Errorf("error while setting up bank periodic operation: %s", err)
	}
	return nil
}

func (m *Module) RunPeriodicOperations() error {
	return HandlePeriodicOperations(m, m.db)
}

func HandlePeriodicOperations(m *Module, db db.BankDb) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	balances := m.BalanceEntries
	tokenBalances := m.TokenBalanceEntries
	err := updateBalances(db, balances, tokenBalances)
	if err != nil {
		return err
	}

	// if success, clear the stored entries
	m.BalanceEntries = nil
	m.TokenBalanceEntries = nil
	return nil
}

func updateBalances(db db.BankDb, balances []AccountBalanceEntry, tokenBalances []TokenAccountBalanceEntry) error {
	errChan := make(chan error, 2)
	go func() {
		errChan <- db.SaveAccountBalances(EntriesToBalances(balances))
	}()
	go func() {
		errChan <- db.SaveAccountTokenBalances(EntriesToTokenBalances(tokenBalances))
	}()
	for i := 0; i < 2; i++ {
		err := <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}
