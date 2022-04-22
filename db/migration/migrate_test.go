package migration_test

import (
	"github.com/forbole/soljuno/db/migration"
)

func (suite *MigrateTestSuite) TestMigrate() {
	err := migration.Down(suite.database)
	suite.Require().NoError(err)

	err = migration.Up(suite.database)
	suite.Require().NoError(err)
}
