package pricefeed

import "github.com/forbole/soljuno/modules/history"

var _ history.HistroyService = &Module{}

func (m *Module) ExecHistory() error {
	prices, err := m.getTokenPrices()
	if err != nil {
		return err
	}
	return m.db.SaveHistoryTokenPrices(prices)
}
