package bpfloader

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/solana/client"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"
	"github.com/rs/zerolog/log"

	"github.com/forbole/soljuno/types"
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
	return "bpfloader"
}

// HandleInstruction implements modules.InstructionModule
func (m *Module) HandleInstruction(instruction types.Instruction, tx types.Tx) error {
	if !tx.Successful() {
		return nil
	}
	if instruction.Program != upgradableLoader.ProgramID {
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
