package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
		ALTER TABLE transaction ADD COLUMN index INT NOT NULL DEFAULT 0;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER TABLE transaction DROP COLUMN index;
	`)
	return err
}
