package system

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/types"
)

// HandleInstruction allows to handle different instructions types for the system module
func HandleInstruction(instruction types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	switch instruction.Parsed.Type {
	case "advanceNonce":
		return handleAdvanceNonce(instruction, tx, db, client)
	case "withdrawFromNonce":
		return handleWithdrawFromNonce(instruction, tx, db, client)
	case "initializeNonce":
		return handleInitializeNonce(instruction, tx, db, client)
	case "authorizeNonce":
		return handleAuthorizeNonce(instruction, tx, db, client)
	}
	return nil
}

// handleAdvanceNonce handles a MsgAdvaceNonce
func handleAdvanceNonce(instruction types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(system.ParsedAdvanceNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "advanceNonce", instruction.Parsed.Type)

	}
	return updateNonce(parsed.NonceAccount, tx.Slot, db, client)
}

// handleAuthorizeNonce handles a MsgWithdrawFromNonce
func handleWithdrawFromNonce(instruction types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(system.ParsedWithdrawNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdrawNonce", instruction.Parsed.Type)

	}
	return updateNonce(parsed.NonceAccount, tx.Slot, db, client)
}

// handleAuthorizeNonce handles a MsgInitializeNonce
func handleInitializeNonce(instruction types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(system.ParsedInitializeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeNonce", instruction.Parsed.Type)

	}
	return updateNonce(parsed.NonceAccount, tx.Slot, db, client)
}

// handleAuthorizeNonce handles a MsgAuthorizeNonce
func handleAuthorizeNonce(instruction types.Instruction, tx types.Tx, db db.SystemDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(system.ParsedAuthorizeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeNonce", instruction.Parsed.Type)

	}
	return updateNonce(parsed.NonceAccount, tx.Slot, db, client)
}
