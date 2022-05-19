package fix_test

import (
	"time"

	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) TestModule_Name() {
	suite.Require().Equal("fix", suite.module.Name())
}

func (suite *ModuleTestSuite) TestModule_HandleBlock() {
	err := suite.module.HandleBlock(types.NewBlock(150, 0, "hash", "leader", nil, time.Now(), []types.Tx{}))
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(1), suite.module.SlotInterval)

	err = suite.module.HandleBlock(types.NewBlock(250, 0, "hash", "leader", nil, time.Now(), []types.Tx{}))
	suite.Require().NoError(err)
	suite.Require().Equal(uint64(2), suite.module.SlotInterval)
}
