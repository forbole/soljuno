package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestSaveNonceAccount() {
	testCases := []struct {
		name     string
		data     dbtypes.NonceAccountRow
		expected dbtypes.NonceAccountRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewNonceAccountRow(
				"address", 1, "owner", "blockhash", 5000,
			),
			expected: dbtypes.NewNonceAccountRow(
				"address", 1, "owner", "blockhash", 5000,
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewNonceAccountRow(
				"address", 0, "pre_owner", "blockhash", 5000,
			),
			expected: dbtypes.NewNonceAccountRow(
				"address", 1, "owner", "blockhash", 5000,
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewNonceAccountRow(
				"address", 1, "same_owner", "blockhash", 5000,
			),
			expected: dbtypes.NewNonceAccountRow(
				"address", 1, "same_owner", "blockhash", 5000,
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewNonceAccountRow(
				"address", 2, "new_owner", "blockhash", 5000,
			),
			expected: dbtypes.NewNonceAccountRow(
				"address", 2, "new_owner", "blockhash", 5000,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveNonceAccount(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.NonceAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM nonce_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestDeleteNonceAccount() {
	err := suite.database.SaveNonceAccount(
		dbtypes.NewNonceAccountRow(
			"address",
			0,
			"owner",
			"hash",
			20,
		),
	)
	suite.Require().NoError(err)
	rows := []dbtypes.NonceAccountRow{}

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM nonce_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	rows = nil

	err = suite.database.DeleteNonceAccount("address")
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM nonce_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}
