package history

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/forbole/soljuno/modules/utils"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", m.Name()).Msg("setting up periodic tasks")

	// Fetch prices of tokens in 15 minutes each
	if _, err := scheduler.Every(15).Minute().Do(func() {
		utils.WatchMethod(m.RunServices)
	}); err != nil {
		return fmt.Errorf("error while setting up history period operations: %s", err)
	}

	return nil
}
