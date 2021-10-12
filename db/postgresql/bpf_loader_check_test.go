package postgresql_test

func (suite *DbTestSuite) TestCheckBufferAccountLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckBufferAccountLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveBufferAccount("address", 1, "owner", "initialized")
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckBufferAccountLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckBufferAccountLatest("address", 2)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckProgramAccountLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckBufferAccountLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveProgramAccount("address", 1, "program_data", "initialized")
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckProgramAccountLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckProgramAccountLatest("address", 2)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckProgramDataAccountLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckProgramDataAccountLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveProgramDataAccount("address", 1, 1, "owner", "initialized")
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckProgramDataAccountLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckProgramDataAccountLatest("address", 2)
	suite.Require().True(isLatest)
}
