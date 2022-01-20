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

	insertStmt := `INSERT INTO token_unit (address, price_id, unit_name, logo_uri, description, website) VALUES`
	paramsStmt := ""
	conflictStmt := `ON CONFLICT (address) DO UPDATE
		SET price_id = EXCLUDED.price_id,
			unit_name = EXCLUDED.unit_name,
			description = EXCLUDED.description,
			website = EXCLUDED.website
	`
	var params []interface{}
	paramNumber := 6

	for i, unit := range units {
		vi := i * paramNumber
		paramsStmt += getParamsStmt(vi, paramNumber)
		params = append(params, unit.Address, unit.PriceID, unit.Name, unit.LogoURI, unit.Description, unit.Website)
	}
	return db.insertWithParams(insertStmt, paramsStmt[:len(paramsStmt)-1], conflictStmt, params)
}

// SaveTokensPrices allows to save the given prices as the most updated ones
func (db *Database) SaveTokenPrices(prices []dbtypes.TokenPriceRow) error {
	if len(prices) == 0 {
		return nil
	}

	insertStmt := `INSERT INTO token_price (id, price, market_cap, symbol, timestamp) VALUES`
	paramsStmt := ""
	conflictStmt := `
	ON CONFLICT (id) DO UPDATE 
		SET price = excluded.price,
			market_cap = excluded.market_cap,
			symbol = excluded.symbol,
			timestamp = excluded.timestamp
	WHERE token_price.timestamp <= excluded.timestamp`
	var params []interface{}
	paramNumber := 5

	for i, ticker := range prices {
		vi := i * paramNumber
		paramsStmt += getParamsStmt(vi, paramNumber)
		params = append(params, ticker.ID, ticker.Price, ticker.MarketCap, ticker.Symbol, ticker.Timestamp)
	}

	return db.insertWithParams(insertStmt, paramsStmt[:len(paramsStmt)-1], conflictStmt, params)
}

// SaveTokensPrices implements db.PriceDb
func (db *Database) SaveHistoryTokenPrices(prices []dbtypes.TokenPriceRow) error {
	if len(prices) == 0 {
		return nil
	}

	insertStmt := `INSERT INTO token_price_history (id, price, market_cap, symbol, timestamp) VALUES`
	paramsStmt := ""
	conflictStmt := ""

	var params []interface{}
	paramNumber := 5

	for i, ticker := range prices {
		vi := i * paramNumber
		paramsStmt += getParamsStmt(vi, paramNumber)
		params = append(params, ticker.ID, ticker.Price, ticker.MarketCap, ticker.Symbol, ticker.Timestamp)
	}

	return db.insertWithParams(insertStmt, paramsStmt[:len(paramsStmt)-1], conflictStmt, params)
}
