package postgresql_test

func (suite *DbTestSuite) TestCheckTokenLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckTokenLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveToken("address", 1, 9, "mint_authority", "freeze_authority")
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckTokenLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckTokenLatest("address", 2)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckTokenAccountLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckTokenAccountLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveTokenAccount("address", 1, "mint", "owner")
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckTokenAccountLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckTokenAccountLatest("address", 2)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckMultisigLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckMultisigLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveMultisig("address", 1, []string{"signer1"}, 1)
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckMultisigLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckMultisigLatest("address", 2)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckTokenDelegateLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckTokenDelegateLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveTokenDelegation("source", "destination", 1, 100)
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckTokenDelegateLatest("source", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckTokenDelegateLatest("source", 2)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckTokenSupplyLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckTokenSupplyLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveTokenSupply("address", 1, 100)
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckTokenSupplyLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckTokenSupplyLatest("address", 2)
	suite.Require().True(isLatest)
}
