package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestCheckStakeAccountLatest() {
	// empty rows returns false
	isLatest := suite.database.CheckStakeAccountLatest("address", 1)
	suite.Require().False(isLatest)

	err := suite.database.SaveStakeAccount(
		dbtypes.NewStakeAccountRow("address", 1, "staker", "withdrawer"),
	)
	suite.Require().NoError(err)

	// db slot is older returns false
	isLatest = suite.database.CheckStakeAccountLatest("address", 2)
	suite.Require().False(isLatest)

	// db slot is latest returns true
	isLatest = suite.database.CheckStakeAccountLatest("address", 1)
	suite.Require().True(isLatest)
}
