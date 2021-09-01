package parser

import (
	"github.com/forbole/soljuno/solana/types"
	"github.com/mr-tron/base58"
)

type ProgramParser interface {
	Parse(accounts []string, data []byte) types.ParsedInstruction
}

func NewParser() *parser {
	var p parser
	p.programs = make(map[string]ProgramParser)
	return &p
}

type Parser interface {
	Register(programID string, programParser ProgramParser)
	Parse(accounts []string, programID string, base58Data string) types.ParsedInstruction
}

type parser struct {
	programs map[string]ProgramParser
}

func (p *parser) Register(programID string, programParser ProgramParser) {
	p.programs[programID] = programParser
}

func (p parser) Parse(accounts []string, programID string, base58Data string) types.ParsedInstruction {
	programParser, ok := p.programs[programID]
	if !ok {
		return types.NewParsedInstruction("unknown", nil)
	}
	bz, err := base58.Decode(base58Data)
	if err != nil {
		return types.NewParsedInstruction("unknown", nil)
	}
	return programParser.Parse(accounts, bz)
}
