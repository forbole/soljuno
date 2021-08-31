package parser

import (
	"fmt"

	"github.com/forbole/soljuno/solana/types"
)

type ProgramParser interface {
	Parse(accounts []string, data []byte) types.ParsedInstruction
}

func NewParser() *parser {
	var p parser
	p.programs = make(map[string]ProgramParser)
	return &p
}

type parser struct {
	programs map[string]ProgramParser
}

func (p *parser) Register(programID string, programParser ProgramParser) error {
	if _, ok := p.programs[programID]; ok {
		return fmt.Errorf("%s is already registered", programID)
	}

	p.programs[programID] = programParser
	return nil
}
