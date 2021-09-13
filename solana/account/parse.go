package account_parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/program/token"
)

func Parse(programID string, bz []byte) interface{} {
	decoder := bincode.NewDecoder()
	switch programID {
	case token.ProgramID:
		return tokenParse(decoder, bz)
	}
	return nil
}
