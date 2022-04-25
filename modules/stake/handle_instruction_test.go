package stake_test

import (
	"github.com/forbole/soljuno/modules/stake"
	stakeProgram "github.com/forbole/soljuno/solana/program/stake"
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
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"initialize",
					stakeProgram.ParsedInitialize{},
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
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"authorize",
					stakeProgram.ParsedAuthorize{},
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
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"delegate",
					stakeProgram.ParsedDelegateStake{},
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
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"split",
					stakeProgram.ParsedSplit{},
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
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"withdraw",
					stakeProgram.ParsedWithdraw{},
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
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"deactivate",
					stakeProgram.ParsedDeactivate{},
				),
			),
		},
		{
			name: "setLockup instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"setLockup",
					stakeProgram.ParsedSetLockup{},
				),
			),
		},
		{
			name: "merge instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"merge",
					stakeProgram.ParsedMerge{},
				),
			),
		},
		{
			name: "authorizeWithSeed instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"authorizeWithSeed",
					stakeProgram.ParsedAuthorizeWithSeed{},
				),
			),
		},
		{
			name: "initializeChecked instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"initializeChecked",
					stakeProgram.ParsedInitializeChecked{},
				),
			),
		},
		{
			name: "authorizeChecked instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"authorizeChecked",
					stakeProgram.ParsedAuthorizeChecked{},
				),
			),
		},
		{
			name: "authorizeCheckedWithSeed instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"authorizeCheckedWithSeed",
					stakeProgram.ParsedAuthorizeCheckedWithSeed{},
				),
			),
		},
		{
			name: "setLockupChecked instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"setLockupChecked",
					stakeProgram.ParsedSetLockupChecked{},
				),
			),
		},
		{
			name: "deactivateDelinquent instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"stake",
				[]string{},
				"",
				solanatypes.NewParsedInstruction(
					"deactivateDelinquent",
					stakeProgram.ParsedDeactivateDelinquent{},
				),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := stake.HandleInstruction(tc.data, suite.db, suite.client)
			suite.Require().NoError(err)
		})
	}
}
