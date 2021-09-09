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
		return fmt.Errorf("instruction does not match initialize mint type: %s", msg.Value.Type())
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
