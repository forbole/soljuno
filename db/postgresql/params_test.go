package postgresql_test

import (
	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestSaveInflationGovernorParam() {
	testCases := []struct {
		name     string
		data     dbtypes.InflationGovernorParamRow
		expected dbtypes.InflationGovernorParamRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewInflationGovernorParamRow(2, 0.1, 0.1, 0.1, 0.1, 0.1),
			expected: dbtypes.NewInflationGovernorParamRow(2, 0.1, 0.1, 0.1, 0.1, 0.1),
		},
		{
			name:     "update with lower epoch",
			data:     dbtypes.NewInflationGovernorParamRow(1, 0.2, 0.1, 0.1, 0.1, 0.1),
			expected: dbtypes.NewInflationGovernorParamRow(2, 0.1, 0.1, 0.1, 0.1, 0.1),
		},
		{
			name:     "update with same epoch",
			data:     dbtypes.NewInflationGovernorParamRow(2, 0.2, 0.1, 0.1, 0.1, 0.1),
			expected: dbtypes.NewInflationGovernorParamRow(2, 0.2, 0.1, 0.1, 0.1, 0.1),
		},
		{
			name:     "update with higher epoch",
			data:     dbtypes.NewInflationGovernorParamRow(3, 0.3, 0.1, 0.1, 0.1, 0.1),
			expected: dbtypes.NewInflationGovernorParamRow(3, 0.3, 0.1, 0.1, 0.1, 0.1),
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveInflationGovernorParam(tc.data)
			suite.Require().NoError(err)
			var rows []dbtypes.InflationGovernorParamRow
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM inflation_governor_param")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveEpochScheduleParam() {
	testCases := []struct {
		name     string
		data     dbtypes.EpochScheduleParamRow
		expected dbtypes.EpochScheduleParamRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewEpochScheduleParamRow(2, 1, 1, 1, true),
			expected: dbtypes.NewEpochScheduleParamRow(2, 1, 1, 1, true),
		},
		{
			name:     "update with lower epoch",
			data:     dbtypes.NewEpochScheduleParamRow(1, 2, 1, 1, true),
			expected: dbtypes.NewEpochScheduleParamRow(2, 1, 1, 1, true),
		},
		{
			name:     "update with same epoch",
			data:     dbtypes.NewEpochScheduleParamRow(2, 2, 1, 1, true),
			expected: dbtypes.NewEpochScheduleParamRow(2, 2, 1, 1, true),
		},
		{
			name:     "update with higher epoch",
			data:     dbtypes.NewEpochScheduleParamRow(3, 3, 1, 1, true),
			expected: dbtypes.NewEpochScheduleParamRow(3, 3, 1, 1, true),
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveEpochScheduleParam(tc.data)
			suite.Require().NoError(err)
			var rows []dbtypes.EpochScheduleParamRow
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM epoch_schedule_param")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
