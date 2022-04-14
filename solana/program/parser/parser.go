package parser

import (
	"github.com/forbole/soljuno/solana/types"
)

type ProgramParser interface {
	Parse(accounts []string, data []byte) types.ParsedInstruction
}
