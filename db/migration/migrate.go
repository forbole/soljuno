package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER INDEX block_timestamp_index RENAME TO block_timestamp_old_index;
	CREATE INDEX block_timestamp_index ON block (timestamp DESC);
	DROP INDEX block_timestamp_old_index;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	ALTER INDEX block_timestamp_index RENAME TO block_timestamp_old_index;
	CREATE INDEX block_timestamp_index ON block (timestamp DESC);
	DROP INDEX block_timestamp_old_index;
	`)
	return err
}
