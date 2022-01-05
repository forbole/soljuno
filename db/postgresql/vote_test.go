package postgresql_test

func (suite *DbTestSuite) TestSaveValidator() {
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
			err := suite.database.SaveValidator(
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
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM validator")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveValidatorStatus() {
	type ValidatorStatusRow struct {
		Address        string `db:"address"`
		Slot           uint64 `db:"slot"`
		ActivatedStake uint64 `db:"activated_stake"`
		LastVote       uint64 `db:"last_vote"`
		RootSlot       uint64 `db:"root_slot"`
		Active         bool   `db:"active"`
	}

	testCases := []struct {
		name     string
		data     ValidatorStatusRow
		expected []ValidatorStatusRow
	}{
		{
			name: "initialize the data",
			data: ValidatorStatusRow{
				"address", 1, 100, 0, 0, true,
			},
			expected: []ValidatorStatusRow{
				{"address", 1, 100, 0, 0, true},
			},
		},
		{
			name: "insert with lower slot",
			data: ValidatorStatusRow{
				"address", 0, 1000, 0, 0, true,
			},
			expected: []ValidatorStatusRow{
				{"address", 1, 100, 0, 0, true},
			},
		},
		{
			name: "insert with same slot",
			data: ValidatorStatusRow{
				"address", 1, 1000, 0, 0, true,
			},
			expected: []ValidatorStatusRow{
				{"address", 1, 1000, 0, 0, true},
			},
		},
		{
			name: "insert with higher slot",
			data: ValidatorStatusRow{
				"address", 2, 2000, 0, 0, true,
			},
			expected: []ValidatorStatusRow{
				{"address", 2, 2000, 0, 0, true},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveValidatorStatus(
				tc.data.Address,
				tc.data.Slot,
				tc.data.ActivatedStake,
				tc.data.LastVote,
				tc.data.RootSlot,
				tc.data.Active,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []ValidatorStatusRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM validator_status")
			suite.Require().NoError(err)
			suite.Require().Len(rows, len(tc.expected))
			suite.Require().Equal(tc.expected, rows)
		})
	}
}
