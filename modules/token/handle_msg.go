package token

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/token"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the token module
func HandleMsg(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	switch msg.Value.Type() {
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
func handleMsgInitializeMint(msg types.Message, tx types.Tx, db db.TokenDb) error {
	instruction, ok := msg.Value.Data().(token.ParsedInitializeMint)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeMint", msg.Value.Type())
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
func handleMsgInitializeAccount(msg types.Message, tx types.Tx, db db.TokenDb) error {
	instruction, ok := msg.Value.Data().(token.ParsedInitializeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeAccount", msg.Value.Type())
	}
	err := db.SaveTokenAccount(
		instruction.Account,
		tx.Slot,
		instruction.Mint,
		instruction.Owner,
		"initialized",
	)
	if err != nil {
		return err
	}
	return nil
}

// handleMsgInitializeMultisig handles a MsgInitializeMultisig and properly stores the new multisig inside the database
func handleMsgInitializeMultisig(msg types.Message, tx types.Tx, db db.TokenDb) error {
	instruction, ok := msg.Value.Data().(token.ParsedInitializeMultisig)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeMultisig", msg.Value.Type())
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
func handleMsgApprove(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedApprove)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approve", msg.Value.Type())
	}
	return updateDelegation(instruction.Source, tx.Slot, db, client)
}

// handleMsgApproveChecked handles a MsgApproveChecked
func handleMsgApproveChecked(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedApproveChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approveChecked", msg.Value.Type())
	}
	return updateDelegation(instruction.Source, tx.Slot, db, client)
}

// handleMsgRevoke handles a MsgRevoke
func handleMsgRevoke(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedRevoke)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "approveChecked", msg.Value.Type())
	}
	return updateDelegation(instruction.Source, tx.Slot, db, client)
}

// handleSetAuthority handles a MsgSetAuthority
func handleSetAuthority(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedSetAuthority)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setAuthority", msg.Value.Type())
	}
	if instruction.Mint != "" {
		return updateToken(instruction.Mint, tx.Slot, db, client)
	}
	return updateTokenAccount(instruction.Account, tx.Slot, db, client)
}

// handleMsgMintTo handles a MsgMintTo
func handleMsgMintTo(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedMintTo)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "mintTo", msg.Value.Type())
	}
	return updateTokenSupply(instruction.Mint, tx.Slot, db, client)
}

// handleMsgMintToChecked handles a MsgMintToChecked
func handleMsgMintToChecked(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedMintToChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "mintToChecked", msg.Value.Type())
	}
	return updateTokenSupply(instruction.Mint, tx.Slot, db, client)
}

// handleBurn handles a MsgBurn
func handleMsgBurn(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedBurn)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "burn", msg.Value.Type())
	}
	return updateTokenSupply(instruction.Mint, tx.Slot, db, client)
}

// handleBurn handles a MsgBurnChecked
func handleMsgBurnChecked(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedBurnChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "burnChecked", msg.Value.Type())
	}
	return updateTokenSupply(instruction.Mint, tx.Slot, db, client)
}

// handleMsgCloseAccount handles a MsgCloseAccount
func handleMsgCloseAccount(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedCloseAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "closeAccount", msg.Value.Type())
	}
	return updateTokenAccount(instruction.Account, tx.Slot, db, client)
}

// handleMsgFreezeAccount handles a MsgFreezeAccount
func handleMsgFreezeAccount(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedFreezeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "freezeAccount", msg.Value.Type())
	}
	return updateTokenAccount(instruction.Account, tx.Slot, db, client)
}

// handleMsgThawAccount handles a MsgThawAccount
func handleMsgThawAccount(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedThawAccount)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "thawAccount", msg.Value.Type())
	}
	return updateTokenAccount(instruction.Account, tx.Slot, db, client)
}
