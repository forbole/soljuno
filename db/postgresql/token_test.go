package postgresql_test

import (
	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestSaveToken() {

	testCases := []struct {
		name     string
		data     dbtypes.TokenRow
		expected dbtypes.TokenRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewTokenRow(
				"mint", 1, 9, "owner", "freeze",
			),
			expected: dbtypes.NewTokenRow(
				"mint", 1, 9, "owner", "freeze",
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewTokenRow(
				"mint", 0, 9, "pre_owner", "freeze",
			),
			expected: dbtypes.NewTokenRow(
				"mint", 1, 9, "owner", "freeze",
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewTokenRow(
				"mint", 1, 9, "new_owner", "freeze",
			),
			expected: dbtypes.NewTokenRow(
				"mint", 1, 9, "new_owner", "freeze",
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewTokenRow(
				"mint", 2, 9, "new_owner", "new_freeze",
			),
			expected: dbtypes.NewTokenRow(
				"mint", 2, 9, "new_owner", "new_freeze",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveToken(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.TokenRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveTokenAccount() {

	testCases := []struct {
		name     string
		data     dbtypes.TokenAccountRow
		expected dbtypes.TokenAccountRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewTokenAccountRow(
				"mint", 1, "mint", "owner",
			),
			expected: dbtypes.NewTokenAccountRow(
				"mint", 1, "mint", "owner",
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewTokenAccountRow(
				"mint", 0, "mint", "pre_owner",
			),
			expected: dbtypes.NewTokenAccountRow(
				"mint", 1, "mint", "owner",
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewTokenAccountRow(
				"mint", 1, "mint", "new_owner",
			),
			expected: dbtypes.NewTokenAccountRow(
				"mint", 1, "mint", "new_owner",
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewTokenAccountRow(
				"mint", 2, "mint", "new_owner",
			),
			expected: dbtypes.NewTokenAccountRow(
				"mint", 2, "mint", "new_owner",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveTokenAccount(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.TokenAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestDeleteTokenAccount() {
	err := suite.database.SaveTokenAccount(dbtypes.NewTokenAccountRow("address", 0, "mint", "owner"))
	suite.Require().NoError(err)
	rows := []dbtypes.TokenAccountRow{}

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	rows = nil

	err = suite.database.DeleteTokenAccount("address")
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}

func (suite *DbTestSuite) TestSaveMultisig() {

	testCases := []struct {
		name     string
		data     dbtypes.MultisigRow
		expected dbtypes.MultisigRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewMultisigRow(
				"mint", 1, []string{"signer1", "signer2"}, 1,
			),
			expected: dbtypes.NewMultisigRow(
				"mint", 1, []string{"signer1", "signer2"}, 1,
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewMultisigRow(
				"mint", 0, []string{"signer1"}, 1,
			),
			expected: dbtypes.NewMultisigRow(
				"mint", 1, []string{"signer1", "signer2"}, 1,
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewMultisigRow(
				"mint", 1, []string{"signer1"}, 1,
			),
			expected: dbtypes.NewMultisigRow(
				"mint", 1, []string{"signer1"}, 1,
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewMultisigRow(
				"mint", 2, []string{"signer1", "signer2"}, 1,
			),
			expected: dbtypes.NewMultisigRow(
				"mint", 2, []string{"signer1", "signer2"}, 1,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveMultisig(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.MultisigRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM multisig")

			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveDelegation() {
	testCases := []struct {
		name     string
		data     dbtypes.TokenDelegationRow
		expected dbtypes.TokenDelegationRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewTokenDelegationRow(
				"source_address", "delegate_address", 1, 1,
			),
			expected: dbtypes.NewTokenDelegationRow(
				"source_address", "delegate_address", 1, 1,
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewTokenDelegationRow(
				"source_address", "delegate_address", 0, 10,
			),
			expected: dbtypes.NewTokenDelegationRow(
				"source_address", "delegate_address", 1, 1,
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewTokenDelegationRow(
				"source_address", "delegate_address", 1, 10,
			),
			expected: dbtypes.NewTokenDelegationRow(
				"source_address", "delegate_address", 1, 10,
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewTokenDelegationRow(
				"source_address", "delegate_address", 1, 100,
			),
			expected: dbtypes.NewTokenDelegationRow(
				"source_address", "delegate_address", 1, 100,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveTokenDelegation(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.TokenDelegationRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_delegation")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestDeleteTokenDelegation() {
	err := suite.database.SaveTokenDelegation(
		dbtypes.NewTokenDelegationRow(
			"address",
			"dest",
			0,
			100,
		),
	)
	suite.Require().NoError(err)
	rows := []struct {
		SourceAddress   string `db:"source_address"`
		DelegateAddress string `db:"delegate_address"`
		Slot            uint64 `db:"slot"`
		Amount          uint64 `db:"amount"`
	}{}

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_delegation")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	rows = nil

	err = suite.database.DeleteTokenDelegation("address")
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_delegation")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}

func (suite *DbTestSuite) SaveTokenSupply() {
	testCases := []struct {
		name     string
		data     dbtypes.TokenSupplyRow
		expected dbtypes.TokenSupplyRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewTokenSupplyRow(
				"mint", 1, 1,
			),
			expected: dbtypes.NewTokenSupplyRow(
				"mint", 1, 1,
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewTokenSupplyRow(
				"mint", 0, 10,
			),
			expected: dbtypes.NewTokenSupplyRow(
				"mint", 1, 1,
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewTokenSupplyRow(
				"mint", 1, 100,
			),
			expected: dbtypes.NewTokenSupplyRow(
				"mint", 1, 100,
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewTokenSupplyRow(
				"mint", 2, 1000,
			),
			expected: dbtypes.NewTokenSupplyRow(
				"mint", 2, 1000,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveTokenSupply(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.TokenSupplyRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM token_supply")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
