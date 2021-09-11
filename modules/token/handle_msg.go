package token

import (
	"encoding/base64"
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/account"
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

	case "approve":
		return handleMsgApprove(msg, tx, db, client)
	case "approveChecked":
		return handleMsgApproveChecked(msg, tx, db, client)

	}

	log.Info().Str("module", "token").Str("message", msg.Value.Type()).Uint64("slot", tx.Slot).
		Msg("handled message")
	return nil
}

//____________________________________________________________________________

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

//____________________________________________________________________________

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
	)
	if err != nil {
		return err
	}
	return nil
}

//____________________________________________________________________________

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

//____________________________________________________________________________

// handleMsgApproveChecked handles a MsgApprove
func handleMsgApprove(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedApprove)
	if !ok {
		return fmt.Errorf("instruction does not match approve type: %s", msg.Value.Type())
	}
	return handleApproveMsgs(instruction.Source, tx.Slot, db, client)
}

// handleMsgApproveChecked handles a MsgApproveChecked
func handleMsgApproveChecked(msg types.Message, tx types.Tx, db db.TokenDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(token.ParsedApproveChecked)
	if !ok {
		return fmt.Errorf("instruction does not match approve type: %s", msg.Value.Type())
	}
	return handleApproveMsgs(instruction.Source, tx.Slot, db, client)
}

// handleApproveMsgs handles approve messages and properly stores the statement of approve inside the database
func handleApproveMsgs(source string, slot uint64, db db.TokenDb, client client.Proxy) error {
	info, err := client.AccountInfo(source)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.SaveDelegate(source, "", slot, 0)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	tokenAccount, ok := account.Parse(token.ProgramID, bz).(account.TokenAccount)
	if !ok {
		return db.SaveDelegate(source, "", slot, 0)
	}

	if !tokenAccount.Delegate.Option.Bool() {
		return db.SaveDelegate(source, "", slot, 0)
	}

	return db.SaveDelegate(source, tokenAccount.Delegate.Value.String(), slot, tokenAccount.DelegateAmount)
}

//____________________________________________________________________________
