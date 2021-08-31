package vote

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/parser"
	"github.com/forbole/soljuno/solana/types"
)

type VoteParser struct {
	parser.ProgramParser
}

func Parse(accounts []string, data []byte) types.ParsedInstruction {
	decoder := bincode.NewDecoder()
	var id InstructionID
	decoder.Decode(data[:4], &id)
	switch id {
	case InitializeAccount:
		return nil
	case Authorize:
		return nil
	case Vote:
		var instruction VoteInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"vote",
			NewParsedVote(accounts[0],
				accounts[1],
				accounts[2],
				accounts[3],
				instruction.Vote),
		)
	case Withdraw:
		return nil
	case UpdateValidatorIdentity:
		return nil
	case UpdateCommission:
		return nil
	case VoteSwitch:
		return nil
	case AuthorizeChecked:
		return nil
	}
	return nil
}

type ParsedVote struct {
	VoteAccount      string   `json:"voteAccount"`
	SlotHashesSysvar string   `json:"slotHashesSysvar"`
	ClockSysvar      string   `json:"clockSysvar"`
	VoteAuthority    string   `json:"voteAuthority"`
	Vote             VoteData `json:"vote"`
}

func NewParsedVote(voteAccount, slotHashesSysvar, clockSysvar, voteAuthority string, vote VoteData) ParsedVote {
	return ParsedVote{
		VoteAccount:      voteAccount,
		SlotHashesSysvar: slotHashesSysvar,
		ClockSysvar:      clockSysvar,
		VoteAuthority:    voteAuthority,
		Vote:             vote,
	}
}
