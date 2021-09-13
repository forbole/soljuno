package postgresql_test

func (suite *DbTestSuite) TestSaveToken() {
	type TokenRow struct {
		Mint            string `db:"mint"`
		Slot            uint64 `db:"slot"`
		Decimals        uint8  `db:"decimals"`
		MintAuthority   string `db:"mint_authority"`
		FreezeAuthority string `db:"freeze_authority"`
	}

	// Save the data
	err := suite.database.SaveToken(
		"mint",
		0,
		9,
		"owner",
		"freeze",
	)
	suite.Require().NoError(err)

	// Verify the data
	var rows []TokenRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM token`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "no duplicated token rows should be inserted")

	expected := TokenRow{"mint", 0, 9, "owner", "freeze"}
	suite.Require().Equal(expected, rows[0])

}
