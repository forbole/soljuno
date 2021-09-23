package bpfloader

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"

	"github.com/forbole/soljuno/types"
)

type Module struct {
	db db.Database
}

func NewModule(db db.Database, client client.Proxy) *Module {
	return &Module{
		db: db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bpfloader"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(msg types.Message, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	if msg.Program != upgradableLoader.ProgramID {
		return nil
	}
	// _, ok := m.db.(db.BpfLoaderDb)
	// if !ok {
	// 	return fmt.Errorf("bpfloader is enabled, but your database does not implement BpfLoaderDb")
	// }
	return nil
}
