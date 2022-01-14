package epoch

import (
	"sync"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
)

type Module struct {
	db       db.Database
	client   client.ClientProxy
	epoch    uint64
	mtx      sync.Mutex
	services []EpochService
}

func NewModule(db db.Database, client client.ClientProxy) *Module {
	return &Module{
		db:     db,
		client: client,
		epoch:  0,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "epoch"
}
