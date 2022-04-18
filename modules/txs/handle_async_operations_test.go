package txs_test

import (
	"fmt"
	"time"

	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) TestHandleBuffer() {
	buffer := make(chan types.Block, 1)
	suite.module.WithBuffer(buffer)
	buffer <- types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{})
	suite.Require().Len(buffer, 1)
	suite.module.HandleBuffer()
	suite.Require().Len(buffer, 0)
}

func (suite *ModuleTestSuite) TestHandleAsyncError() {
	buffer := make(chan types.Block, 1)
	suite.module.WithBuffer(buffer)
	suite.module.HandleAsyncError(
		fmt.Errorf("error"),
		types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}),
	)
	suite.Require().Len(buffer, 1)
	suite.Require().Equal(types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}), <-buffer)
}
