package migration_test

import (
	"github.com/forbole/soljuno/db/migration"
	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *MigrateTestSuite) TestMigrate() {
	err := migration.Down(suite.database)
	suite.Require().NoError(err)

	// custom query
	err = suite.database.SaveTokenAccount(dbtypes.NewTokenAccountRow("source", 0, "mint", "owner"))
	suite.Require().NoError(err)
	err = suite.database.SaveTokenAccount(dbtypes.NewTokenAccountRow("dest", 0, "mint", "owner"))
	suite.Require().NoError(err)
	err = suite.database.SaveTokenDelegation(dbtypes.NewTokenDelegationRow("source", "dest", 0, 1))
	suite.Require().NoError(err)
	err = suite.database.SaveTokenDelegation(dbtypes.NewTokenDelegationRow("dest", "dest1", 0, 1))
	suite.Require().NoError(err)

	var count []int
	err = suite.database.Sqlx.Select(&count, "SELECT COUNT(*) FROM token_delegation")
	suite.Require().NoError(err)
	suite.Require().Equal([]int{2}, count)
	count = nil

	err = migration.Up(suite.database)
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&count, "SELECT COUNT(*) FROM token_delegation")
	suite.Require().NoError(err)
	suite.Require().Equal([]int{1}, count)
}
