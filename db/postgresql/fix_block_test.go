package postgresql_test

import (
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestGetMissingHeight() {
	height, err := suite.database.GetMissingHeight(1, 2)
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), height)

	// save blocks having 1 and 3 height
	err = suite.database.SaveBlock(dbtypes.NewBlockRow(1, 1, "hash1", "leader", time.Now(), 0))
	suite.Require().NoError(err)
	err = suite.database.SaveBlock(dbtypes.NewBlockRow(3, 3, "hash3", "leader", time.Now(), 0))
	suite.Require().NoError(err)

	height, err = suite.database.GetMissingHeight(1, 3)
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(2), height)
}

func (suite *DbTestSuite) TestGetMissingSlotRange() {
	start, end, err := suite.database.GetMissingSlotRange(2)
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(0), start)
	suite.Require().Equal(uint64(0), end)

	// save blocks having 1 and 3 height
	err = suite.database.SaveBlock(dbtypes.NewBlockRow(1, 1, "hash1", "leader", time.Now(), 0))
	suite.Require().NoError(err)
	err = suite.database.SaveBlock(dbtypes.NewBlockRow(3, 3, "hash3", "leader", time.Now(), 0))
	suite.Require().NoError(err)

	start, end, err = suite.database.GetMissingSlotRange(2)
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(2), start)
	suite.Require().Equal(uint64(2), end)
}
