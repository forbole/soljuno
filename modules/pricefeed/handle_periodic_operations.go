package pricefeed

import (
	"fmt"

	"github.com/forbole/soljuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "pricefeed").Msg("setting up periodic tasks")

	// Fetch total supply of token in 30 seconds each
	if _, err := scheduler.Every(30).Second().Do(func() {
		utils.WatchMethod(m.updatePrice)
	}); err != nil {
		return fmt.Errorf("error while setting up pricefeed period operations: %s", err)
	}

	return nil
}

// updatePrice fetch total amount of coins in the system from RPC and store it into database
func (m *Module) updatePrice() error {
	return nil
}
