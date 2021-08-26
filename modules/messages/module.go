package messages

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/solana/bincode"
)

var _ modules.Module = &Module{}

// Module represents the module allowing to store messages properly inside a dedicated table
type Module struct {
	parser MessageParser

	cdc bincode.Decoder
	db  db.Database
}

func NewModule(parser MessageParser, cdc bincode.Decoder, db db.Database) *Module {
	return &Module{
		parser: parser,
		cdc:    cdc,
		db:     db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "messages"
}

// HandleMsg implements modules.MessageModule
// func (m *Module) HandleInstruction(index int, instruction types.Instruction, tx *types.Tx) error {
// 	return HandleMsg(index, instruction, tx, m.parser, m.cdc, m.db)
// }
