package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestSaveTxs() {
	err := suite.database.CreateTxPartition(0)
	suite.Require().NoError(err)

	testCases := []struct {
		name        string
		data        dbtypes.TxRow
		expectedLen int
		shouldErr   bool
	}{
		{
			name:        "initialize the data",
			data:        dbtypes.NewTxRow("hash", 1, true, 500, []string{"logs"}, "{}"),
			expectedLen: 1,
		},
		{
			name: "insert the wrong tx",
			data: dbtypes.TxRow{
				Hash:        "hash",
				Slot:        1,
				Error:       true,
				Fee:         500,
				Logs:        []string{"logs"},
				Messages:    "{}",
				PartitionId: 100,
			},
			shouldErr: true,
		},
		{
			name:        "insert the existed data",
			data:        dbtypes.NewTxRow("hash", 1, true, 500, []string{"logs"}, "{}"),
			expectedLen: 1,
		},
		{
			name:        "insert the new data",
			data:        dbtypes.NewTxRow("hash2", 2, true, 500, []string{"logs"}, "{}"),
			expectedLen: 2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveTxs([]dbtypes.TxRow{tc.data})
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				// Verify the data
				rows := []dbtypes.TxRow{}
				err = suite.database.Sqlx.Select(&rows, "SELECT * FROM transaction")
				suite.Require().NoError(err)
				suite.Require().Len(rows, tc.expectedLen)
			}
		})
	}
}

func (suite *DbTestSuite) TestPruneTxsBeforeSlot() {
	err := suite.database.CreateTxPartition(0)
	suite.Require().NoError(err)

	err = suite.database.SaveTxs([]dbtypes.TxRow{
		dbtypes.NewTxRow("hash", 1, true, 500, []string{"logs"}, "{}"),
	})
	suite.Require().NoError(err)

	rows := []dbtypes.TxRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM transaction")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)

	err = suite.database.PruneTxsBeforeSlot(10000)
	suite.Require().NoError(err)

	rows = nil
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM transaction")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}
