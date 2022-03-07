package manager

import (
	"fmt"

	"github.com/forbole/soljuno/solana/parser"
	associatedTokenAccount "github.com/forbole/soljuno/solana/program/associated-token-account"
	"github.com/forbole/soljuno/solana/program/bpfloader"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/solana/program/token"
	tokenswap "github.com/forbole/soljuno/solana/program/token-swap"
	"github.com/forbole/soljuno/solana/program/vote"
	"github.com/forbole/soljuno/solana/types"
	"github.com/mr-tron/base58"
	"github.com/rs/zerolog/log"
)

type ParserManager interface {
	Register(programID string, parser parser.ProgramParser)
	Parse(accounts []string, programID string, base58Data string) types.ParsedInstruction
}

func NewParserManager() ParserManager {
	var p manager
	p.programs = make(map[string]parser.ProgramParser)
	return &p
}

type manager struct {
	programs map[string]parser.ProgramParser
}

func (m *manager) Register(programID string, parser parser.ProgramParser) {
	m.programs[programID] = parser
}

func (m manager) Parse(accounts []string, programID string, base58Data string) types.ParsedInstruction {
	var parsed types.ParsedInstruction
	defer func() {
		r := recover()
		if r != nil {
			log.Err(fmt.Errorf("failed to parsed message on program %v with data: %v", programID, base58Data)).Send()
			parsed = types.NewParsedInstruction("unknown", nil)
		}
	}()

	parser, ok := m.programs[programID]
	if !ok {
		return types.NewParsedInstruction("unknown", nil)
	}

	if len(base58Data) != 0 {
		bz, err := base58.Decode(base58Data)
		if err != nil {
			return types.NewParsedInstruction("unknown", nil)
		}
		return parser.Parse(accounts, bz)
	}

	parsed = parser.Parse(accounts, []byte{})
	return parsed
}

func NewDefaultManager() ParserManager {
	parser := NewParserManager()
	parser.Register(vote.ProgramID, vote.Parser{})
	parser.Register(stake.ProgramID, stake.Parser{})
	parser.Register(system.ProgramID, system.Parser{})
	parser.Register(token.ProgramID, token.Parser{})
	parser.Register(bpfloader.ProgramID, bpfloader.Parser{})
	parser.Register(upgradableLoader.ProgramID, upgradableLoader.Parser{})
	parser.Register(associatedTokenAccount.ProgramID, associatedTokenAccount.Parser{})
	parser.Register(tokenswap.ProgramID, tokenswap.Parser{})
	return parser
}
