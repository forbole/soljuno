package instructions

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
		err := m.db.SaveInstructions(msgs)
		m.handleAsyncError(err, msgs)
		return nil
	})

	m.handleAsyncError(err, msgs)
}

func (m *Module) getMsgRows() []dbtypes.InstructionRow {
	var instructions []dbtypes.InstructionRow
	timeout := time.After(5 * time.Second)
	for {
		select {
		case instruction := <-m.buffer:
			instructions = append(instructions, instruction)
		case <-timeout:
			return instructions
		}
	}
}

func (m *Module) handleAsyncError(err error, instructions []dbtypes.InstructionRow) {
	if err != nil {
		// re-enqueueing failed instructions in the same goroutine
		log.Error().Str("module", m.Name()).Err(err).Send()
		log.Info().Str("module", m.Name()).Msg("re-enqueueing failed instructions")
		for _, instruction := range instructions {
			m.buffer <- instruction
		}
	}
}
