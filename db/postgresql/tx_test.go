package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestSaveTxs() {
	err := suite.database.CreateTxPartition(0)
	suite.Require().NoError(err)

	testCases := []struct {
		name      string
		data      dbtypes.TxRow
		expected  []dbtypes.TxRow
		shouldErr bool
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewTxRow("signature", 1, true, 500, []string{"logs"}, 2),
			expected: []dbtypes.TxRow{
				dbtypes.NewTxRow("signature", 1, true, 500, []string{"logs"}, 2),
			},
		},
		{
			name: "insert the wrong tx",
			data: dbtypes.TxRow{
				Signature:   "signature",
				Slot:        1,
				Error:       true,
				Fee:         500,
				Logs:        []string{"logs"},
				PartitionId: 100,
			},
			shouldErr: true,
		},
		{
			name: "insert the existed data",
			data: dbtypes.NewTxRow("signature", 1, true, 500, []string{"logs"}, 2),
			expected: []dbtypes.TxRow{
				dbtypes.NewTxRow("signature", 1, true, 500, []string{"logs"}, 2),
			},
		},
		{
			name: "insert the new data",
			data: dbtypes.NewTxRow("signature2", 2, true, 500, []string{"logs"}, 2),
			expected: []dbtypes.TxRow{
				dbtypes.NewTxRow("signature", 1, true, 500, []string{"logs"}, 2),
				dbtypes.NewTxRow("signature2", 2, true, 500, []string{"logs"}, 2),
			},
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
				suite.Require().Len(rows, len(tc.expected))
			}
		})
	}
}

func (suite *DbTestSuite) TestPruneTxsBeforeSlot() {
	err := suite.database.CreateTxPartition(0)
	suite.Require().NoError(err)

	err = suite.database.SaveTxs([]dbtypes.TxRow{
		dbtypes.NewTxRow("signature", 1, true, 500, []string{"logs"}, 2),
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
