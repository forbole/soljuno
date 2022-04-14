package associated_token_account

import (
	"github.com/forbole/soljuno/solana/program/parser"
	"github.com/forbole/soljuno/solana/types"
)

type Parser struct {
	parser.ProgramParser
}

func (Parser) Parse(accounts []string, data []byte) types.ParsedInstruction {
	return types.NewParsedInstruction(
		"create",
		NewParsedCreate(accounts),
	)
}
