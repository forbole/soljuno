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
	switch msg.Value.Type() {
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
	instruction, ok := msg.Value.Data().(system.ParsedAdvanceNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "advanceNonce", msg.Value.Type())

	}
	return updateNonce(instruction.NonceAccount, db, client)
}

// handleMsgAuthorizeNonce handles a MsgWithdrawFromNonce
func handleMsgWithdrawFromNonce(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(system.ParsedWithdrawNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdrawNonce", msg.Value.Type())

	}
	return updateNonce(instruction.NonceAccount, db, client)
}

// handleMsgAuthorizeNonce handles a MsgInitializeNonce
func handleMsgInitializeNonce(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(system.ParsedInitializeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeNonce", msg.Value.Type())

	}
	return updateNonce(instruction.NonceAccount, db, client)
}

// handleMsgAuthorizeNonce handles a MsgAuthorizeNonce
func handleMsgAuthorizeNonce(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(system.ParsedAuthorizeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeNonce", msg.Value.Type())

	}
	return updateNonce(instruction.NonceAccount, db, client)
}
