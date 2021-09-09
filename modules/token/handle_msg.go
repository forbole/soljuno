package token

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/token"
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

func HandleMsg(msg types.Message, tx types.Tx, db db.TokenDb) error {
	switch msg.Value.Type() {
	case "initializeMint":
	case "initializeMint2":
		handleMsgInitializeMint(msg, tx, db)

	case "initializeAccount":
	case "initializeAccount2":
	case "initializeAccount3":
		handleMsgInitializeAccount(msg, tx, db)

	case "initializeMultisig":
	case "initializeMultisig2":

	}
	log.Info().Str("module", "token").Str("message", msg.Value.Type()).Uint64("slot", tx.Slot).
		Msg("handled message")
	return nil
}

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
