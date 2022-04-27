package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
		ALTER TABLE transaction ADD COLUMN involved_accounts TEXT[] NOT NULL DEFAULT array[]::TEXT[];
		CREATE INDEX transaction_accounts_index ON transaction USING GIN(involved_accounts);

		CREATE FUNCTION transactions_by_address(
			addresses TEXT[],
			"start_slot" BIGINT = 0,
			"end_slot" BIGINT = 0
			)
			RETURNS SETOF transaction AS
		$$
			SELECT * FROM transaction WHERE 
			(slot < "end_slot" AND slot >= "start_slot") AND
			involved_accounts @> addresses ORDER BY slot+0 DESC
		$$ LANGUAGE sql STABLE;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP FUNCTION transactions_by_address;
	DROP INDEX transaction_accounts_index;
	ALTER TABLE transaction DROP COLUMN involved_accounts;
	`)
	return err
}
