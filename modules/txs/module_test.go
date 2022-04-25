package txs_test

import (
	"time"

	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) TestName() {
	suite.Require().Equal("txs", suite.module.Name())
}

func (suite *ModuleTestSuite) TestHandleBlock() {
	buffer := make(chan types.Block, 1)
	suite.module.WithBuffer(buffer)
	err := suite.module.HandleBlock(types.NewBlock(0, 0, "hash", "leader", nil, time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}))
	suite.Require().Len(buffer, 1)
	suite.Require().Equal(types.NewBlock(0, 0, "hash", "leader", nil, time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}), <-buffer)
	suite.Require().NoError(err)
}

func (suite *ModuleTestSuite) TestPrune() {
	err := suite.module.Prune(0)
	suite.Require().NoError(err)
}
