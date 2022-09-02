package postgresql_test

import (
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestGetLastBlock() {
	_, err := suite.database.GetLastBlock()
	suite.Require().Error(err)

	date := time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC)
	err = suite.database.SaveBlock(dbtypes.NewBlockRow(
		1, 1, "hash", "leader", date, 0,
	))
	suite.Require().NoError(err)

	expected := dbtypes.NewBlockRow(1, 1, "hash", "leader", date, 0)

	result, err := suite.database.GetLastBlock()
	suite.Require().NoError(err)
	suite.Require().True(expected.Equal(result))
}

func (suite *DbTestSuite) TestGetBlockHourAgo() {
	timeAgo, err := time.Parse(time.RFC3339, "2020-01-01T15:00:00Z")
	suite.Require().NoError(err)

	timeNow := timeAgo.Add(time.Hour)
	result, found, err := suite.database.GetBlockHourAgo(timeNow)
	suite.Require().False(found)
	suite.Require().NoError(err)

	slot := uint64(1)
	err = suite.database.SaveBlock(dbtypes.NewBlockRow(
		1, 1, "leader", "hash", timeAgo, 0,
	))
	suite.Require().NoError(err)
	result, found, err = suite.database.GetBlockHourAgo(timeNow)
	suite.Require().True(found)
	suite.Require().NoError(err)

	suite.Require().True(result.Timestamp.Equal(timeAgo))
	suite.Require().Equal(slot, result.Slot)
}

func (suite *DbTestSuite) TestSaveAverageSlotTimePerHour() {
	testCases := []struct {
		name     string
		data     dbtypes.AverageTimeRow
		expected dbtypes.AverageTimeRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewAverageTimeRow(10, 0.5),
			expected: dbtypes.NewAverageTimeRow(10, 0.5),
		},
		{
			name:     "update with lower slot",
			data:     dbtypes.NewAverageTimeRow(5, 1),
			expected: dbtypes.NewAverageTimeRow(10, 0.5),
		},
		{
			name:     "update with same slot",
			data:     dbtypes.NewAverageTimeRow(10, 1),
			expected: dbtypes.NewAverageTimeRow(10, 1),
		},
		{
			name:     "update with higher slot",
			data:     dbtypes.NewAverageTimeRow(15, 1),
			expected: dbtypes.NewAverageTimeRow(15, 1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveAverageSlotTimePerHour(tc.data.Slot, tc.data.AverageTime)
			suite.Require().NoError(err)
			var rows []dbtypes.AverageTimeRow
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM average_slot_time_per_hour")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
