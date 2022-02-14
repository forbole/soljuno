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
	err := <-m.pool.DoAsync(func() error {
		err := m.db.SaveMessages(msgs)

		// re-enqueueing failed messages in the same goroutine
		if err != nil {
			log.Error().Str("module", m.Name()).Err(err).Send()
			for _, msg := range msgs {
				m.buffer <- msg
			}
			log.Info().Str("module", m.Name()).Msg("re-enqueueing failed messages")
		}
		return nil
	})

	if err != nil {
		log.Error().Str("module", m.Name()).Err(err).Send()
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
