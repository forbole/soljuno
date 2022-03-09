package system

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the system module
func HandleMsg(msg types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	switch msg.Parsed.Type {
	case "advanceNonce":
		return handleMsgAdvanceNonce(msg, tx, db, client)
	case "withdrawFromNonce":
		return handleMsgWithdrawFromNonce(msg, tx, db, client)
	case "initializeNonce":
		return handleMsgInitializeNonce(msg, tx, db, client)
	case "authorizeNonce":
		return handleMsgAuthorizeNonce(msg, tx, db, client)
	}
	return nil
}

// handleMsgAdvanceNonce handles a MsgAdvaceNonce
func handleMsgAdvanceNonce(msg types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(system.ParsedAdvanceNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "advanceNonce", msg.Parsed.Type)

	}
	return updateNonce(instruction.NonceAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeNonce handles a MsgWithdrawFromNonce
func handleMsgWithdrawFromNonce(msg types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(system.ParsedWithdrawNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdrawNonce", msg.Parsed.Type)

	}
	return updateNonce(instruction.NonceAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeNonce handles a MsgInitializeNonce
func handleMsgInitializeNonce(msg types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(system.ParsedInitializeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeNonce", msg.Parsed.Type)

	}
	return updateNonce(instruction.NonceAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeNonce handles a MsgAuthorizeNonce
func handleMsgAuthorizeNonce(msg types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(system.ParsedAuthorizeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeNonce", msg.Parsed.Type)

	}
	return updateNonce(instruction.NonceAccount, tx.Slot, db, client)
}
