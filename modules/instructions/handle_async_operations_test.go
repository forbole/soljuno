package instructions_test

import (
	"fmt"

	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *ModuleTestSuit) TestHandleBuffer() {
	buffer := make(chan dbtypes.InstructionRow, 1)
	suite.module.WithBuffer(buffer)
	buffer <- dbtypes.NewInstructionRow("sig", 0, 0, 0, "program", []string{"account"}, "", "unknown", []byte("null"))
	suite.Require().Len(buffer, 1)
	suite.module.HandleBuffer()
	suite.Require().Len(buffer, 0)
}

func (suite *ModuleTestSuit) TestHandleAsyncError() {
	buffer := make(chan dbtypes.InstructionRow, 1)
	suite.module.WithBuffer(buffer)
	suite.module.HandleAsyncError(
		fmt.Errorf("error"),
		[]dbtypes.InstructionRow{
			dbtypes.NewInstructionRow("sig", 0, 0, 0, "program", []string{"account"}, "", "unknown", []byte("null")),
		},
	)
	suite.Require().Len(buffer, 1)
	suite.Require().Equal(
		dbtypes.NewInstructionRow("sig", 0, 0, 0, "program", []string{"account"}, "", "unknown", []byte("null")),
		<-buffer)
}
