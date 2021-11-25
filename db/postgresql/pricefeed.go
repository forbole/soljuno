package postgresql

import (
	"fmt"

	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/types"
)

func (db *Database) GetTokenUnits() ([]types.TokenUnit, error) {
	query := `SELECT * FROM token_unit`

	var dbUnits []dbtypes.TokenUnitRow
	err := db.Sqlx.Select(&dbUnits, query)
	if err != nil {
		return nil, err
	}

	var units = make([]types.TokenUnit, len(dbUnits))
	for index, unit := range dbUnits {
		units[index] = types.NewTokenUnit(
			unit.PriceID,
			unit.Address,
			unit.TokenName,
		)
	}

	return units, nil
}

func (db *Database) SaveTokenUnit(unit types.TokenUnit) error {
	stmt := `INSERT INTO token_unit (price_id, address, token_name) VALUES ($1, $2, $3)`
	_, err := db.Sqlx.Exec(stmt, unit.ID, unit.Address, unit.Name)
	return err
}

// SaveTokensPrices allows to save the given prices as the most updated ones
func (db *Database) SaveTokensPrices(prices []types.TokenPrice) error {
	if len(prices) == 0 {
		return nil
	}

	stmt := `INSERT INTO token_price (unit_name, price, market_cap, timestamp) VALUES`
	var param []interface{}

	for i, ticker := range prices {
		vi := i * 4
		stmt += fmt.Sprintf("($%d,$%d,$%d,$%d),", vi+1, vi+2, vi+3, vi+4)
		param = append(param, ticker.UnitName, ticker.Price, ticker.MarketCap, ticker.Timestamp)
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ","
	_, err := db.Sqlx.Exec(stmt, param...)
	if err != nil {
		return fmt.Errorf("error while saving tokens prices: %s", err)
	}

	return nil
}
