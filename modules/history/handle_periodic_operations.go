package history

import (
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", m.Name()).Msg("setting up periodic tasks")
	return m.RunServices(scheduler)
}

func (m *Module) RunPeriodicOperations() error {
	for _, service := range m.services {
		err := service.ExecHistory()
		if err != nil {
			return err
		}
	}
	return nil
}
