package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestCheckValidatorLatest() {
	// empty rows returns false
	isLatest := suite.database.CheckValidatorLatest("address", 1)
	suite.Require().False(isLatest)

	err := suite.database.SaveValidator(
		dbtypes.NewVoteAccountRow(
			"address", 1, "node", "voter", "withdrawer", 100,
		),
	)
	suite.Require().NoError(err)

	// db slot is older returns false
	isLatest = suite.database.CheckValidatorLatest("address", 2)
	suite.Require().False(isLatest)

	// db slot is latest returns true
	isLatest = suite.database.CheckValidatorLatest("address", 1)
	suite.Require().True(isLatest)
}
