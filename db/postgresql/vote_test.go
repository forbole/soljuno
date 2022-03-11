package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

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

	testCases := []struct {
		name     string
		data     []dbtypes.ValidatorStatusRow
		expected []dbtypes.ValidatorStatusRow
	}{
		{
			name: "initialize the data",
			data: []dbtypes.ValidatorStatusRow{
				{"address", 1, 100, 0, 0, true},
			},
			expected: []dbtypes.ValidatorStatusRow{
				{"address", 1, 100, 0, 0, true},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveValidatorStatuses(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.ValidatorStatusRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM validator_status")
			suite.Require().NoError(err)
			suite.Require().Len(rows, len(tc.expected))
			suite.Require().Equal(tc.expected, rows)
		})
	}
}

func (suite *DbTestSuite) TestSaveValidatorSkipRates() {

	testCases := []struct {
		name     string
		data     dbtypes.ValidatorSkipRateRow
		expected []dbtypes.ValidatorSkipRateRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewValidatorSkipRateRow("address", 1, 0.1, 10, 1),
			expected: []dbtypes.ValidatorSkipRateRow{dbtypes.NewValidatorSkipRateRow("address", 1, 0.1, 10, 1)},
		},
		{
			name:     "insert with lower epoch",
			data:     dbtypes.NewValidatorSkipRateRow("address", 0, 0.2, 10, 2),
			expected: []dbtypes.ValidatorSkipRateRow{dbtypes.NewValidatorSkipRateRow("address", 1, 0.1, 10, 1)},
		},
		{
			name:     "insert with same epoch",
			data:     dbtypes.NewValidatorSkipRateRow("address", 1, 0.2, 10, 2),
			expected: []dbtypes.ValidatorSkipRateRow{dbtypes.NewValidatorSkipRateRow("address", 1, 0.2, 10, 2)},
		},
		{
			name:     "insert with higher epoch",
			data:     dbtypes.NewValidatorSkipRateRow("address", 2, 0.3, 10, 3),
			expected: []dbtypes.ValidatorSkipRateRow{dbtypes.NewValidatorSkipRateRow("address", 2, 0.3, 10, 3)},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveValidatorSkipRates(
				[]dbtypes.ValidatorSkipRateRow{tc.data},
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.ValidatorSkipRateRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM validator_skip_rate")
			suite.Require().NoError(err)
			suite.Require().Len(rows, len(tc.expected))
			suite.Require().Equal(tc.expected, rows)
		})
	}
}

func (suite *DbTestSuite) TestSaveHistoryValidatorSkipRates() {

	testCases := []struct {
		name     string
		data     dbtypes.ValidatorSkipRateRow
		expected []dbtypes.ValidatorSkipRateRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewValidatorSkipRateRow("address", 1, 0.1, 10, 1),
			expected: []dbtypes.ValidatorSkipRateRow{dbtypes.NewValidatorSkipRateRow("address", 1, 0.1, 10, 1)},
		},
		{
			name:     "insert with same epoch",
			data:     dbtypes.NewValidatorSkipRateRow("address", 1, 0.2, 10, 2),
			expected: []dbtypes.ValidatorSkipRateRow{dbtypes.NewValidatorSkipRateRow("address", 1, 0.1, 10, 1)},
		},
		{
			name: "insert with higher epoch",
			data: dbtypes.NewValidatorSkipRateRow("address", 2, 0.3, 10, 3),
			expected: []dbtypes.ValidatorSkipRateRow{
				dbtypes.NewValidatorSkipRateRow("address", 1, 0.1, 10, 1),
				dbtypes.NewValidatorSkipRateRow("address", 2, 0.3, 10, 3),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveHistoryValidatorSkipRates(
				[]dbtypes.ValidatorSkipRateRow{tc.data},
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.ValidatorSkipRateRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM validator_skip_rate_history")
			suite.Require().NoError(err)
			suite.Require().Len(rows, len(tc.expected))
			suite.Require().Equal(tc.expected, rows)
		})
	}
}
