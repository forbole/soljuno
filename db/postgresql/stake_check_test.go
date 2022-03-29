package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestCheckStakeAccountLatest() {
	// empty rows returns true
	isLatest := suite.database.CheckStakeAccountLatest("address", 1)
	suite.Require().True(isLatest)

	err := suite.database.SaveStakeAccount(
		dbtypes.NewStakeAccountRow("address", 1, "staker", "withdrawer"),
	)
	suite.Require().NoError(err)

	// older slot returns false
	isLatest = suite.database.CheckStakeAccountLatest("address", 1)
	suite.Require().False(isLatest)

	// latest slot returns true
	isLatest = suite.database.CheckStakeAccountLatest("address", 2)
	suite.Require().True(isLatest)
}
