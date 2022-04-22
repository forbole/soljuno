package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestCheckNonceAccountLatest() {
	// empty rows returns false
	isLatest := suite.database.CheckNonceAccountLatest("address", 1)
	suite.Require().False(isLatest)

	err := suite.database.SaveNonceAccount(
		dbtypes.NewNonceAccountRow(
			"address", 1, "new_owner", "blockhash", 5000,
		),
	)
	suite.Require().NoError(err)

	// db slot is older returns false
	isLatest = suite.database.CheckNonceAccountLatest("address", 2)
	suite.Require().False(isLatest)

	// db slot is latest returns true
	isLatest = suite.database.CheckNonceAccountLatest("address", 1)
	suite.Require().True(isLatest)
}
