package postgresql_test

func (suite *DbTestSuite) TestSaveNonceAccount() {
	type NonceRow struct {
		Address              string `db:"address"`
		Slot                 uint64 `db:"slot"`
		Authority            string `db:"authority"`
		Blockhash            string `db:"blockhash"`
		LamportsPerSignature uint64 `db:"lamports_per_signature"`
		State                string `db:"state"`
	}

	testCases := []struct {
		name     string
		data     NonceRow
		expected NonceRow
	}{
		{
			name: "initialize the data",
			data: NonceRow{
				"address", 1, "owner", "blockhash", 5000, "initialized",
			},
			expected: NonceRow{
				"address", 1, "owner", "blockhash", 5000, "initialized",
			},
		},
		{
			name: "update with lower slot",
			data: NonceRow{
				"address", 0, "pre_owner", "blockhash", 5000, "initialized",
			},
			expected: NonceRow{
				"address", 1, "owner", "blockhash", 5000, "initialized",
			},
		},
		{
			name: "update with same slot",
			data: NonceRow{
				"address", 1, "same_owner", "blockhash", 5000, "initialized",
			},
			expected: NonceRow{
				"address", 1, "same_owner", "blockhash", 5000, "initialized",
			},
		},
		{
			name: "update with higher slot",
			data: NonceRow{
				"address", 2, "new_owner", "blockhash", 5000, "initialized",
			},
			expected: NonceRow{
				"address", 2, "new_owner", "blockhash", 5000, "initialized",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveNonce(
				tc.data.Address,
				tc.data.Slot,
				tc.data.Authority,
				tc.data.Blockhash,
				tc.data.LamportsPerSignature,
				tc.data.State,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []NonceRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM nonce")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
