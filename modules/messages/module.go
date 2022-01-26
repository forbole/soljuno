package messages

import (
	"sync"

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
	db     db.MsgDb
	buffer chan dbtypes.MsgRow
	pool   *ants.Pool

	mtx sync.Mutex
}

func NewModule(db db.MsgDb, pool *ants.Pool) *Module {
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
	err := m.createPartition(msg.Slot)
	if err != nil {
		return err
	}
	m.buffer <- dbtypes.NewMsgRowFromMessage(msg)
	return nil
}

// createPartition creates a new partition for the msgs module
func (m *Module) createPartition(slot uint64) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	err := m.db.CreateMsgPartition(int(slot / 1000))
	if err != nil {
		return err
	}
	return nil
}

// Prune implements pruning.PruningService
func (m *Module) Prune(slot uint64) error {
	return m.db.PruneMsgsBeforeSlot(slot)
}
