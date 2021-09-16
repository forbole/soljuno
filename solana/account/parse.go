package account_parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/solana/program/token"
)

func Parse(programID string, bz []byte) interface{} {
	decoder := bincode.NewDecoder()
	switch programID {
	case token.ProgramID:
		return tokenParse(decoder, bz)
	case system.ProgramID:
		return systemParse(decoder, bz)
	case stake.ProgramID:
		return stakeParse(decoder, bz)
	}
	return nil
}
