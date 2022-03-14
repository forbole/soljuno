package pruning

import (
	"github.com/forbole/soljuno/types/logging"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

var _ modules.Module = &Module{}

// PruningService represents a service allowing to clean the database
type PruningService interface {
	Name() string
	Prune(slot uint64) error
}

// Module represents the pruning module allowing to clean the database periodically
type Module struct {
	cfg      types.PruningConfig
	db       db.Database
	logger   logging.Logger
	signal   chan bool
	services []PruningService
}

// NewModule builds a new Module instance
func NewModule(cfg types.PruningConfig, db db.Database, logger logging.Logger) *Module {
	return &Module{
		cfg:    cfg,
		db:     db,
		logger: logger,
		signal: make(chan bool),
	}
}

// RegisterService allows to register pruning services in to the module
func (m *Module) RegisterService(services ...PruningService) {
	m.services = append(m.services, services...)
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pruning"
}

// HandleBlock implements modules.BlockModule
func (m *Module) HandleBlock(block types.Block) error {
	if m.cfg == nil {
		// Nothing to do, pruning is disabled
		return nil
	}

	if block.Height%uint64(m.cfg.GetInterval()) != 0 {
		// Not an interval height, so just skip
		return nil
	}
	m.signal <- true
	return nil
}
