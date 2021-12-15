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
		time.Sleep(5 * time.Second)
	}
}

func (m *Module) consumeMsgs() {
	for len(m.buffer) != 0 {
		msgs := m.GetMsgs(1000)
		err := m.db.SaveMessages(msgs)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}
}

func (m *Module) GetMsgs(num int) []types.Message {
	var msgs []types.Message
	for len(m.buffer) != 0 && len(msgs) < num {
		msgs = append(msgs, <-m.buffer)
	}
	return msgs
}
