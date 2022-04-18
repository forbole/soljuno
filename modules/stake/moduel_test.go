package stake_test

import (
	"fmt"

	stakeProgram "github.com/forbole/soljuno/solana/program/stake"
	solanatypes "github.com/forbole/soljuno/solana/types"
	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) TestModule_Name() {
	suite.Require().Equal("stake", suite.module.Name())
}

func (suite *ModuleTestSuite) TestModule_HandleInstruction() {
	testCases := []struct {
		name        string
		tx          types.Tx
		instruction types.Instruction
		shouldErr   bool
	}{
		{
			name: "failed tx skip properly",
			tx:   types.NewTx("sig", 0, fmt.Errorf("failed"), 0, nil, nil, nil, nil, nil),
		},
		{
			name: "non stake instruction skips properly",
			instruction: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"unknown",
				nil,
				"",
				solanatypes.NewParsedInstruction(
					"initialize",
					nil,
				),
			),
		},
		{
			name: "fail to handle instruction return error",
			tx:   types.NewTx("sig", 0, nil, 0, nil, nil, nil, nil, nil),
			instruction: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				stakeProgram.ProgramID,
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"initialize",
					nil,
				),
			),
			shouldErr: true,
		},
		{
			name: "instruction works properly",
			tx:   types.NewTx("sig", 0, nil, 0, nil, nil, nil, nil, nil),
			instruction: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				stakeProgram.ProgramID,
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"initialize",
					stakeProgram.ParsedInitialize{},
				),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.module.HandleInstruction(tc.instruction, tc.tx)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
