package token

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/token"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

// HandleMsg allows to handle different messages types for the token module
func HandleMsg(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	switch msg.Value.Type() {
	case "initializeMint":
	case "initializeMint2":
		return handleMsgInitializeMint(msg, tx, db)

	case "initializeAccount":
	case "initializeAccount2":
	case "initializeAccount3":
		return handleMsgInitializeAccount(msg, tx, db)

	case "initializeMultisig":
	case "initializeMultisig2":
		return handleMsgInitializeMultisig(msg, tx, db)

	// TODO: make decision if handle token balance instead of bank module
	case "transfer":
	case "transferChecked":
		return nil

	// Delegate msgs
	case "approve":
		return handleMsgApprove(msg, db, client)
	case "approveChecked":
		return handleMsgApproveChecked(msg, db, client)
	case "revoke":
		return handleMsgRevoke(msg, db, client)

	case "setAuthority":
		return nil

	// Token supply msgs
	case "mintTo":
		return handleMsgMintTo(msg, db, client)
	case "mintToChecked":
		return handleMsgMintToChecked(msg, db, client)
	case "burn":
		return handleMsgBurn(msg, db, client)
	case "burnChecked":
		return handleMsgBurnChecked(msg, db, client)

	// Account state msgs
	case "closeAccount":
		return handleMsgCloseAccount(msg, db, client)
	case "freezeAccount":
		return handleMsgFreezeAccount(msg, db, client)
	case "thawAccount":
		return handleMsgThawAccount(msg, db, client)
	}

	log.Info().Str("module", "token").Str("message", msg.Value.Type()).Uint64("slot", tx.Slot).
		Msg("handled message")
	return nil
}

// handleMsgInitializeMint handles a MsgInitializeMint and properly stores the new token inside the database
func handleMsgInitializeMint(msg types.Message, tx types.Tx, db db.TokenDb) error {
	instruction, ok := msg.Value.Data().(token.ParsedInitializeMint)
	if !ok {
		return fmt.Errorf("instruction does not match initializeMint type: %s", msg.Value.Type())
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
	return nil
}

// handleMsgInitializeAccount handles a MsgInitializeAccount and properly stores the new token account inside the database
func handleMsgInitializeAccount(msg types.Message, tx types.Tx, db db.TokenDb) error {
	instruction, ok := msg.Value.Data().(token.ParsedInitializeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match initializeAccount type: %s", msg.Value.Type())
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
		return fmt.Errorf("instruction does not match initializeMultisig type: %s", msg.Value.Type())
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
func handleMsgApprove(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedApprove)
	if !ok {
		return fmt.Errorf("instruction does not match approve type: %s", msg.Value.Type())
	}
	return updateDelegation(instruction.Source, db, client)
}

// handleMsgApproveChecked handles a MsgApproveChecked
func handleMsgApproveChecked(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedApproveChecked)
	if !ok {
		return fmt.Errorf("instruction does not match approveChecked type: %s", msg.Value.Type())
	}
	return updateDelegation(instruction.Source, db, client)
}

// handleMsgRevoke handles a MsgRevoke
func handleMsgRevoke(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedRevoke)
	if !ok {
		return fmt.Errorf("instruction does not match approveChecked type: %s", msg.Value.Type())
	}
	return updateDelegation(instruction.Source, db, client)
}

// handleMsgMintTo handles a MsgMintTo
func handleMsgMintTo(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedMintTo)
	if !ok {
		return fmt.Errorf("instruction does not match mintTo type: %s", msg.Value.Type())
	}
	return updateTokenSupply(instruction.Mint, db, client)
}

// handleMsgMintToChecked handles a MsgMintToChecked
func handleMsgMintToChecked(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedMintToChecked)
	if !ok {
		return fmt.Errorf("instruction does not match mintToChecked type: %s", msg.Value.Type())
	}
	return updateTokenSupply(instruction.Mint, db, client)
}

// handleBurn handles a MsgBurn
func handleMsgBurn(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedBurn)
	if !ok {
		return fmt.Errorf("instruction does not match burn type: %s", msg.Value.Type())
	}
	return updateTokenSupply(instruction.Mint, db, client)
}

// handleBurn handles a MsgBurnChecked
func handleMsgBurnChecked(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedBurn)
	if !ok {
		return fmt.Errorf("instruction does not match burnChecked type: %s", msg.Value.Type())
	}
	return updateTokenSupply(instruction.Mint, db, client)
}

// handleMsgCloseAccount handles a MsgCloseAccount
func handleMsgCloseAccount(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedCloseAccount)
	if !ok {
		return fmt.Errorf("instruction does not match closeAccount type: %s", msg.Value.Type())
	}
	return updateAccountState(instruction.Account, db, client)
}

// handleMsgFreezeAccount handles a MsgFreezeAccount
func handleMsgFreezeAccount(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedFreezeAccount)
	if !ok {
		return fmt.Errorf("instruction does not match freezeAccount type: %s", msg.Value.Type())
	}
	return updateAccountState(instruction.Account, db, client)
}

// handleMsgThawAccount handles a MsgThawAccount
func handleMsgThawAccount(msg types.Message, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedThawAccount)
	if !ok {
		return fmt.Errorf("instruction does not match thawAccount type: %s", msg.Value.Type())
	}
	return updateAccountState(instruction.Account, db, client)
}
