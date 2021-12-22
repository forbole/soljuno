package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	ALTER TABLE block ADD COLUMN num_tx INT NOT NULL DEFAULT 0
	`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`ALTER TABLE block DROP COLUMN num_tx`)
	return err
}
