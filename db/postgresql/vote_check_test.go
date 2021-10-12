package postgresql_test

func (suite *DbTestSuite) TestCheckVoteAccountLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckVoteAccountLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveVoteAccount("address", 1, "node", "voter", "withdrawer", 100)
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckVoteAccountLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckVoteAccountLatest("address", 2)
	suite.Require().True(isLatest)
}
