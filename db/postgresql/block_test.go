package postgresql_test

import (
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestSaveBlock() {
	// save a block
	block := dbtypes.NewBlockRow(1, 1, "hash", "leader", time.Date(2022, 3, 29, 0, 0, 0, 0, time.UTC).UTC(), 0)
	err := suite.database.SaveBlock(block)
	suite.Require().NoError(err)

	// verify the data
	rows := []dbtypes.BlockRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM block")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(block.Equal(rows[0]))
	rows = nil

	// Save a duplicated block
	duplicatedBlock := dbtypes.NewBlockRow(1, 1, "hash2", "leader2", time.Date(2022, 3, 29, 0, 0, 0, 0, time.UTC).UTC(), 0)
	err = suite.database.SaveBlock(duplicatedBlock)
	suite.Require().NoError(err)

	// verify the data
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM block")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(block.Equal(rows[0]))
	rows = nil
}

func (suite *DbTestSuite) TestHasBlock() {
	// save a block
	block := dbtypes.NewBlockRow(1, 1, "hash", "leader", time.Date(2022, 3, 29, 0, 0, 0, 0, time.UTC).UTC(), 0)
	err := suite.database.SaveBlock(block)
	suite.Require().NoError(err)

	suite.Require().False(suite.database.HasBlock(0))
	suite.Require().True(suite.database.HasBlock(1))
}
