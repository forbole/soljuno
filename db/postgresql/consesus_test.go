package postgresql_test

import (
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/types"
)

func (suite *DbTestSuite) TestGetLastBlock() {
	_, err := suite.database.GetLastBlock()
	suite.Require().Error(err)

	date := time.Date(2020, 10, 10, 15, 00, 00, 000, time.UTC)
	err = suite.database.SaveBlock(types.Block{
		Slot:      1,
		Height:    1,
		Proposer:  "proposer",
		Hash:      "hash",
		Timestamp: date,
	})
	suite.Require().NoError(err)

	expected := dbtypes.NewBlockRow(1, 1, "hash", "proposer", date)

	result, err := suite.database.GetLastBlock()
	suite.Require().NoError(err)
	suite.Require().Equal(expected.Slot, result.Slot)
	suite.Require().Equal(expected.Height, result.Height)
	suite.Require().Equal(expected.Proposer, result.Proposer)
	suite.Require().Equal(expected.Hash, result.Hash)
	suite.Require().Equal(expected.Timestamp.Unix(), result.Timestamp.Unix())
}

func (suite *DbTestSuite) TestGetBlockHourAgo() {
	timeAgo, err := time.Parse(time.RFC3339, "2020-01-01T15:00:00Z")
	suite.Require().NoError(err)

	timeNow := timeAgo.Add(time.Hour)
	result, err := suite.database.GetBlockHourAgo(timeNow)
	suite.Require().Error(err)

	slot := uint64(1)
	err = suite.database.SaveBlock(types.Block{
		Slot:      1,
		Height:    1,
		Proposer:  "proposer",
		Hash:      "hash",
		Timestamp: timeAgo,
	})
	suite.Require().NoError(err)
	result, err = suite.database.GetBlockHourAgo(timeNow)
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
