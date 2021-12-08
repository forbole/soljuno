package migration_test

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/forbole/soljuno/db"
	database "github.com/forbole/soljuno/db/postgresql"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
	"github.com/stretchr/testify/suite"
)

func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(MigrateTestSuite))
}

type MigrateTestSuite struct {
	suite.Suite
	database *database.Database
}

func (suite *MigrateTestSuite) SetupTest() {
	dbCfg := types.NewDatabaseConfig(
		"soljuno",
		"localhost",
		5433,
		"soljuno",
		"password",
		"",
		"public",
		-1,
		-1,
	)
	db, err := database.Builder(db.NewContext(dbCfg, logging.DefaultLogger()))
	suite.Require().NoError(err)

	solDb, ok := (db).(*database.Database)
	suite.Require().True(ok)

	// Delete the public schema
	_, err = db.Exec(`DROP SCHEMA public CASCADE;`)
	suite.Require().NoError(err)

	// Re-create the schema
	_, err = db.Exec(`CREATE SCHEMA public;`)
	suite.Require().NoError(err)

	dirPath := path.Join("../", "schema")
	dir, err := ioutil.ReadDir(dirPath)
	suite.Require().NoError(err)

	for _, fileInfo := range dir {
		file, err := ioutil.ReadFile(filepath.Join(dirPath, fileInfo.Name()))
		suite.Require().NoError(err)

		commentsRegExp := regexp.MustCompile(`/\*.*\*/`)
		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err := db.Exec(commentsRegExp.ReplaceAllString(request, ""))
			suite.Require().NoError(err)
		}
	}

	suite.database = solDb

}

func (suite *MigrateTestSuite) TearDownSuite() {
	suite.database.Close()
}
