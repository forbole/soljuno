package vote_test

import (
	"github.com/forbole/soljuno/modules/vote"
	voteProgram "github.com/forbole/soljuno/solana/program/vote"
	solanatypes "github.com/forbole/soljuno/solana/types"
	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) Test_HandleInstruction() {
	testCases := []struct {
		name string
		data types.Instruction
	}{
		{
			name: "initialize instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"vote",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"initialize",
					voteProgram.ParsedInitializeAccount{},
				),
			),
		},
		{
			name: "authorize instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"vote",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"authorize",
					voteProgram.ParsedAuthorize{},
				),
			),
		},
		{
			name: "delegate instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"vote",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"withdraw",
					voteProgram.ParsedWithdraw{},
				),
			),
		},
		{
			name: "split instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"vote",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"updateValidatorIdentity",
					voteProgram.ParsedUpdateValidatorIdentity{},
				),
			),
		},
		{
			name: "withdraw instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"vote",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"updateCommission",
					voteProgram.ParsedUpdateCommission{},
				),
			),
		},
		{
			name: "deactivate instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"vote",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"authorizeChecked",
					voteProgram.ParsedAuthorizeChecked{},
				),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := vote.HandleInstruction(tc.data, suite.db, suite.client)
			suite.Require().NoError(err)
		})
	}
}
