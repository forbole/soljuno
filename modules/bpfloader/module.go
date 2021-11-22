package bpfloader

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"
	"github.com/rs/zerolog/log"

	"github.com/forbole/soljuno/types"
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

	err := HandleMsg(msg, tx, m.db, m.client)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Str("tx", tx.Hash).Uint64("slot", tx.Slot).
		Msg("handled msg")
	return nil
}
