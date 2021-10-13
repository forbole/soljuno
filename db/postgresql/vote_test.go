package postgresql_test

func (suite *DbTestSuite) TestSaveNonce() {
	type VoteAccountRow struct {
		Address    string `db:"address"`
		Slot       uint64 `db:"slot"`
		Node       string `db:"node"`
		Voter      string `db:"voter"`
		Withdrawer string `db:"withdrawer"`
		Commission uint8  `db:"commission"`
	}

	testCases := []struct {
		name     string
		data     VoteAccountRow
		expected VoteAccountRow
	}{
		{
			name: "initialize the data",
			data: VoteAccountRow{
				"address", 1, "node", "voter", "withdrawer", 5,
			},
			expected: VoteAccountRow{
				"address", 1, "node", "voter", "withdrawer", 5,
			},
		},
		{
			name: "update with lower slot",
			data: VoteAccountRow{
				"address", 0, "pre_node", "voter", "withdrawer", 5,
			},
			expected: VoteAccountRow{
				"address", 1, "node", "voter", "withdrawer", 5,
			},
		},
		{
			name: "update with same slot",
			data: VoteAccountRow{
				"address", 1, "same_node", "voter", "withdrawer", 5,
			},
			expected: VoteAccountRow{
				"address", 1, "same_node", "voter", "withdrawer", 5,
			},
		},
		{
			name: "update with higher slot",
			data: VoteAccountRow{
				"address", 2, "new_node", "voter", "withdrawer", 5,
			},
			expected: VoteAccountRow{
				"address", 2, "new_node", "voter", "withdrawer", 5,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveVoteAccount(
				tc.data.Address,
				tc.data.Slot,
				tc.data.Node,
				tc.data.Voter,
				tc.data.Withdrawer,
				tc.data.Commission,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []VoteAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM vote_accounts")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
