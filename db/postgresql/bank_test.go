package postgresql_test

import "github.com/forbole/soljuno/solana/client/types"

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
		balances []types.TransactionTokenBalance
		expected BalanceRow
	}{
		{
			name:     "initialize the data",
			slot:     1,
			accounts: []string{"address"},
			balances: []types.TransactionTokenBalance{
				{
					AccountIndex: 0,
					Mint:         "mint",
					UiTokenAmount: types.UiTokenAmount{
						Amount: "1",
					},
				},
			},
			expected: BalanceRow{
				"address", 1, 1,
			},
		},
		{
			name:     "update with lower slot",
			slot:     0,
			accounts: []string{"address"},
			balances: []types.TransactionTokenBalance{
				{
					AccountIndex: 0,
					Mint:         "mint",
					UiTokenAmount: types.UiTokenAmount{
						Amount: "10",
					},
				},
			},
			expected: BalanceRow{
				"address", 1, 1,
			},
		},
		{
			name:     "update with same slot",
			slot:     1,
			accounts: []string{"address"},
			balances: []types.TransactionTokenBalance{
				{
					AccountIndex: 0,
					Mint:         "mint",
					UiTokenAmount: types.UiTokenAmount{
						Amount: "100",
					},
				},
			},
			expected: BalanceRow{
				"address", 1, 100,
			},
		},
		{
			name:     "update with higher slot",
			slot:     2,
			accounts: []string{"address"},
			balances: []types.TransactionTokenBalance{
				{
					AccountIndex: 0,
					Mint:         "mint",
					UiTokenAmount: types.UiTokenAmount{
						Amount: "1000",
					},
				},
			},
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
