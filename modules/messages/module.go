package messages

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/types"
)

var _ modules.Module = &Module{}

// Module represents the module allowing to store messages properly inside a dedicated table
type Module struct {
	cdc bincode.Decoder
	db  db.Database
}

func NewModule(cdc bincode.Decoder, db db.Database) *Module {
	return &Module{
		cdc: cdc,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "messages"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(index int, msg types.Message, tx types.Tx) error {
	return m.db.SaveMessage(msg)
}
