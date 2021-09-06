package upgradableLoader

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/solana/parser"
	"github.com/forbole/soljuno/solana/types"
)

type Parser struct {
	parser.ProgramParser
}

// nolint: gocyclo
func (Parser) Parse(accounts []string, data []byte) types.ParsedInstruction {
	decoder := bincode.NewDecoder()
	var id InstructionID
	decoder.Decode(data[:4], &id)
	switch id {
	case InitializeBuffer:
		return types.NewParsedInstruction(
			"initializeBuffer",
			NewParsedInitializeBuffer(
				accounts,
			),
		)

	case Write:
		var instruction WriteInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"write",
			NewParsedWrite(
				accounts,
				instruction,
			),
		)

	case DeployWithMaxDataLen:
		var instruction DeployWithMaxDataLenInstruction
		decoder.Decode(data[4:], &instruction)
		return types.NewParsedInstruction(
			"deployWithMaxDataLen",
			NewParsedDeployWithMaxDataLen(
				accounts,
				instruction,
			),
		)

	case Upgrade:
		return types.NewParsedInstruction(
			"upgrade",
			NewParsedUpgrade(
				accounts,
			),
		)

	case SetAuthority:
		return types.NewParsedInstruction(
			"setAuthority",
			NewParsedSetAuthority(
				accounts,
			),
		)

	case Close:
		return types.NewParsedInstruction(
			"close",
			NewParsedClose(
				accounts,
			),
		)
	}
	return types.NewParsedInstruction(
		"unknown",
		nil,
	)
}
