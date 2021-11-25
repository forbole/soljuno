package postgresql_test

import (
	"time"

	"github.com/forbole/soljuno/types"
)

func (suite *DbTestSuite) insertToken(unit types.TokenUnit) {
	err := suite.database.SaveTokenUnit(unit)
	suite.Require().NoError(err)
}

func (suite *DbTestSuite) TestGetTokenUnits() {
	suite.insertToken(types.NewTokenUnit("sol", "sol", "sol"))
	suite.insertToken(types.NewTokenUnit("usdc", "usdc", "usdc"))

	units, err := suite.database.GetTokenUnits()
	suite.Require().NoError(err)

	var expected = []types.TokenUnit{
		types.NewTokenUnit("sol", "sol", "sol"),
		types.NewTokenUnit("usdc", "usdc", "usdc"),
	}
	suite.Require().Len(units, len(expected))
	for _, name := range expected {
		suite.Require().Contains(units, name)
	}
}

func (suite *DbTestSuite) TestSaveTokenPrices() {
	suite.insertToken(types.NewTokenUnit("sol", "sol", "sol"))
	suite.insertToken(types.NewTokenUnit("usdc", "usdc", "usdc"))

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
		suite.Require().True(expected[i].Timestamp.Equal(row.Timestamp))
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
		suite.Require().True(expected[i].Timestamp.Equal(row.Timestamp))
	}
}
