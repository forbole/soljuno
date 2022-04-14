package parser

import (
	"github.com/forbole/soljuno/solana/bincode"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"
	"github.com/forbole/soljuno/solana/program/config"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/solana/program/token"
	"github.com/forbole/soljuno/solana/program/vote"
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
	case vote.ProgramID:
		return voteParse(decoder, bz)
	case config.ProgramID:
		return configParse(decoder, bz)
	case upgradableLoader.ProgramID:
		return bpfLoaderParse(decoder, bz)
	}
	return nil
}
