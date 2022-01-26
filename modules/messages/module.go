package messages

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/modules/pruning"
	"github.com/forbole/soljuno/types"
	"github.com/panjf2000/ants/v2"
)

var _ modules.Module = &Module{}
var _ pruning.PruningService = &Module{}

// Module represents the module allowing to store messages properly inside a dedicated table
type Module struct {
	db     db.Database
	buffer chan dbtypes.MsgRow
	pool   *ants.Pool
}

func NewModule(db db.Database, pool *ants.Pool) *Module {
	return &Module{
		db:     db,
		buffer: make(chan dbtypes.MsgRow),
		pool:   pool,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "messages"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(msg types.Message, tx types.Tx) error {
	m.buffer <- dbtypes.NewMsgRowFromMessage(msg)
	return nil
}

// Prune implements pruning.PruningService
func (m *Module) Prune(slot uint64) error {
	return nil
}
