package consensus

import (
	"fmt"

	"github.com/forbole/soljuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	if _, err := scheduler.Every(1).Hour().Do(func() {
		utils.WatchMethod(m, m.updateSlotTimeInHour)
	}); err != nil {
		return fmt.Errorf("error while setting up consensus periodic operation: %s", err)
	}
	return nil
}

// updateSlotTimeInHour insert average slot time in the latest hour
func (m *Module) updateSlotTimeInHour() error {
	log.Trace().Str("module", "consensus").Str("operation", "block time").
		Msg("updating block time in hours")

	block, err := m.db.GetLastBlock()
	if err != nil {
		return fmt.Errorf("error while getting last block: %s", err)
	}

	hour, err := m.db.GetBlockHourAgo(block.Timestamp)
	if err != nil {
		return fmt.Errorf("error while getting block an hour ago: %s", err)
	}
	newBlockTime := block.Timestamp.Sub(hour.Timestamp).Seconds() / float64(block.Slot-hour.Slot)

	return m.db.SaveAverageSlotTimePerHour(block.Slot, newBlockTime)
}
