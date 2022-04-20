package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestCheckTokenLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckTokenLatest("mint", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveToken(dbtypes.NewTokenRow("mint", 1, 9, "mint_authority", "freeze_authority"))
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckTokenLatest("mint", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckTokenLatest("mint", 2)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckTokenAccountLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckTokenAccountLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveTokenAccount(dbtypes.NewTokenAccountRow("address", 1, "mint", "owner"))
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

	err := suite.database.SaveMultisig(
		dbtypes.NewMultisigRow(
			"address", 1, []string{"signer1"}, 1,
		),
	)
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckMultisigLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckMultisigLatest("address", 2)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckTokenDelegateLatest() {
	err := suite.database.SaveTokenAccount(dbtypes.NewTokenAccountRow("source", 0, "mint", "owner"))
	suite.NoError(err)
	err = suite.database.SaveTokenAccount(dbtypes.NewTokenAccountRow("dest", 0, "mint", "owner"))
	suite.NoError(err)

	// empty rows returns true
	isLatest := suite.database.CheckTokenDelegateLatest("source", 1)
	suite.Require().True(isLatest)

	err = suite.database.SaveTokenDelegation(
		dbtypes.NewTokenDelegationRow(
			"source", "dest", 1, 100,
		),
	)
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
	isLatest := suite.database.CheckTokenSupplyLatest("mint", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveTokenSupply(dbtypes.NewTokenSupplyRow("mint", 1, 100))
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckTokenSupplyLatest("mint", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckTokenSupplyLatest("mint", 2)
	suite.Require().True(isLatest)
}
