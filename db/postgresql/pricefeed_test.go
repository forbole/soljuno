package postgresql_test

import (
	"fmt"
	"time"

	"github.com/forbole/soljuno/types"
)

func (suite *DbTestSuite) insertToken(name string, address string) {
	query := fmt.Sprintf(
		`INSERT INTO token_unit (token_name, address) VALUES ('%[1]s', '%[2]s')`,
		name, address)
	_, err := suite.database.Sqlx.Query(query)
	suite.Require().NoError(err)
}

func (suite *DbTestSuite) SaveTokenPrice() {
	suite.insertToken("sol", "sol")
	suite.insertToken("usdc", "usdc")

	type TokenPriceRow struct {
		Name      string    `db:"unit_name"`
		Price     float64   `db:"price"`
		MarketCap int64     `db:"market_cap"`
		Timestamp time.Time `db:"timestamp"`
	}

	// Save data
	tickers := []types.TokenPrice{
		types.NewTokenPrice(
			"sol",
			100.01,
			10,
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
		types.NewTokenPrice(
			"usdc",
			200.01,
			20,
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
	}
	err := suite.database.SaveTokensPrices(tickers)
	suite.Require().NoError(err)

	// Verify data
	expected := []TokenPriceRow{
		{
			"sol",
			100.01,
			10,
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		},
		{
			"usdc",
			200.01,
			20,
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		},
	}

	var rows []TokenPriceRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM token_price`)
	suite.Require().NoError(err)
	for i, row := range rows {
		suite.Require().True(expected[i].Name == row.Name)
		suite.Require().True(expected[i].Price == row.Price)
		suite.Require().True(expected[i].MarketCap == row.MarketCap)
		suite.Require().True(expected[i].Timestamp == row.Timestamp)
	}

	// Update data
	tickers = []types.TokenPrice{
		types.NewTokenPrice(
			"sol",
			100.01,
			10,
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
		types.NewTokenPrice(
			"usdc",
			1,
			20,
			time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC),
		),
	}
	err = suite.database.SaveTokensPrices(tickers)
	suite.Require().NoError(err)

	// Verify data
	expected = []TokenPriceRow{
		{
			"sol",
			100.01,
			10,
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		},
		{
			"usdc",
			1,
			20,
			time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC),
		},
	}

	rows = []TokenPriceRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM token_price ORDER BY timestamp`)
	suite.Require().NoError(err)
	for i, row := range rows {
		suite.Require().True(expected[i].Name == row.Name)
		suite.Require().True(expected[i].Price == row.Price)
		suite.Require().True(expected[i].MarketCap == row.MarketCap)
		suite.Require().True(expected[i].Timestamp == row.Timestamp)
	}
}
