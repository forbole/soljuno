package system

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
)

type Module struct {
	db     db.Database
	client client.Proxy
}

func NewModule(db db.Database, client client.Proxy) *Module {
	return &Module{
		db:     db,
		client: client,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "system"
}
