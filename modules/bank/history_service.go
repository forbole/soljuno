package bank

import (
	"time"

	"github.com/forbole/soljuno/modules/history"
)

var _ history.HistroyService = &Module{}

func (m *Module) ExecHistory() error {
	balanceEntries, tokenBalanceEntries := m.getHistoryEntries()
	errChan := make(chan error)
	go func() {
		timestamp := time.Now()
		_, accounts, balances := EntriesToBalances(balanceEntries)
		errChan <- m.db.SaveAccountHistoryBalances(timestamp, accounts, balances)

	}()
	go func() {
		timestamp := time.Now()
		_, accounts, balances := EntriesToTokenBalances(tokenBalanceEntries)
		errChan <- m.db.SaveAccountHistoryTokenBalances(timestamp, accounts, balances)
	}()
	for i := 0; i < 2; i++ {
		err := <-errChan
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Module) getHistoryEntries() ([]AccountBalanceEntry, []TokenAccountBalanceEntry) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	balanceEntries := m.HistoryBalanceEntries
	tokenBalanceEntries := m.HistoryTokenBalanceEntries
	m.HistoryBalanceEntries = nil
	m.HistoryTokenBalanceEntries = nil
	return balanceEntries, tokenBalanceEntries
}

// update very 30 minutes
func (m *Module) Cron() string { return "*/30 * * * *" }
