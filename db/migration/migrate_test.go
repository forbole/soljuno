package migration_test

import (
	"github.com/forbole/soljuno/db/migration"
	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *MigrateTestSuite) TestMigrate() {
	err := migration.Down(suite.database)
	suite.Require().NoError(err)

	suite.database.Exec(`
INSERT INTO validator_config
    (address, slot, owner, name, keybase_username, website, details)
VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		"address",
		0,
		"owner",
		"name",
		"keybase_username",
		"website",
		"details",
	)

	err = migration.Up(suite.database, suite.keybaseClient)
	suite.Require().NoError(err)
	expected := dbtypes.NewValidatorConfigRow(
		"address",
		0,
		"owner",
		dbtypes.NewParsedValidatorConfig(
			"name", "keybase_username", "website", "details",
		),
		"keybase_username",
	)

	rows := []dbtypes.ValidatorConfigRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM validator_config")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().Equal(expected, rows[0])
}
