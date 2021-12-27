package postgresql

import (
	"fmt"

	dbtypes "github.com/forbole/soljuno/db/types"
)

func (db *Database) GetTokenUnits() ([]dbtypes.TokenUnitRow, error) {
	query := `SELECT * FROM token_unit`

	var units []dbtypes.TokenUnitRow
	err := db.Sqlx.Select(&units, query)
	if err != nil {
		return nil, err
	}
	return units, nil
}

func (db *Database) SaveTokenUnit(unit dbtypes.TokenUnitRow) error {
	stmt := `INSERT INTO token_unit (address, price_id, unit_name, logo_url) VALUES ($1, $2, $3, $4)`
	_, err := db.Sqlx.Exec(stmt, unit.Address, unit.PriceID, unit.Name, unit.LogoURL)
	return err
}

// SaveTokensPrices allows to save the given prices as the most updated ones
func (db *Database) SaveTokensPrices(prices []dbtypes.TokenPriceRow) error {
	if len(prices) == 0 {
		return nil
	}

	stmt := `INSERT INTO token_price (id, price, market_cap, symbol, timestamp) VALUES`
	var param []interface{}
	paramNumber := 5

	for i, ticker := range prices {
		vi := i * paramNumber
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d),", vi+1, vi+2, vi+3, vi+4, vi+5)
		param = append(param, ticker.ID, ticker.Price, ticker.MarketCap, ticker.Symbol, ticker.Timestamp)
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ","

	stmt += `
ON CONFLICT (id) DO UPDATE 
	SET price = excluded.price,
	    market_cap = excluded.market_cap,
	    timestamp = excluded.timestamp
WHERE token_price.timestamp <= excluded.timestamp`
	_, err := db.Sqlx.Exec(stmt, param...)
	if err != nil {
		return fmt.Errorf("error while saving tokens prices: %s", err)
	}

	return nil
}
