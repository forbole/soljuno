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
	switch msg.Value.Type() {
	case "initializeBuffer":
		return handleMsgInitializeBuffer(msg, db, client)
	case "deployWithMaxDataLen":
		return handleMsgDeployWithMaxDataLen(msg, db, client)
	case "upgrade":
		return handleMsgUpgrade(msg, db, client)
	case "setAuthority":
		return handleMsgSetAuthority(msg, db, client)
	case "close":
		return handleMsgClose(msg, db, client)
	}
	return nil
}

// handleMsgInitializeBuffer handles a MsgInitializeBuffer
func handleMsgInitializeBuffer(msg types.Message, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(upgradableLoader.ParsedInitializeBuffer)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeBuffer", msg.Value.Type())

	}
	return updateBufferAccount(instruction.Account, db, client)
}

// handleMsgDeployWithMaxDataLen handles a MsgDeployWithMaxDataLen
func handleMsgDeployWithMaxDataLen(msg types.Message, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(upgradableLoader.ParsedDeployWithMaxDataLen)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "deployWithMaxDataLen", msg.Value.Type())

	}
	return updateProgramAccount(instruction.ProgramAccount, db, client)

}

// handleMsgUpgrade handles a MsgUpgrade
func handleMsgUpgrade(msg types.Message, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(upgradableLoader.ParsedUpgrade)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "upgrade", msg.Value.Type())

	}
	return updateProgramDataAccount(instruction.ProgramDataAccount, db, client)
}

// handleMsgSetAuthority handles a MsgSetAuthority
func handleMsgSetAuthority(msg types.Message, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(upgradableLoader.ParsedSetAuthority)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setAuthority", msg.Value.Type())

	}

	return updateProgramDataAccount(instruction.Account, db, client)
}

// handleMsgClose handles a MsgClose
func handleMsgClose(msg types.Message, db db.BpfLoaderDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(upgradableLoader.ParsedClose)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "close", msg.Value.Type())

	}
	return updateProgramDataAccount(instruction.Account, db, client)
}
