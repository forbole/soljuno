package messages

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

var _ modules.Module = &Module{}

// Module represents the module allowing to store messages properly inside a dedicated table
type Module struct {
	db     db.Database
	buffer chan types.Message
}

func NewModule(db db.Database) *Module {
	return &Module{
		db:     db,
		buffer: make(chan types.Message, 1000),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "messages"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(msg types.Message, tx types.Tx) error {
	m.buffer <- msg
	return nil
}
