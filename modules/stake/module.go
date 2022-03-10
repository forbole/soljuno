package stake

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

type Module struct {
	db     db.Database
	client client.ClientProxy
}

func NewModule(db db.Database, client client.ClientProxy) *Module {
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

	err := HandleInstruction(instruction, tx, m.db, m.client)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Str("tx", tx.Hash).Uint64("slot", tx.Slot).
		Msg("handled msg")
	return nil
}
