package pricefeed

import (
	"github.com/forbole/soljuno/db"
	"github.com/go-co-op/gocron"
)

type Module struct {
	db db.Database
}

func NewModule(db db.Database) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pricefeed"
}

func (m *Module) PeriodicOperationsModule(scheduler *gocron.Scheduler) error {
	return nil
}
