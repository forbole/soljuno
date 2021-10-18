package bank

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

func HandleTx(tx types.Tx, db db.BankDb) error {
	err := db.SaveAccountBalances(tx.Slot, tx.Accounts, tx.PostBalances)
	if err != nil {
		return err
	}

	return db.SaveAccountTokenBalances(tx.Slot, tx.Accounts, tx.PostTokenBalances)
}
