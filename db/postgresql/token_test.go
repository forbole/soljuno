package postgresql_test

import "github.com/lib/pq"

func (suite *DbTestSuite) TestSaveToken() {
	type TokenRow struct {
		Mint            string `db:"mint"`
		Slot            uint64 `db:"slot"`
		Decimals        uint8  `db:"decimals"`
		MintAuthority   string `db:"mint_authority"`
		FreezeAuthority string `db:"freeze_authority"`
	}

	testCases := []struct {
		name     string
		data     TokenRow
		expected TokenRow
	}{
		{
			name: "initialize the data",
			data: TokenRow{
				"mint", 1, 9, "owner", "freeze",
			},
			expected: TokenRow{
				"mint", 1, 9, "owner", "freeze",
			},
		},
		{
			name: "update with lower slot",
			data: TokenRow{
				"mint", 0, 9, "pre_owner", "freeze",
			},
			expected: TokenRow{
				"mint", 1, 9, "owner", "freeze",
			},
		},
		{
			name: "update with same slot",
			data: TokenRow{
				"mint", 1, 9, "new_owner", "freeze",
			},
			expected: TokenRow{
				"mint", 1, 9, "new_owner", "freeze",
			},
		},
		{
			name: "update with higher slot",
			data: TokenRow{
				"mint", 2, 9, "new_owner", "new_freeze",
			},
			expected: TokenRow{
				"mint", 2, 9, "new_owner", "new_freeze",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveToken(
				tc.data.Mint,
				tc.data.Slot,
				tc.data.Decimals,
				tc.data.MintAuthority,
				tc.data.FreezeAuthority,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []TokenRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveTokenAccount() {
	type TokenAccountRow struct {
		Address string `db:"address"`
		Slot    uint64 `db:"slot"`
		Mint    string `db:"mint"`
		Owner   string `db:"owner"`
		State   string `db:"state"`
	}

	testCases := []struct {
		name     string
		data     TokenAccountRow
		expected TokenAccountRow
	}{
		{
			name: "initialize the data",
			data: TokenAccountRow{
				"mint", 1, "mint", "owner", "state",
			},
			expected: TokenAccountRow{
				"mint", 1, "mint", "owner", "state",
			},
		},
		{
			name: "update with lower slot",
			data: TokenAccountRow{
				"mint", 0, "mint", "pre_owner", "state",
			},
			expected: TokenAccountRow{
				"mint", 1, "mint", "owner", "state",
			},
		},
		{
			name: "update with same slot",
			data: TokenAccountRow{
				"mint", 1, "mint", "new_owner", "state",
			},
			expected: TokenAccountRow{
				"mint", 1, "mint", "new_owner", "state",
			},
		},
		{
			name: "update with higher slot",
			data: TokenAccountRow{
				"mint", 2, "mint", "new_owner", "new_state",
			},
			expected: TokenAccountRow{
				"mint", 2, "mint", "new_owner", "new_state",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveTokenAccount(
				tc.data.Address,
				tc.data.Slot,
				tc.data.Mint,
				tc.data.Owner,
				tc.data.State,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []TokenAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveMultisig() {
	type MultisigRow struct {
		Address string         `db:"address"`
		Slot    uint64         `db:"slot"`
		Signers pq.StringArray `db:"signers"`
		M       uint8          `db:"m"`
	}

	testCases := []struct {
		name     string
		data     MultisigRow
		expected MultisigRow
	}{
		{
			name: "initialize the data",
			data: MultisigRow{
				"mint", 1, []string{"signer1", "signer2"}, 1,
			},
			expected: MultisigRow{
				"mint", 1, []string{"signer1", "signer2"}, 1,
			},
		},
		{
			name: "update with lower slot",
			data: MultisigRow{
				"mint", 0, []string{"signer1"}, 1,
			},
			expected: MultisigRow{
				"mint", 1, []string{"signer1", "signer2"}, 1,
			},
		},
		{
			name: "update with same slot",
			data: MultisigRow{
				"mint", 1, []string{"signer1"}, 1,
			},
			expected: MultisigRow{
				"mint", 1, []string{"signer1"}, 1,
			},
		},
		{
			name: "update with higher slot",
			data: MultisigRow{
				"mint", 2, []string{"signer1", "signer2"}, 1,
			},
			expected: MultisigRow{
				"mint", 2, []string{"signer1", "signer2"}, 1,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveMultisig(
				tc.data.Address,
				tc.data.Slot,
				tc.data.Signers,
				tc.data.M,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []MultisigRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM multisig")

			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveDelegate() {
	type DelegateRow struct {
		SourceAddress   string `db:"source_address"`
		DelegateAddress string `db:"delegate_address"`
		Slot            uint64 `db:"slot"`
		Amount          uint64 `db:"amount"`
	}

	testCases := []struct {
		name     string
		data     DelegateRow
		expected DelegateRow
	}{
		{
			name: "initialize the data",
			data: DelegateRow{
				"source_address", "delegate_address", 1, 1,
			},
			expected: DelegateRow{
				"source_address", "delegate_address", 1, 1,
			},
		},
		{
			name: "update with lower slot",
			data: DelegateRow{
				"source_address", "delegate_address", 0, 10,
			},
			expected: DelegateRow{
				"source_address", "delegate_address", 1, 1,
			},
		},
		{
			name: "update with same slot",
			data: DelegateRow{
				"source_address", "delegate_address", 1, 10,
			},
			expected: DelegateRow{
				"source_address", "delegate_address", 1, 10,
			},
		},
		{
			name: "update with higher slot",
			data: DelegateRow{
				"source_address", "delegate_address", 1, 100,
			},
			expected: DelegateRow{
				"source_address", "delegate_address", 1, 100,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveDelegate(
				tc.data.SourceAddress,
				tc.data.DelegateAddress,
				tc.data.Slot,
				tc.data.Amount,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []DelegateRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_delegate")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) SaveTokenSupply() {
	type TokenSupplyRow struct {
		Mint   string `db:"mint"`
		Slot   uint64 `db:"slot"`
		Supply uint64 `db:"supply"`
	}

	testCases := []struct {
		name     string
		data     TokenSupplyRow
		expected TokenSupplyRow
	}{
		{
			name: "initialize the data",
			data: TokenSupplyRow{
				"mint", 1, 1,
			},
			expected: TokenSupplyRow{
				"mint", 1, 1,
			},
		},
		{
			name: "update with lower slot",
			data: TokenSupplyRow{
				"mint", 0, 10,
			},
			expected: TokenSupplyRow{
				"mint", 1, 1,
			},
		},
		{
			name: "update with same slot",
			data: TokenSupplyRow{
				"mint", 1, 100,
			},
			expected: TokenSupplyRow{
				"mint", 1, 100,
			},
		},
		{
			name: "update with higher slot",
			data: TokenSupplyRow{
				"mint", 2, 1000,
			},
			expected: TokenSupplyRow{
				"mint", 2, 1000,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveTokenSupply(
				tc.data.Mint,
				tc.data.Slot,
				tc.data.Supply,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []TokenSupplyRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_supply")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
