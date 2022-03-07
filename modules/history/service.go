package history

import (
	"fmt"

	"github.com/forbole/soljuno/modules/utils"
	"github.com/go-co-op/gocron"
)

func (m *Module) RegisterService(services ...HistroyService) {
	m.services = append(m.services, services...)
}

func (m *Module) RunServices(scheduler *gocron.Scheduler) error {
	for _, service := range m.services {
		m.registerService(scheduler, service)
	}
	return nil
}

func (m *Module) registerService(scheduler *gocron.Scheduler, service HistroyService) error {
	if _, err := scheduler.Cron(service.Cron()).Do(func() {
		utils.WatchMethod(m, service.ExecHistory)
	}); err != nil {
		return fmt.Errorf("error while setting up history period operations: %s", err)
	}
	return nil
}
