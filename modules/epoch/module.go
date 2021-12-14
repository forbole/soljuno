package epoch

import (
	"sync"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
)

type Module struct {
	db     db.Database
	client client.Proxy
	epoch  uint64
	mtx    sync.Mutex
}

func NewModule(db db.Database, client client.Proxy) *Module {
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
