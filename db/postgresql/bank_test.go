package postgresql_test

import (
	"time"
)

func (suite *DbTestSuite) TestSaveAccountBalances() {
	type BalanceRow struct {
		Address string `db:"address"`
		Slot    uint64 `db:"slot"`
		Balance uint64 `db:"balance"`
	}

	testCases := []struct {
		name     string
		slot     uint64
		accounts []string
		balances []uint64
		expected BalanceRow
	}{
		{
			name:     "initialize the data",
			slot:     1,
			accounts: []string{"address"},
			balances: []uint64{1},
			expected: BalanceRow{
				"address", 1, 1,
			},
		},
		{
			name:     "update with lower slot",
			slot:     0,
			accounts: []string{"address"},
			balances: []uint64{100},
			expected: BalanceRow{
				"address", 1, 1,
			},
		},
		{
			name:     "update with same slot",
			slot:     1,
			accounts: []string{"address"},
			balances: []uint64{100},
			expected: BalanceRow{
				"address", 1, 100,
			},
		},
		{
			name:     "update with higher slot",
			slot:     2,
			accounts: []string{"address"},
			balances: []uint64{1000},
			expected: BalanceRow{
				"address", 2, 1000,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveAccountBalances(
				tc.slot,
				tc.accounts,
				tc.balances,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []BalanceRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM account_balance")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveAccountTokenBalances() {
	type BalanceRow struct {
		Address string `db:"address"`
		Slot    uint64 `db:"slot"`
		Balance uint64 `db:"balance"`
	}

	testCases := []struct {
		name     string
		slot     uint64
		accounts []string
		balances []uint64
		expected BalanceRow
	}{
		{
			name:     "initialize the data",
			slot:     1,
			accounts: []string{"address"},
			balances: []uint64{1},
			expected: BalanceRow{
				"address", 1, 1,
			},
		},
		{
			name:     "update with lower slot",
			slot:     0,
			accounts: []string{"address"},
			balances: []uint64{10},
			expected: BalanceRow{
				"address", 1, 1,
			},
		},
		{
			name:     "update with same slot",
			slot:     1,
			accounts: []string{"address"},
			balances: []uint64{100},
			expected: BalanceRow{
				"address", 1, 100,
			},
		},
		{
			name:     "update with higher slot",
			slot:     2,
			accounts: []string{"address"},
			balances: []uint64{1000},
			expected: BalanceRow{
				"address", 2, 1000,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveAccountTokenBalances(
				tc.slot,
				tc.accounts,
				tc.balances,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []BalanceRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_account_balance")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveAccountHistoryBalances() {
	type BalanceRow struct {
		Address   string    `db:"address"`
		Timestamp time.Time `db:"timestamp"`
		Balance   uint64    `db:"balance"`
	}

	testCases := []struct {
		name      string
		timestamp time.Time
		accounts  []string
		balances  []uint64
		expected  []BalanceRow
	}{
		{
			name:      "initialize the data",
			timestamp: time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC),
			accounts:  []string{"address"},
			balances:  []uint64{1},
			expected: []BalanceRow{
				{
					"address", time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC), 1,
				},
			},
		},
		{
			name:      "insert another data",
			timestamp: time.Date(2020, 10, 10, 16, 05, 00, 000, time.UTC),
			accounts:  []string{"address"},
			balances:  []uint64{100},
			expected: []BalanceRow{
				{
					"address", time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC), 1,
				},
				{
					"address", time.Date(2020, 10, 10, 16, 05, 00, 000, time.UTC), 100,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveAccountHistoryBalances(
				tc.timestamp,
				tc.accounts,
				tc.balances,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []BalanceRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM account_balance_history ORDER BY timestamp")
			suite.Require().NoError(err)
			suite.Require().Len(rows, len(tc.expected))
			for i, row := range rows {
				suite.Require().True(tc.expected[i].Address == row.Address)
				suite.Require().True(tc.expected[i].Timestamp.Equal(row.Timestamp))
				suite.Require().True(tc.expected[i].Balance == row.Balance)
			}
		})
	}
}

func (suite *DbTestSuite) TestSaveAccountHistoryTokenBalances() {
	type BalanceRow struct {
		Address   string    `db:"address"`
		Timestamp time.Time `db:"timestamp"`
		Balance   uint64    `db:"balance"`
	}

	testCases := []struct {
		name      string
		timestamp time.Time
		accounts  []string
		balances  []uint64
		expected  []BalanceRow
	}{
		{
			name:      "initialize the data",
			timestamp: time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC),
			accounts:  []string{"address"},
			balances:  []uint64{1},
			expected: []BalanceRow{
				{
					"address", time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC), 1,
				},
			},
		},
		{
			name:      "insert another data",
			timestamp: time.Date(2020, 10, 10, 16, 05, 00, 000, time.UTC),
			accounts:  []string{"address"},
			balances:  []uint64{10},
			expected: []BalanceRow{
				{
					"address", time.Date(2020, 10, 10, 15, 05, 00, 000, time.UTC), 1,
				},
				{
					"address", time.Date(2020, 10, 10, 16, 05, 00, 000, time.UTC), 10,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveAccountHistoryTokenBalances(
				tc.timestamp,
				tc.accounts,
				tc.balances,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []BalanceRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_account_balance_history ORDER BY timestamp")
			suite.Require().NoError(err)
			suite.Require().Len(rows, len(tc.expected))
			for i, row := range rows {
				suite.Require().True(tc.expected[i].Address == row.Address)
				suite.Require().True(tc.expected[i].Timestamp.Equal(row.Timestamp))
				suite.Require().True(tc.expected[i].Balance == row.Balance)
			}
		})
	}
}
