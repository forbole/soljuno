package postgresql_test

import (
	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestSaveEpoch() {
	testCases := []struct {
		name     string
		data     dbtypes.EpochInfoRow
		expected dbtypes.EpochInfoRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewEpochInfoRow(2, 1),
			expected: dbtypes.NewEpochInfoRow(2, 1),
		},
		{
			name:     "update with lower epoch",
			data:     dbtypes.NewEpochInfoRow(1, 1),
			expected: dbtypes.NewEpochInfoRow(2, 1),
		},
		{
			name:     "update with higher epoch",
			data:     dbtypes.NewEpochInfoRow(3, 1),
			expected: dbtypes.NewEpochInfoRow(3, 1),
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveEpoch(tc.data)
			suite.Require().NoError(err)
			var rows []dbtypes.EpochInfoRow
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM epoch_info")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveInflationRate() {
	testCases := []struct {
		name     string
		data     dbtypes.InflationRateRow
		expected dbtypes.InflationRateRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewInflationRateRow(2, 0.1, 0.1, 0.1),
			expected: dbtypes.NewInflationRateRow(2, 0.1, 0.1, 0.1),
		},
		{
			name:     "update with lower epoch",
			data:     dbtypes.NewInflationRateRow(1, 0.2, 0.1, 0.1),
			expected: dbtypes.NewInflationRateRow(2, 0.1, 0.1, 0.1),
		},
		{
			name:     "update with same epoch",
			data:     dbtypes.NewInflationRateRow(2, 0.2, 0.1, 0.1),
			expected: dbtypes.NewInflationRateRow(2, 0.2, 0.1, 0.1),
		},
		{
			name:     "update with higher epoch",
			data:     dbtypes.NewInflationRateRow(3, 0.3, 0.1, 0.1),
			expected: dbtypes.NewInflationRateRow(3, 0.3, 0.1, 0.1),
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveInflationRate(tc.data)
			suite.Require().NoError(err)
			var rows []dbtypes.InflationRateRow
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM inflation_rate")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveSupplyInfo() {
	testCases := []struct {
		name     string
		data     dbtypes.SupplyInfoRow
		expected dbtypes.SupplyInfoRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewSupplyInfoRow(2, 1, 1, 1),
			expected: dbtypes.NewSupplyInfoRow(2, 1, 1, 1),
		},
		{
			name:     "update with lower epoch",
			data:     dbtypes.NewSupplyInfoRow(1, 2, 1, 1),
			expected: dbtypes.NewSupplyInfoRow(2, 1, 1, 1),
		},
		{
			name:     "update with same epoch",
			data:     dbtypes.NewSupplyInfoRow(2, 2, 1, 1),
			expected: dbtypes.NewSupplyInfoRow(2, 2, 1, 1),
		},
		{
			name:     "update with higher epoch",
			data:     dbtypes.NewSupplyInfoRow(3, 3, 1, 1),
			expected: dbtypes.NewSupplyInfoRow(3, 3, 1, 1),
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveSupplyInfo(tc.data)
			suite.Require().NoError(err)
			var rows []dbtypes.SupplyInfoRow
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM supply_info")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
