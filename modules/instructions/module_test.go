package instructions_test

import (
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
	solana "github.com/forbole/soljuno/solana/types"
	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuit) TestName() {
	suite.Require().Equal("instructions", suite.module.Name())
}

func (suite *ModuleTestSuit) TestHandleBlock() {
	err := suite.module.HandleBlock(types.NewBlock(0, 0, "hash", "leader", time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC), []types.Tx{}))
	suite.Require().NoError(err)
}

func (suite *ModuleTestSuit) TestHandleTxs() {
	buffer := make(chan dbtypes.InstructionRow, 10)
	suite.module.WithBuffer(buffer)
	err := suite.module.HandleInstruction(
		types.NewInstruction("sig", 0, 0, 0, "program", []string{"account"}, "", solana.NewParsedInstruction("unknown", nil)),
		types.NewTx("hash", 0, nil, 0, nil, nil, nil, nil, nil),
	)
	suite.Require().Len(buffer, 1)
	suite.Require().Equal(dbtypes.NewInstructionRow("sig", 0, 0, 0, "program", []string{"account"}, "", "unknown", []byte("null")), <-buffer)
	suite.Require().NoError(err)
}

func (suite *ModuleTestSuit) TestPrune() {
	err := suite.module.Prune(0)
	suite.Require().NoError(err)
}
