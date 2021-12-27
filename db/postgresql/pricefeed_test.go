package postgresql_test

import (
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) insertToken(unit dbtypes.TokenUnitRow) {
	err := suite.database.SaveTokenUnit(unit)
	suite.Require().NoError(err)
}

func (suite *DbTestSuite) TestGetTokenUnits() {
	suite.insertToken(dbtypes.NewTokenUnitRow("sol", "sol", "sol", "url"))
	suite.insertToken(dbtypes.NewTokenUnitRow("usdc", "usdc", "usdc", "url"))

	units, err := suite.database.GetTokenUnits()
	suite.Require().NoError(err)

	var expected = []dbtypes.TokenUnitRow{
		dbtypes.NewTokenUnitRow("sol", "sol", "sol", "url"),
		dbtypes.NewTokenUnitRow("usdc", "usdc", "usdc", "url"),
	}
	suite.Require().Len(units, len(expected))
	for _, name := range expected {
		suite.Require().Contains(units, name)
	}
}

func (suite *DbTestSuite) TestSaveTokenPrices() {
	suite.insertToken(dbtypes.NewTokenUnitRow("sol", "sol", "sol", "url"))
	suite.insertToken(dbtypes.NewTokenUnitRow("usdc", "usdc", "usdc", "url"))

	// Save data
	tickers := []dbtypes.TokenPriceRow{
		dbtypes.NewTokenPriceRow(
			"sol",
			100.01,
			10,
			"sol",
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
		dbtypes.NewTokenPriceRow(
			"usdc",
			200.01,
			20,
			"usdc",
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
	}
	err := suite.database.SaveTokensPrices(tickers)
	suite.Require().NoError(err)

	// Verify data
	expected := []dbtypes.TokenPriceRow{
		dbtypes.NewTokenPriceRow(
			"sol",
			100.01,
			10,
			"sol",
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
		dbtypes.NewTokenPriceRow(
			"usdc",
			200.01,
			20,
			"usdc",
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
	}

	var rows []dbtypes.TokenPriceRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM token_price`)
	suite.Require().NoError(err)
	for i, row := range rows {
		suite.Require().True(expected[i].ID == row.ID)
		suite.Require().True(expected[i].Price == row.Price)
		suite.Require().True(expected[i].Timestamp.Equal(row.Timestamp))
	}
	rows = nil

	// Update data
	tickers = []dbtypes.TokenPriceRow{
		dbtypes.NewTokenPriceRow(
			"sol",
			100.01,
			10,
			"sol",
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
		dbtypes.NewTokenPriceRow(
			"usdc",
			1,
			20,
			"usdc",
			time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC),
		),
	}
	err = suite.database.SaveTokensPrices(tickers)
	suite.Require().NoError(err)

	// Verify data
	expected = []dbtypes.TokenPriceRow{
		dbtypes.NewTokenPriceRow(
			"sol",
			100.01,
			10,
			"sol",
			time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC),
		),
		dbtypes.NewTokenPriceRow(
			"usdc",
			1,
			20,
			"usdc",
			time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC),
		),
	}

	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM token_price ORDER BY timestamp`)
	suite.Require().NoError(err)
	for i, row := range rows {
		suite.Require().True(expected[i].ID == row.ID)
		suite.Require().True(expected[i].Price == row.Price)
		suite.Require().True(expected[i].MarketCap == row.MarketCap)
		suite.Require().True(expected[i].Timestamp.Equal(row.Timestamp))
	}
}
