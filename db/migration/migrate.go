package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
		CREATE INDEX CONCURRENTLY token_account_balance_index ON token_account_balance(balance DESC);
		CREATE INDEX CONCURRENTLY token_account_mint_index ON token_account (mint);
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP INDEX token_account_balance_index;
	DROP INDEX token_account_mint_index;
	`)
	return err
}
