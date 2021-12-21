package messages

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
	"github.com/panjf2000/ants/v2"
)

var _ modules.Module = &Module{}

// Module represents the module allowing to store messages properly inside a dedicated table
type Module struct {
	db     db.Database
	buffer chan types.Message
	pool   *ants.Pool
}

func NewModule(db db.Database, pool *ants.Pool) *Module {
	return &Module{
		db:     db,
		buffer: make(chan types.Message),
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
