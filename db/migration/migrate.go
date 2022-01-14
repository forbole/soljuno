package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	CREATE TABLE validator_skip_rate
	(
		address      TEXT     NOT NULL PRIMARY KEY,
		epoch        BIGINT   NOT NULL,
		skip_rate    FLOAT    NOT NULL
	);
	`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`
	DROP TABLE validator_skip_rate;
	`)
	return err
}
