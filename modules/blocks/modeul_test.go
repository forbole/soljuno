package blocks_test

import (
	"time"

	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) TestModule_Name() {
	suite.Require().Equal("blocks", suite.module.Name())
}

func (suite *ModuleTestSuite) TestModule_HandleBlock() {
	err := suite.module.HandleBlock(types.NewBlock(0, 0, "hash", "leader", nil, time.Now(), []types.Tx{}))
	suite.Require().NoError(err)
}
