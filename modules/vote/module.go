package vote

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/solana/program/vote"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

type ClientProxy interface {
	GetAccountInfo(string) (clienttypes.AccountInfo, error)
	GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error)
	GetLeaderSchedule(uint64) (clienttypes.LeaderSchedule, error)
}

var (
	_ modules.Module            = &Module{}
	_ modules.InstructionModule = &Module{}
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
	return "vote"
}

// HandleInstruction implements modules.InstructionModule
func (m *Module) HandleInstruction(instruction types.Instruction, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	if instruction.Program != vote.ProgramID {
		return nil
	}

	err := HandleInstruction(instruction, m.db, m.client)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Str("tx", tx.Signature).Uint64("slot", tx.Slot).
		Msg("handled instruction")
	return nil
}
