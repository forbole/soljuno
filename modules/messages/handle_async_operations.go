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
	err := <-m.pool.DoAsync(func() error { return m.db.SaveMessages(msgs) })
	if err != nil {
		log.Error().Str("module", m.Name()).Err(err).Send()
		log.Info().Str("module", m.Name()).Msg("re-enqueueing failed messages")
		for _, msg := range msgs {
			m.buffer <- msg
		}
	}
}

func (m *Module) getMsgRows() []dbtypes.MsgRow {
	var msgs []dbtypes.MsgRow
	timeout := time.After(5 * time.Second)
	for {
		select {
		case msg := <-m.buffer:
			msgs = append(msgs, msg)
		case <-timeout:
			return msgs
		}
	}
}
