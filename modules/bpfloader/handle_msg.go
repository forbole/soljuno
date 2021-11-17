package bpfloader

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	upgradableLoader "github.com/forbole/soljuno/solana/program/bpfloader/upgradeable"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the bpfloader module
func HandleMsg(msg types.Message, tx types.Tx, db db.BpfLoaderDb, client client.Proxy) error {
	switch msg.Parsed.Type() {
	case "initializeBuffer":
		return handleMsgInitializeBuffer(msg, tx, db, client)
	case "deployWithMaxDataLen":
		return handleMsgDeployWithMaxDataLen(msg, tx, db, client)
	case "upgrade":
		return handleMsgUpgrade(msg, tx, db, client)
	case "setAuthority":
		return handleMsgSetAuthority(msg, tx, db, client)
	case "close":
		return handleMsgClose(msg, tx, db, client)
	}
	return nil
}

// handleMsgInitializeBuffer handles a MsgInitializeBuffer
func handleMsgInitializeBuffer(msg types.Message, tx types.Tx, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(upgradableLoader.ParsedInitializeBuffer)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeBuffer", msg.Parsed.Type())

	}
	return updateBufferAccount(instruction.Account, tx.Slot, db, client)
}

// handleMsgDeployWithMaxDataLen handles a MsgDeployWithMaxDataLen
func handleMsgDeployWithMaxDataLen(msg types.Message, tx types.Tx, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(upgradableLoader.ParsedDeployWithMaxDataLen)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "deployWithMaxDataLen", msg.Parsed.Type())

	}
	if err := updateBufferAccount(instruction.BufferAccount, tx.Slot, db, client); err != nil {
		return err
	}
	if err := updateProgramAccount(instruction.ProgramAccount, tx.Slot, db, client); err != nil {
		return err
	}
	return updateProgramDataAccount(instruction.ProgramDataAccount, tx.Slot, db, client)
}

// handleMsgUpgrade handles a MsgUpgrade
func handleMsgUpgrade(msg types.Message, tx types.Tx, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(upgradableLoader.ParsedUpgrade)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "upgrade", msg.Parsed.Type())

	}
	if err := updateBufferAccount(instruction.BufferAccount, tx.Slot, db, client); err != nil {
		return err
	}
	if err := updateProgramAccount(instruction.ProgramAccount, tx.Slot, db, client); err != nil {
		return err
	}
	return updateProgramDataAccount(instruction.ProgramDataAccount, tx.Slot, db, client)
}

// handleMsgSetAuthority handles a MsgSetAuthority
func handleMsgSetAuthority(msg types.Message, tx types.Tx, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(upgradableLoader.ParsedSetAuthority)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setAuthority", msg.Parsed.Type())
	}
	if err := updateBufferAccount(instruction.Account, tx.Slot, db, client); err != nil {
		return err
	}
	return updateProgramDataAccount(instruction.Account, tx.Slot, db, client)
}

// handleMsgClose handles a MsgClose
func handleMsgClose(msg types.Message, tx types.Tx, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(upgradableLoader.ParsedClose)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "close", msg.Parsed.Type())

	}
	return updateProgramDataAccount(instruction.Account, tx.Slot, db, client)
}
