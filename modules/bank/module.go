package bank

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

type Module struct {
	db    db.Database
	tasks chan func()
}

func NewModule(db db.Database) *Module {
	return &Module{
		db:    db,
		tasks: make(chan func()),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bank"
}

// HandleBank implements modules.BankModule
func (m *Module) HandleBlock(block types.Block) error {
	m.tasks <- func() {
		err := HandleBank(block, m.db)
		if err != nil {
			log.Error().Str("module", m.Name()).Uint64("slot", block.Slot).Err(err).Send()
		}
	}
	log.Debug().Str("module", m.Name()).Uint64("slot", block.Slot).Msg("handled block")
	return nil
}
