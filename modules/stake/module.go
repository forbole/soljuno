package stake

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

type ClientProxy interface {
	GetAccountInfo(string) (clienttypes.AccountInfo, error)
}

var (
	_ modules.Module            = &Module{}
	_ modules.InstructionModule = &Module{}
)

type Module struct {
	db     db.StakeDb
	client ClientProxy
}

func NewModule(db db.StakeDb, client ClientProxy) *Module {
	return &Module{
		db:     db,
		client: client,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "stake"
}

// HandleInstruction implements modules.InstructionModule
func (m *Module) HandleInstruction(instruction types.Instruction, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	if instruction.Program != stake.ProgramID {
		return nil
	}

	err := HandleInstruction(instruction, m.db, m.client)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Str("tx", instruction.TxSignature).Uint64("slot", tx.Slot).
		Msg("handled instruction")
	return nil
}
