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
	err = suite.database.SaveTokenDelegation(dbtypes.NewTokenDelegationRow("source", "dest", 0, 1))
	suite.Require().NoError(err)

	err = migration.Up(suite.database)
	suite.Require().NoError(err)
}
