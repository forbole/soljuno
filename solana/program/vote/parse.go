package vote

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/program/parser"
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
	case InitializeAccount:
		var instruction InitializeAccountInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"initialize",
			NewParsedInitializeAccount(
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

	case Vote:
		var instruction VoteInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"vote",
			NewParsedVote(
				accounts,
				instruction),
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

	case UpdateValidatorIdentity:
		return types.NewParsedInstruction(
			"updateValidatorIdentity",
			NewParsedUpdateValidatorIdentity(
				accounts,
			),
		)

	case UpdateCommission:
		var instruction UpdateCommissionInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"updateCommission",
			NewParsedUpdateCommission(
				accounts,
				instruction,
			),
		)

	case VoteSwitch:
		var instruction VoteSwitchInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"voteSwitch",
			NewParsedVoteSwitch(
				accounts,
				instruction,
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
	}

	return types.NewParsedInstruction(
		"unknown",
		nil,
	)
}
