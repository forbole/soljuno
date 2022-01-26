package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestSaveMessages() {
	err := suite.database.CreateMsgPartition(0)
	suite.Require().NoError(err)

	testCases := []struct {
		name        string
		data        dbtypes.MsgRow
		expectedLen int
		shouldErr   bool
	}{
		{
			name:        "initialize the data",
			data:        dbtypes.NewMsgRow("hash", 1, 1, 0, "program", []string{"address"}, "raw", "type", "{}"),
			expectedLen: 1,
		},
		{
			name: "insert the wrong tx",
			data: dbtypes.MsgRow{
				TxHash:           "txHash",
				Slot:             1,
				Index:            1,
				InnerIndex:       0,
				Program:          "program",
				InvolvedAccounts: []string{"address"},
				RawData:          "raw",
				Type:             "type",
				Value:            "{}",
				PartitionId:      100,
			},
			shouldErr: true,
		},
		{
			name:        "insert the existed data",
			data:        dbtypes.NewMsgRow("hash", 1, 1, 0, "program", []string{"address"}, "raw", "type", "{}"),
			expectedLen: 1,
		},
		{
			name:        "insert the new data",
			data:        dbtypes.NewMsgRow("hash", 1, 1, 1, "program", []string{"address"}, "raw", "type", "{}"),
			expectedLen: 2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveMessages([]dbtypes.MsgRow{tc.data})
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				// Verify the data
				rows := []dbtypes.MsgRow{}
				err = suite.database.Sqlx.Select(&rows, "SELECT * FROM message")
				suite.Require().NoError(err)
				suite.Require().Len(rows, tc.expectedLen)
			}
		})
	}
}

func (suite *DbTestSuite) TestPruneMsgsBeforeSlot() {
	err := suite.database.CreateMsgPartition(0)
	suite.Require().NoError(err)

	err = suite.database.SaveMessages([]dbtypes.MsgRow{
		dbtypes.NewMsgRow("hash", 1, 1, 0, "program", []string{"address"}, "raw", "type", "{}"),
	})
	suite.Require().NoError(err)

	err = suite.database.PruneMsgsBeforeSlot(10000)
	suite.Require().NoError(err)

	rows := []dbtypes.MsgRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM message")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}
