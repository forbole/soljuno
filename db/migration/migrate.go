package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	ALTER TABLE token_delegation ALTER COLUMN amount TYPE NUMERIC(20,0);
	`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE token_delegation ALTER COLUMN amount TYPE BIGINT;
	`)
	return err
}
