package messages

import (
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		m.consumeMsgs()
	}
}

func (m *Module) consumeMsgs() {
	msgs := m.getMsgRows()
	_ = m.pool.Submit(func() {
		err := m.db.SaveMessages(msgs)
		if err != nil {
			log.Error().Str("module", m.Name()).Uint64("slot", msgs[0].Slot).Err(err).Send()
		}
	})
}

func (m *Module) getMsgRows() []dbtypes.MsgRow {
	var msgs []dbtypes.MsgRow
	for {
		select {
		case msg := <-m.buffer:
			msgs = append(msgs, msg)
		case <-time.After(1 * time.Second):
			return msgs
		}
	}
}
