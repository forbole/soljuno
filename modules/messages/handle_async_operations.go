package messages

import (
	"time"

	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

// RunAsyncOperations implements modules.Module
func (m *Module) RunAsyncOperations() {
	for {
		m.consumeMsgs()
	}
}

func (m *Module) consumeMsgs() {
	msgs := m.GetMsgs(1000)
	err := m.db.SaveMessages(msgs)
	if err != nil {
		log.Error().Str("module", m.Name()).Uint64("slot", msgs[0].Slot).Err(err).Send()
		log.Error().Err(err).Send()
	}
}

func (m *Module) GetMsgs(num int) []types.Message {
	var msgs []types.Message
	for {
		select {
		case msg := <-m.buffer:
			msgs = append(msgs, msg)
			if len(msgs) >= num {
				return msgs
			}
		case <-time.After(100 * time.Millisecond):
			return msgs
		}
	}
}
