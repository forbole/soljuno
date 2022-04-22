package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestCheckBufferAccountLatest() {
	// empty rows returns false
	isLatest := suite.database.CheckBufferAccountLatest("address", 1)
	suite.Require().False(isLatest)

	err := suite.database.SaveBufferAccount(
		dbtypes.NewBufferAccountRow("address", 1, "owner"),
	)
	suite.Require().NoError(err)

	// db slot is older returns false
	isLatest = suite.database.CheckBufferAccountLatest("address", 2)
	suite.Require().False(isLatest)

	// db slot is latest returns true
	isLatest = suite.database.CheckBufferAccountLatest("address", 1)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckProgramAccountLatest() {
	// empty rows returns false
	isLatest := suite.database.CheckBufferAccountLatest("address", 1)
	suite.Require().False(isLatest)

	err := suite.database.SaveProgramAccount(
		dbtypes.NewProgramAccountRow("address", 1, "program_data"),
	)
	suite.Require().NoError(err)

	// db slot is older returns false
	isLatest = suite.database.CheckProgramAccountLatest("address", 2)
	suite.Require().False(isLatest)

	// db slot is latest returns true
	isLatest = suite.database.CheckProgramAccountLatest("address", 1)
	suite.Require().True(isLatest)
}

func (suite *DbTestSuite) TestCheckProgramDataAccountLatest() {
	// empty rows returns false
	isLatest := suite.database.CheckProgramDataAccountLatest("address", 1)
	suite.Require().False(isLatest)

	err := suite.database.SaveProgramDataAccount(
		dbtypes.NewProgramDataAccountRow("address", 1, 1, "owner"),
	)
	suite.Require().NoError(err)

	// db slot is older returns false
	isLatest = suite.database.CheckProgramDataAccountLatest("address", 2)
	suite.Require().False(isLatest)

	// db slot is latest returns true
	isLatest = suite.database.CheckProgramDataAccountLatest("address", 1)
	suite.Require().True(isLatest)
}
