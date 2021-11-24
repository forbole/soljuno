package postgresql

import (
	"fmt"

	"github.com/forbole/soljuno/types"
)

// SaveTokensPrices allows to save the given prices as the most updated ones
func (db *Database) SaveTokensPrices(prices []types.TokenPrice) error {
	if len(prices) == 0 {
		return nil
	}

	query := `INSERT INTO token_price (unit_name, price, market_cap, timestamp) VALUES`
	var param []interface{}

	for i, ticker := range prices {
		vi := i * 4
		query += fmt.Sprintf("($%d,$%d,$%d,$%d),", vi+1, vi+2, vi+3, vi+4)
		param = append(param, ticker.UnitName, ticker.Price, ticker.MarketCap, ticker.Timestamp)
	}

	query = query[:len(query)-1] // Remove trailing ","
	_, err := db.Sqlx.Exec(query, param...)
	if err != nil {
		return fmt.Errorf("error while saving tokens prices: %s", err)
	}

	return nil
}
