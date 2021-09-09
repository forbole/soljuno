package bank

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"

	"github.com/rs/zerolog/log"
)

func HandleTx(tx types.Tx, db db.BankDb) error {
	if !tx.Successful() {
		return nil
	}

	err := db.SaveAccountBalances(tx.Slot, tx.Accounts, tx.PostBalances)
	if err != nil {
		return err
	}

	err = db.SaveAccountTokenBalances(tx.Slot, tx.Accounts, tx.PostTokenBalances)
	if err != nil {
		return err
	}

	log.Info().Str("module", "bank").Str("tx", tx.Hash).Uint64("slot", tx.Slot).
		Msg("handled tx")
	return nil
}
