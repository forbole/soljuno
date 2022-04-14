package tokenswap

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
	decoder.Decode(data[:1], &id)
	switch id {
	case Initialize:
		return types.NewParsedInstruction(
			"initialize",
			nil,
		)

	case Swap:
		return types.NewParsedInstruction(
			"swap",
			nil,
		)

	case DepositAllTokenTypes:
		return types.NewParsedInstruction(
			"depositAllTokenTypes",
			nil,
		)

	case WithdrawAllTokenTypes:
		return types.NewParsedInstruction(
			"withdrawAllTokenTypes",
			nil,
		)

	case DepositSingleTokenTypeExactAmountIn:
		return types.NewParsedInstruction(
			"depositSingleTokenTypeExactAmountIn",
			nil,
		)

	case WithdrawSingleTokenTypeExactAmountOut:
		return types.NewParsedInstruction(
			"withdrawSingleTokenTypeExactAmountOut",
			nil,
		)
	}
	return types.NewParsedInstruction(
		"unknown",
		nil,
	)
}
