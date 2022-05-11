package votestatus

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

type ClientProxy interface {
	GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error)
	GetLeaderSchedule(uint64) (clienttypes.LeaderSchedule, error)
}

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	db     db.VoteDb
	client ClientProxy
}

func NewModule(db db.VoteDb, client ClientProxy) *Module {
	return &Module{
		db:     db,
		client: client,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "vote-status"
}
