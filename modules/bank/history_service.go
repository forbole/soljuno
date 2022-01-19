package bank

import (
	"time"

	"github.com/forbole/soljuno/modules/history"
)

var _ history.HistroyService = &Module{}

func (m *Module) ExecHistory() error {
	m.mtx.Lock()
	balanceEntries := m.balanceEntries
	tokenBalanceEntries := m.tokenBalanceEntries
	m.balanceEntries = nil
	m.tokenBalanceEntries = nil
	m.mtx.Unlock()

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
