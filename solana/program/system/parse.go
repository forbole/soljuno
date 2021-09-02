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
				accounts,
				instruction,
			),
		)

	case Assign:
		var instruction AssignInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"assign",
			NewParsedAssign(
				accounts,
				instruction,
			),
		)

	case Transfer:
		var instruction TransferInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"transfer",
			NewParsedTransfer(
				accounts,
				instruction,
			),
		)

	case CreateAccountWithSeed:
		var instruction CreateAccountWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"createAccountWithSeed",
			NewParsedCreateAccountWithSeed(
				accounts,
				instruction,
			),
		)

	case AdvanceNonceAccount:
		return types.NewParsedInstruction(
			"advanceNonce",
			NewParsedAdvanceNonceAccount(
				accounts,
			),
		)

	case WithdrawNonceAccount:
		var instruction WithdrawNonceAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"withdrawFromNonce",
			NewParsedWithdrawNonceAccount(
				accounts,
				instruction,
			),
		)

	case InitializeNonceAccount:
		var instruction InitializeNonceAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initializeNonce",
			NewParsedInitializeNonceAccount(
				accounts,
				instruction,
			),
		)

	case AuthorizeNonceAccount:
		var instruction AuthorizeNonceAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"authorizeNonce",
			NewParsedAuthorizeNonceAccount(
				accounts,
				instruction,
			),
		)

	case Allocate:
		var instruction AllocateInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"allocate",
			NewParsedAllocate(
				accounts,
				instruction,
			),
		)

	case AllocateWithSeed:
		var instruction AllocateWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"allocateWithSeed",
			NewParsedAllocateWithSeed(
				accounts,
				instruction,
			),
		)

	case AssignWithSeed:
		var instruction AssignWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"assignWithSeed",
			NewParsedAssignWithSeed(
				accounts,
				instruction,
			),
		)

	case TransferWithSeed:
		var instruction TransferWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"transferWithSeed",
			NewParsedTransferWithSeed(
				accounts,
				instruction,
			),
		)
	}

	return nil
}
