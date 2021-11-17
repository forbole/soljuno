package system

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the system module
func HandleMsg(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	switch msg.Parsed.Type() {
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
func handleMsgAdvanceNonce(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(system.ParsedAdvanceNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "advanceNonce", msg.Parsed.Type())

	}
	return updateNonce(instruction.NonceAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeNonce handles a MsgWithdrawFromNonce
func handleMsgWithdrawFromNonce(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(system.ParsedWithdrawNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdrawNonce", msg.Parsed.Type())

	}
	return updateNonce(instruction.NonceAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeNonce handles a MsgInitializeNonce
func handleMsgInitializeNonce(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(system.ParsedInitializeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeNonce", msg.Parsed.Type())

	}
	return updateNonce(instruction.NonceAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeNonce handles a MsgAuthorizeNonce
func handleMsgAuthorizeNonce(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	instruction, ok := msg.Parsed.Data().(system.ParsedAuthorizeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeNonce", msg.Parsed.Type())

	}
	return updateNonce(instruction.NonceAccount, tx.Slot, db, client)
}
