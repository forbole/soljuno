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
			data: dbtypes.NewTxRow("hash", 1, true, 500, []string{"logs"}, "{}"),
			expected: []dbtypes.TxRow{
				dbtypes.NewTxRow("hash", 1, true, 500, []string{"logs"}, "{}"),
			},
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
				PartitionId: 100},
			shouldErr: true,
		},
		{
			name: "insert the existed data",
			data: dbtypes.NewTxRow("hash", 1, true, 500, []string{"logs"}, "{}"),
			expected: []dbtypes.TxRow{
				dbtypes.NewTxRow("hash", 1, true, 500, []string{"logs"}, "{}"),
			},
		},
		{
			name: "insert the new data",
			data: dbtypes.NewTxRow("hash2", 2, true, 500, []string{"logs"}, "{}"),
			expected: []dbtypes.TxRow{
				dbtypes.NewTxRow("hash", 1, true, 500, []string{"logs"}, "{}"),
				dbtypes.NewTxRow("hash2", 1, true, 500, []string{"logs"}, "{}"),
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
