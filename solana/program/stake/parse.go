package stake

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/parser"
	"github.com/forbole/soljuno/solana/types"
)

type Parser struct {
	parser.ProgramParser
}

// nolint: gocyclo
func (Parser) Parse(accounts []string, data []byte) types.ParsedInstruction {
	decoder := bincode.NewDecoder()
	var id InstructionID
	decoder.Decode(data[:4], &id)
	switch id {
	case Initialize:
		var instruction InitializeInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initialize",
			NewParsedInitialize(
				accounts,
				instruction,
			),
		)

	case Authorize:
		var instruction AuthorizeInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"authorize",
			NewParsedAuthorize(
				accounts,
				instruction,
			),
		)

	case DelegateStake:
		var instruction InitializeInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"delegate",
			NewParsedDelegateStake(
				accounts,
			),
		)

	case Split:
		var instruction SplitInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"split",
			NewParsedSplit(
				accounts,
				instruction,
			),
		)

	case Withdraw:
		var instruction WithdrawInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"withdraw",
			NewParsedWithdraw(
				accounts,
				instruction,
			),
		)

	case Deactivate:
		return types.NewParsedInstruction(
			"deactivate",
			NewParsedDeactivate(
				accounts,
			),
		)

	case SetLockup:
		var instruction SetLockupInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"setLockup",
			NewParsedSetLockup(
				accounts,
				instruction,
			),
		)

	case Merge:
		return types.NewParsedInstruction(
			"merge",
			NewParsedMerge(
				accounts,
			),
		)

	case AuthorizeWithSeed:
		var instruction AuthorizeWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"authorizeWithSeed",
			NewParsedAuthorizedWithSeed(
				accounts,
				instruction,
			),
		)

	case InitializeChecked:
		return types.NewParsedInstruction(
			"initializeChecked",
			NewInitializeChecked(
				accounts,
			),
		)

	case AuthorizeChecked:
		var instruction AuthorizeCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"authorizeChecked",
			NewParsedAuthorizeChecked(
				accounts,
				instruction,
			),
		)

	case AuthorizeCheckedWithSeed:
		var instruction AuthorizeCheckedWithSeedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"authorizeCheckedWithSeed",
			NewParsedAuthorizeCheckedWithSeed(
				accounts,
				instruction,
			),
		)

	case SetLockupChecked:
		var instruction SetLockupCheckedInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"setLockupChecked",
			NewParsedSetLockupChecked(
				accounts,
				instruction,
			),
		)

	}
	return types.NewParsedInstruction(
		"unknown",
		nil,
	)
}
