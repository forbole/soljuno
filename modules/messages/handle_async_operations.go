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
	_, err := m.pool.DoAsync(func() error {
		err := m.db.SaveMessages(msgs)
		m.handleAsyncError(err, msgs)
		return nil
	})

	m.handleAsyncError(err, msgs)
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

func (m *Module) handleAsyncError(err error, msgs []dbtypes.MsgRow) {
	if err != nil {
		// re-enqueueing failed messages in the same goroutine
		log.Error().Str("module", m.Name()).Err(err).Send()
		log.Info().Str("module", m.Name()).Msg("re-enqueueing failed messages")
		for _, msg := range msgs {
			m.buffer <- msg
		}
	}
}
