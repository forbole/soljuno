package pricefeed

import (
	"github.com/forbole/soljuno/apis/coingecko"
	"github.com/forbole/soljuno/db"
	"github.com/go-co-op/gocron"
)

type Module struct {
	db     db.Database
	client coingecko.Client
}

func NewModule(db db.Database) *Module {
	return &Module{
		db:     db,
		client: coingecko.NewDefaultClient(),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pricefeed"
}

func (m *Module) PeriodicOperationsModule(scheduler *gocron.Scheduler) error {
	return m.RegisterPeriodicOperations(scheduler)
}
