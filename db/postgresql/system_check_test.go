package postgresql_test

func (suite *DbTestSuite) TestCheckNonceAccountLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckNonceAccountLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveNonceAccount("address", 1, "new_owner", "blockhash", 5000, "initialized")
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckNonceAccountLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckNonceAccountLatest("address", 2)
	suite.Require().True(isLatest)
}
