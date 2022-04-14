package txs_test

import (
	"fmt"
	"time"

	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuit) TestName() {
	suite.Require().Equal("txs", suite.module.Name())
}

func (suite *ModuleTestSuit) TestHandleBlock() {
	buffer := make(chan types.Block, 1)
	suite.module.WithBuffer(buffer)
	err := suite.module.HandleBlock(types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}))
	suite.Require().Len(buffer, 1)
	suite.Require().Equal(types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}), <-buffer)
	suite.Require().NoError(err)
}

func (suite *ModuleTestSuit) TestHandleBuffer() {
	buffer := make(chan types.Block, 1)
	suite.module.WithBuffer(buffer)
	buffer <- types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{})
	suite.Require().Len(buffer, 1)
	suite.module.HandleBuffer()
	suite.Require().Len(buffer, 0)
}

func (suite *ModuleTestSuit) TestHandleAsyncError() {
	buffer := make(chan types.Block, 1)
	suite.module.WithBuffer(buffer)
	suite.module.HandleAsyncError(
		fmt.Errorf("error"),
		types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}),
	)
	suite.Require().Len(buffer, 1)
	suite.Require().Equal(types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}), <-buffer)
}
