package bpfloader

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/program/parser"
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
	case Write:
		var instruction WriteInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"write",
			NewParsedWrite(accounts, instruction),
		)

	case Finalize:
		return types.NewParsedInstruction(
			"finalize",
			NewParsedFinalize(accounts),
		)
	}
	return types.NewParsedInstruction(
		"unknown",
		nil,
	)
}
