package system

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/parser"
	"github.com/forbole/soljuno/solana/types"
)

type Parser struct {
	parser.ProgramParser
}

func (Parser) Parse(accounts []string, data []byte) types.ParsedInstruction {
	decoder := bincode.NewDecoder()
	var id InstructionID
	decoder.Decode(data[:4], &id)
	switch id {
	case CreateAccount:
		var instruction CreateAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"createAccount",
			NewParsedCreateAccount(
				accounts[0],
				accounts[1],
				instruction,
			),
		)

	case Assign:
		var instruction AssignInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"assign",
			NewParsedAssign(
				accounts[0],
				instruction,
			),
		)

	case Transfer:
		var instruction TransferInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"transfer",
			NewParsedTransfer(
				accounts[0],
				accounts[1],
				instruction,
			),
		)

	case CreateAccountWithSeed:
		var instruction CreateAccountWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"createAccountWithSeed",
			NewParsedCreateAccountWithSeed(
				accounts[0],
				accounts[1],
				instruction,
			),
		)

	case AdvanceNonceAccount:
		return types.NewParsedInstruction(
			"advanceNonce",
			NewParsedAdvanceNonceAccount(
				accounts[0],
				accounts[1],
				accounts[2],
			),
		)

	case WithdrawNonceAccount:
		var instruction WithdrawNonceAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"withdrawFromNonce",
			NewParsedWithdrawNonceAccount(
				accounts[0],
				accounts[1],
				accounts[2],
				accounts[3],
				accounts[4],
				instruction,
			),
		)

	case InitializeNonceAccount:
		var instruction InitializeNonceAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initializeNonce",
			NewParsedInitializeNonceAccount(
				accounts[0],
				accounts[1],
				accounts[2],
				instruction,
			),
		)

	case AuthorizeNonceAccount:
		var instruction AuthorizeNonceAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"authorizeNonce",
			NewParsedAuthorizeNonceAccount(
				accounts[0],
				accounts[1],
				instruction,
			),
		)

	case Allocate:
		var instruction AllocateInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"allocate",
			NewParsedAllocate(
				accounts[0],
				instruction,
			),
		)

	case AllocateWithSeed:
		var instruction AllocateWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"allocateWithSeed",
			NewParsedAllocateWithSeed(
				accounts[0],
				instruction,
			),
		)

	case AssignWithSeed:
		var instruction AssignWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"assignWithSeed",
			NewParsedAssignWithSeed(
				accounts[0],
				instruction,
			),
		)

	case TransferWithSeed:
		var instruction TransferWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"transferWithSeed",
			NewParsedTransferWithSeed(
				accounts[0],
				accounts[1],
				accounts[2],
				instruction,
			),
		)
	}

	return nil
}
