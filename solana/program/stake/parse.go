package stake

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
	case Initialize:
	case Authorize:
	case DelegateStake:
	case Split:
	case Withdraw:
	case Deactivate:
	case SetLockup:
	case Merge:
	case AuthorizeWithSeed:
	case InitializeChecked:
	case AuthorizeChecked:
	case AuthorizeCheckedWithSeed:
	case SetLockupChecked:
	}
	return nil
}
