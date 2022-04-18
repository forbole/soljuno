package system_test

import (
	"github.com/forbole/soljuno/modules/system"
	systemProgram "github.com/forbole/soljuno/solana/program/system"
	solanatypes "github.com/forbole/soljuno/solana/types"
	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) Test_HandleInstruction() {
	testCases := []struct {
		name string
		data types.Instruction
	}{
		{
			name: "advanceNonce instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"system",
				[]string{"nonce", "sysvar", "owner"},
				"",
				solanatypes.NewParsedInstruction(
					"advanceNonce",
					systemProgram.ParsedAdvanceNonceAccount{},
				),
			),
		},
		{
			name: "withdrawFromNonce instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"system",
				[]string{"nonce", "sysvar", "owner"},
				"",
				solanatypes.NewParsedInstruction(
					"withdrawFromNonce",
					systemProgram.ParsedWithdrawNonceAccount{},
				),
			),
		},
		{
			name: "initializeNonce instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"system",
				[]string{"nonce", "sysvar", "owner"},
				"",
				solanatypes.NewParsedInstruction(
					"initializeNonce",
					systemProgram.ParsedInitializeNonceAccount{},
				),
			),
		},
		{
			name: "authorizeNonce instruction works properly",
			data: types.NewInstruction(
				"sig",
				1,
				0,
				0,
				"system",
				[]string{"nonce", "sysvar", "owner"},
				"",
				solanatypes.NewParsedInstruction(
					"authorizeNonce",
					systemProgram.ParsedAuthorizeNonceAccount{},
				),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := system.HandleInstruction(tc.data, suite.db, suite.client)
			suite.Require().NoError(err)
		})
	}
}
