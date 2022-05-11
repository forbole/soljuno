package bpfloader

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"
	"github.com/forbole/soljuno/types"
)

// HandleInstruction allows to handle different instructions types for the bpfloader module
func HandleInstruction(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	switch instruction.Parsed.Type {
	case "initializeBuffer":
		return handleInitializeBuffer(instruction, tx, db, client)
	case "deployWithMaxDataLen":
		return handleDeployWithMaxDataLen(instruction, tx, db, client)
	case "upgrade":
		return handleUpgrade(instruction, tx, db, client)
	case "setAuthority":
		return handleSetAuthority(instruction, tx, db, client)
	case "close":
		return handleClose(instruction, tx, db, client)
	}
	return nil
}

// handleInitializeBuffer handles a instruction of InitializeBuffer
func handleInitializeBuffer(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(upgradableLoader.ParsedInitializeBuffer)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeBuffer", instruction.Parsed.Type)

	}
	return updateBufferAccount(parsed.Account, tx.Slot, db, client)
}

// handleDeployWithMaxDataLen handles a instruction of DeployWithMaxDataLen
func handleDeployWithMaxDataLen(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
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

// handleUpgrade handles a instruction of Upgrade
func handleUpgrade(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
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

// handleSetAuthority handles a instruction of SetAuthority
func handleSetAuthority(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(upgradableLoader.ParsedSetAuthority)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setAuthority", instruction.Parsed.Type)
	}
	if err := updateBufferAccount(parsed.Account, tx.Slot, db, client); err != nil {
		return err
	}
	return updateProgramDataAccount(parsed.Account, tx.Slot, db, client)
}

// handleClose handles a instruction ofClose
func handleClose(instruction types.Instruction, tx types.Tx, db db.BpfLoaderDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(upgradableLoader.ParsedClose)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "close", instruction.Parsed.Type)

	}
	if err := updateBufferAccount(parsed.Account, tx.Slot, db, client); err != nil {
		return err
	}
	return updateProgramDataAccount(parsed.Account, tx.Slot, db, client)
}
