package stake

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
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
	return "stake"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(msg types.Message, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	if msg.Program != stake.ProgramID {
		return nil
	}
	stakeDb, ok := m.db.(db.StakeDb)
	if !ok {
		return fmt.Errorf("stake is enabled, but your database does not implement StakeDb")
	}

	err := HandleMsg(msg, tx, stakeDb, m.client)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Str("tx", tx.Hash).Uint64("slot", tx.Slot).
		Msg("handled msg")
	return nil
}
