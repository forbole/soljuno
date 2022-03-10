package bpfloader

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different instructions types for the bpfloader module
func HandleInstruction(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	switch instruction.Parsed.Type {
	case "initializeBuffer":
		return handleMsgInitializeBuffer(instruction, tx, db, client)
	case "deployWithMaxDataLen":
		return handleMsgDeployWithMaxDataLen(instruction, tx, db, client)
	case "upgrade":
		return handleMsgUpgrade(instruction, tx, db, client)
	case "setAuthority":
		return handleMsgSetAuthority(instruction, tx, db, client)
	case "close":
		return handleMsgClose(instruction, tx, db, client)
	}
	return nil
}

// handleMsgInitializeBuffer handles a MsgInitializeBuffer
func handleMsgInitializeBuffer(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(upgradableLoader.ParsedInitializeBuffer)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeBuffer", instruction.Parsed.Type)

	}
	return updateBufferAccount(parsed.Account, tx.Slot, db, client)
}

// handleMsgDeployWithMaxDataLen handles a MsgDeployWithMaxDataLen
func handleMsgDeployWithMaxDataLen(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(upgradableLoader.ParsedDeployWithMaxDataLen)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "deployWithMaxDataLen", instruction.Parsed.Type)

	}
	if err := updateBufferAccount(parsed.BufferAccount, tx.Slot, db, client); err != nil {
		return err
	}
	if err := updateProgramAccount(parsed.ProgramAccount, tx.Slot, db, client); err != nil {
		return err
	}
	return updateProgramDataAccount(parsed.ProgramDataAccount, tx.Slot, db, client)
}

// handleMsgUpgrade handles a MsgUpgrade
func handleMsgUpgrade(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(upgradableLoader.ParsedUpgrade)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "upgrade", instruction.Parsed.Type)

	}
	if err := updateBufferAccount(parsed.BufferAccount, tx.Slot, db, client); err != nil {
		return err
	}
	if err := updateProgramAccount(parsed.ProgramAccount, tx.Slot, db, client); err != nil {
		return err
	}
	return updateProgramDataAccount(parsed.ProgramDataAccount, tx.Slot, db, client)
}

// handleMsgSetAuthority handles a MsgSetAuthority
func handleMsgSetAuthority(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(upgradableLoader.ParsedSetAuthority)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setAuthority", instruction.Parsed.Type)
	}
	if err := updateBufferAccount(parsed.Account, tx.Slot, db, client); err != nil {
		return err
	}
	return updateProgramDataAccount(parsed.Account, tx.Slot, db, client)
}

// handleMsgClose handles a MsgClose
func handleMsgClose(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(upgradableLoader.ParsedClose)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "close", instruction.Parsed.Type)

	}
	return updateProgramDataAccount(parsed.Account, tx.Slot, db, client)
}
