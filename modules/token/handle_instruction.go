package token

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/token"
	"github.com/forbole/soljuno/types"
)

// HandleInstruction allows to handle different messages types for the token module
func HandleInstruction(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	switch instruction.Parsed.Type {
	case "initializeMint":
		return handleInitializeMint(instruction, tx, db)
	case "initializeMint2":
		return handleInitializeMint(instruction, tx, db)

	case "initializeAccount":
		return handleInitializeAccount(instruction, tx, db)
	case "initializeAccount2":
		return handleInitializeAccount(instruction, tx, db)
	case "initializeAccount3":
		return handleInitializeAccount(instruction, tx, db)

	case "initializeMultisig":
		return handleInitializeMultisig(instruction, tx, db)
	case "initializeMultisig2":
		return handleInitializeMultisig(instruction, tx, db)

	// TODO: make decision if handle token balance instead of bank module
	case "transfer":
		return nil
	case "transferChecked":
		return nil

	// Delegate msgs
	case "approve":
		return handleApprove(instruction, tx, db, client)
	case "approveChecked":
		return handleApproveChecked(instruction, tx, db, client)
	case "revoke":
		return handleRevoke(instruction, tx, db, client)

	case "setAuthority":
		return handleSetAuthority(instruction, tx, db, client)

	// Token supply msgs
	case "mintTo":
		return handleMintTo(instruction, tx, db, client)
	case "mintToChecked":
		return handleMintToChecked(instruction, tx, db, client)
	case "burn":
		return handleBurn(instruction, tx, db, client)
	case "burnChecked":
		return handleBurnChecked(instruction, tx, db, client)

	// Account state msgs
	case "closeAccount":
		return handleCloseAccount(instruction, tx, db, client)
	case "freezeAccount":
		return handleFreezeAccount(instruction, tx, db, client)
	case "thawAccount":
		return handleThawAccount(instruction, tx, db, client)
	}
	return nil
}

// handleInitializeMint handles a MsgInitializeMint and properly stores the new token inside the database
func handleInitializeMint(instruction types.Instruction, tx types.Tx, db db.TokenDb) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedInitializeMint)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeMint", instruction.Parsed.Type)
	}
	err := db.SaveToken(
		parsed.Mint,
		tx.Slot,
		parsed.Decimals,
		parsed.MintAuthority,
		parsed.FreezeAuthority,
	)
	if err != nil {
		return err
	}
	return db.SaveTokenSupply(parsed.Mint, tx.Slot, 0)
}

// handleInitializeAccount handles a MsgInitializeAccount and properly stores the new token account inside the database
func handleInitializeAccount(instruction types.Instruction, tx types.Tx, db db.TokenDb) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedInitializeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeAccount", instruction.Parsed.Type)
	}
	err := db.SaveTokenAccount(
		parsed.Account,
		tx.Slot,
		parsed.Mint,
		parsed.Owner,
	)
	if err != nil {
		return err
	}
	return nil
}

// handleInitializeMultisig handles a MsgInitializeMultisig and properly stores the new multisig inside the database
func handleInitializeMultisig(instruction types.Instruction, tx types.Tx, db db.TokenDb) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedInitializeMultisig)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeMultisig", instruction.Parsed.Type)
	}
	err := db.SaveMultisig(
		parsed.MultiSig,
		tx.Slot,
		parsed.Signers,
		parsed.M,
	)
	if err != nil {
		return err
	}
	return nil
}

// handleApproveChecked handles a MsgApprove
func handleApprove(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedApprove)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approve", instruction.Parsed.Type)
	}
	return updateDelegation(parsed.Source, tx.Slot, db, client)
}

// handleApproveChecked handles a MsgApproveChecked
func handleApproveChecked(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedApproveChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approveChecked", instruction.Parsed.Type)
	}
	return updateDelegation(parsed.Source, tx.Slot, db, client)
}

// handleRevoke handles a MsgRevoke
func handleRevoke(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedRevoke)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approveChecked", instruction.Parsed.Type)
	}
	return updateDelegation(parsed.Source, tx.Slot, db, client)
}

// handleSetAuthority handles a MsgSetAuthority
func handleSetAuthority(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedSetAuthority)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setAuthority", instruction.Parsed.Type)
	}
	if parsed.Mint != "" {
		return updateToken(parsed.Mint, tx.Slot, db, client)
	}
	return updateTokenAccount(parsed.Account, tx.Slot, db, client)
}

// handleMintTo handles a MsgMintTo
func handleMintTo(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedMintTo)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "mintTo", instruction.Parsed.Type)
	}
	return updateTokenSupply(parsed.Mint, tx.Slot, db, client)
}

// handleMintToChecked handles a MsgMintToChecked
func handleMintToChecked(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedMintToChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "mintToChecked", instruction.Parsed.Type)
	}
	return updateTokenSupply(parsed.Mint, tx.Slot, db, client)
}

// handleBurn handles a MsgBurn
func handleBurn(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedBurn)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "burn", instruction.Parsed.Type)
	}
	return updateTokenSupply(parsed.Mint, tx.Slot, db, client)
}

// handleBurn handles a MsgBurnChecked
func handleBurnChecked(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedBurnChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "burnChecked", instruction.Parsed.Type)
	}
	return updateTokenSupply(parsed.Mint, tx.Slot, db, client)
}

// handleCloseAccount handles a MsgCloseAccount
func handleCloseAccount(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedCloseAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "closeAccount", instruction.Parsed.Type)
	}
	return updateTokenAccount(parsed.Account, tx.Slot, db, client)
}

// handleFreezeAccount handles a MsgFreezeAccount
func handleFreezeAccount(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedFreezeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "freezeAccount", instruction.Parsed.Type)
	}
	return updateTokenAccount(parsed.Account, tx.Slot, db, client)
}

// handleThawAccount handles a MsgThawAccount
func handleThawAccount(instruction types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(token.ParsedThawAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "thawAccount", instruction.Parsed.Type)
	}
	return updateTokenAccount(parsed.Account, tx.Slot, db, client)
}
