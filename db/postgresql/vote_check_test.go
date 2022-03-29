package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestCheckValidatorLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckValidatorLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveValidator(
		dbtypes.NewVoteAccountRow(
			"address", 1, "node", "voter", "withdrawer", 100,
		),
	)
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckValidatorLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckValidatorLatest("address", 2)
	suite.Require().True(isLatest)
}
