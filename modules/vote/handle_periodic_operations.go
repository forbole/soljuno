package vote

import (
	"fmt"

	"github.com/forbole/soljuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "vote").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Minute().Do(func() {
		utils.WatchMethod(m.updateValidatorStatus)
	}); err != nil {
		return fmt.Errorf("error while setting up vote periodic operation: %s", err)
	}

	return nil
}

// updateValidatorStatus insert average block time in the latest minute
func (m *Module) updateValidatorStatus() error {
	_, _, err := m.client.ValidatorsWithSlot()
	if err != nil {
		return nil
	}
	return nil
}
