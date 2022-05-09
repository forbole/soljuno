package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
		DROP INDEX token_account_balance_index ON token_account_balance;

		CREATE MATERIALIZED VIEW token_account_balance_ordering AS 
		SELECT tab.address, tab.balance, ta.mint FROM token_account_balance AS tab 
		INNER JOIN token_account AS ta ON ta.address = tab.address;

		CREATE INDEX CONCURRENTLY token_account_balance_ordering_index ON token_account_balance_ordering(mint, balance DESC);
	`)

	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
		DROP INDEX token_account_balance_ordering_index;
		DROP INDEX MATERIALIZED VIEW token_account_balance_ordering;
		CREATE INDEX token_account_balance_index ON token_account_balance(balance DESC);
	`)
	return err
}
