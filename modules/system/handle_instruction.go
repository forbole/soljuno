package system

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/system"
	"github.com/forbole/soljuno/types"
)

// HandleInstruction allows to handle different instructions types for the system module
func HandleInstruction(instruction types.Instruction, db db.SystemDb, client ClientProxy) error {
	switch instruction.Parsed.Type {
	case "advanceNonce":
		return handleAdvanceNonce(instruction, db, client)
	case "withdrawFromNonce":
		return handleWithdrawFromNonce(instruction, db, client)
	case "initializeNonce":
		return handleInitializeNonce(instruction, db, client)
	case "authorizeNonce":
		return handleAuthorizeNonce(instruction, db, client)
	}
	return nil
}

// handleAdvanceNonce handles a instruction of AdvaceNonce
func handleAdvanceNonce(instruction types.Instruction, db db.SystemDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(system.ParsedAdvanceNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "advanceNonce", instruction.Parsed.Type)

	}
	return UpdateNonceAccount(parsed.NonceAccount, instruction.Slot, db, client)
}

// handleAuthorizeNonce handles a instruction of WithdrawFromNonce
func handleWithdrawFromNonce(instruction types.Instruction, db db.SystemDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(system.ParsedWithdrawNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdrawNonce", instruction.Parsed.Type)

	}
	return UpdateNonceAccount(parsed.NonceAccount, instruction.Slot, db, client)
}

// handleAuthorizeNonce handles a instruction of InitializeNonce
func handleInitializeNonce(instruction types.Instruction, db db.SystemDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(system.ParsedInitializeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeNonce", instruction.Parsed.Type)

	}
	return UpdateNonceAccount(parsed.NonceAccount, instruction.Slot, db, client)
}

// handleAuthorizeNonce handles a instruction of AuthorizeNonce
func handleAuthorizeNonce(instruction types.Instruction, db db.SystemDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(system.ParsedAuthorizeNonceAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeNonce", instruction.Parsed.Type)

	}
	return UpdateNonceAccount(parsed.NonceAccount, instruction.Slot, db, client)
}
