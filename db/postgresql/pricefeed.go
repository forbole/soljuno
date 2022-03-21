package postgresql

import (
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

func (db *Database) SaveTokenUnits(units []dbtypes.TokenUnitRow) error {
	if len(units) == 0 {
		return nil
	}

	insertStmt := `INSERT INTO token_unit (mint, price_id, unit_name, logo_uri, description, website) VALUES`
	conflictStmt := `ON CONFLICT (mint) DO UPDATE
		SET price_id = EXCLUDED.price_id,
			unit_name = EXCLUDED.unit_name,
			description = EXCLUDED.description,
			website = EXCLUDED.website
	`
	var params []interface{}
	paramsNumber := 6
	params = make([]interface{}, 0, paramsNumber*len(units))
	for _, unit := range units {
		params = append(params, unit.Mint, unit.PriceID, unit.Name, unit.LogoURI, unit.Description, unit.Website)
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

// SaveTokensPrices allows to save the given prices as the most updated ones
func (db *Database) SaveTokenPrices(prices []dbtypes.TokenPriceRow) error {
	if len(prices) == 0 {
		return nil
	}

	insertStmt := `INSERT INTO token_price (id, price, market_cap, symbol, timestamp, volume) VALUES`
	conflictStmt := `
	ON CONFLICT (id) DO UPDATE 
		SET price = excluded.price,
			market_cap = excluded.market_cap,
			symbol = excluded.symbol,
			timestamp = excluded.timestamp,
			volume = excluded.volume
	WHERE token_price.timestamp <= excluded.timestamp`
	var params []interface{}
	paramsNumber := 6
	params = make([]interface{}, 0, paramsNumber*len(prices))

	for _, ticker := range prices {
		params = append(params, ticker.ID, ticker.Price, ticker.MarketCap, ticker.Symbol, ticker.Timestamp, ticker.Volume)
	}

	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

// SaveTokensPrices implements db.PriceDb
func (db *Database) SaveHistoryTokenPrices(prices []dbtypes.TokenPriceRow) error {
	if len(prices) == 0 {
		return nil
	}

	insertStmt := `INSERT INTO token_price_history (id, price, market_cap, symbol, timestamp, volume) VALUES`
	conflictStmt := ""
	var params []interface{}
	paramsNumber := 6
	params = make([]interface{}, 0, paramsNumber*len(prices))

	for _, ticker := range prices {
		params = append(params, ticker.ID, ticker.Price, ticker.MarketCap, ticker.Symbol, ticker.Timestamp, ticker.Volume)
	}

	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}
