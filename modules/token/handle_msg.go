package token

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/token"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the token module
func HandleMsg(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	switch msg.Parsed.Type {
	case "initializeMint":
		return handleMsgInitializeMint(msg, tx, db)
	case "initializeMint2":
		return handleMsgInitializeMint(msg, tx, db)

	case "initializeAccount":
		return handleMsgInitializeAccount(msg, tx, db)
	case "initializeAccount2":
		return handleMsgInitializeAccount(msg, tx, db)
	case "initializeAccount3":
		return handleMsgInitializeAccount(msg, tx, db)

	case "initializeMultisig":
		return handleMsgInitializeMultisig(msg, tx, db)
	case "initializeMultisig2":
		return handleMsgInitializeMultisig(msg, tx, db)

	// TODO: make decision if handle token balance instead of bank module
	case "transfer":
		return nil
	case "transferChecked":
		return nil

	// Delegate msgs
	case "approve":
		return handleMsgApprove(msg, tx, db, client)
	case "approveChecked":
		return handleMsgApproveChecked(msg, tx, db, client)
	case "revoke":
		return handleMsgRevoke(msg, tx, db, client)

	case "setAuthority":
		return handleSetAuthority(msg, tx, db, client)

	// Token supply msgs
	case "mintTo":
		return handleMsgMintTo(msg, tx, db, client)
	case "mintToChecked":
		return handleMsgMintToChecked(msg, tx, db, client)
	case "burn":
		return handleMsgBurn(msg, tx, db, client)
	case "burnChecked":
		return handleMsgBurnChecked(msg, tx, db, client)

	// Account state msgs
	case "closeAccount":
		return handleMsgCloseAccount(msg, tx, db, client)
	case "freezeAccount":
		return handleMsgFreezeAccount(msg, tx, db, client)
	case "thawAccount":
		return handleMsgThawAccount(msg, tx, db, client)
	}
	return nil
}

// handleMsgInitializeMint handles a MsgInitializeMint and properly stores the new token inside the database
func handleMsgInitializeMint(msg types.Instruction, tx types.Tx, db db.TokenDb) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedInitializeMint)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeMint", msg.Parsed.Type)
	}
	err := db.SaveToken(
		instruction.Mint,
		tx.Slot,
		instruction.Decimals,
		instruction.MintAuthority,
		instruction.FreezeAuthority,
	)
	if err != nil {
		return err
	}
	return db.SaveTokenSupply(instruction.Mint, tx.Slot, 0)
}

// handleMsgInitializeAccount handles a MsgInitializeAccount and properly stores the new token account inside the database
func handleMsgInitializeAccount(msg types.Instruction, tx types.Tx, db db.TokenDb) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedInitializeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeAccount", msg.Parsed.Type)
	}
	err := db.SaveTokenAccount(
		instruction.Account,
		tx.Slot,
		instruction.Mint,
		instruction.Owner,
	)
	if err != nil {
		return err
	}
	return nil
}

// handleMsgInitializeMultisig handles a MsgInitializeMultisig and properly stores the new multisig inside the database
func handleMsgInitializeMultisig(msg types.Instruction, tx types.Tx, db db.TokenDb) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedInitializeMultisig)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeMultisig", msg.Parsed.Type)
	}
	err := db.SaveMultisig(
		instruction.MultiSig,
		tx.Slot,
		instruction.Signers,
		instruction.M,
	)
	if err != nil {
		return err
	}
	return nil
}

// handleMsgApproveChecked handles a MsgApprove
func handleMsgApprove(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedApprove)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approve", msg.Parsed.Type)
	}
	return updateDelegation(instruction.Source, tx.Slot, db, client)
}

// handleMsgApproveChecked handles a MsgApproveChecked
func handleMsgApproveChecked(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedApproveChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approveChecked", msg.Parsed.Type)
	}
	return updateDelegation(instruction.Source, tx.Slot, db, client)
}

// handleMsgRevoke handles a MsgRevoke
func handleMsgRevoke(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedRevoke)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approveChecked", msg.Parsed.Type)
	}
	return updateDelegation(instruction.Source, tx.Slot, db, client)
}

// handleSetAuthority handles a MsgSetAuthority
func handleSetAuthority(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedSetAuthority)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setAuthority", msg.Parsed.Type)
	}
	if instruction.Mint != "" {
		return updateToken(instruction.Mint, tx.Slot, db, client)
	}
	return updateTokenAccount(instruction.Account, tx.Slot, db, client)
}

// handleMsgMintTo handles a MsgMintTo
func handleMsgMintTo(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedMintTo)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "mintTo", msg.Parsed.Type)
	}
	return updateTokenSupply(instruction.Mint, tx.Slot, db, client)
}

// handleMsgMintToChecked handles a MsgMintToChecked
func handleMsgMintToChecked(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedMintToChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "mintToChecked", msg.Parsed.Type)
	}
	return updateTokenSupply(instruction.Mint, tx.Slot, db, client)
}

// handleBurn handles a MsgBurn
func handleMsgBurn(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedBurn)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "burn", msg.Parsed.Type)
	}
	return updateTokenSupply(instruction.Mint, tx.Slot, db, client)
}

// handleBurn handles a MsgBurnChecked
func handleMsgBurnChecked(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedBurnChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "burnChecked", msg.Parsed.Type)
	}
	return updateTokenSupply(instruction.Mint, tx.Slot, db, client)
}

// handleMsgCloseAccount handles a MsgCloseAccount
func handleMsgCloseAccount(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedCloseAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "closeAccount", msg.Parsed.Type)
	}
	return updateTokenAccount(instruction.Account, tx.Slot, db, client)
}

// handleMsgFreezeAccount handles a MsgFreezeAccount
func handleMsgFreezeAccount(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedFreezeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "freezeAccount", msg.Parsed.Type)
	}
	return updateTokenAccount(instruction.Account, tx.Slot, db, client)
}

// handleMsgThawAccount handles a MsgThawAccount
func handleMsgThawAccount(msg types.Instruction, tx types.Tx, db db.TokenDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(token.ParsedThawAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "thawAccount", msg.Parsed.Type)
	}
	return updateTokenAccount(instruction.Account, tx.Slot, db, client)
}
