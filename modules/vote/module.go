package vote

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/vote"
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
	return "vote"
}

// HandleMsg implements modules.MessageModule
func (m *Module) HandleMsg(msg types.Message, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	if msg.Program != vote.ProgramID {
		return nil
	}
	voteDb, ok := m.db.(db.VoteDb)
	if !ok {
		return fmt.Errorf("vote is enabled, but your database does not implement VoteDb")
	}
	return HandleMsg(msg, tx, voteDb, m.client)
}
