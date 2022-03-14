package config

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/config"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

var _ modules.InstructionModule = &Module{}

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
	return "config"
}

// HandleInstruction implements modules.InstructionModule
func (m *Module) HandleInstruction(instruction types.Instruction, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	if instruction.Program != config.ProgramID {
		return nil
	}

	err := HandleInstruction(instruction, tx, m.db, m.client)
	if err != nil {
		return err
	}
	log.Debug().Str("module", m.Name()).Str("tx", tx.Signature).Uint64("slot", tx.Slot).
		Msg("handled instruction")
	return nil
}
