package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER INDEX transaction_slot_index RENAME TO transaction_slot_old_index;
	CREATE INDEX transaction_slot_index ON transaction (slot DESC);
	DROP INDEX transaction_slot_old_index;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER INDEX transaction_slot_index RENAME TO transaction_slot_old_index;
	CREATE INDEX transaction_slot_index ON transaction (slot);
	DROP INDEX transaction_slot_old_index;
	`)
	return err
}
