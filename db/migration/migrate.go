package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP TABLE validator_skip_rate;
	CREATE TABLE validator_skip_rate
	(
		address      TEXT     NOT NULL PRIMARY KEY,
		epoch        BIGINT   NOT NULL,
		skip_rate    FLOAT    NOT NULL,
		total        INT      NOT NULL,
		skip         INT      NOT NULL
	);

	CREATE TABLE validator_skip_rate_history
	(
		address      TEXT     NOT NULL,
		epoch        BIGINT   NOT NULL,
		skip_rate    FLOAT    NOT NULL,
		total        INT      NOT NULL,
		skip         INT      NOT NULL
	);
	ALTER TABLE validator_skip_rate_history ADD UNIQUE (address, epoch);
	CREATE INDEX validator_skip_rate_history_address ON validator_skip_rate_history(address);
	CREATE INDEX validator_skip_rate_history_epoch ON validator_skip_rate_history(epoch);
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP TABLE validator_skip_rate;
	DROP TABLE validator_skip_rate_history;

	CREATE TABLE validator_skip_rate
	(
		address      TEXT     NOT NULL PRIMARY KEY,
		epoch        BIGINT   NOT NULL,
		skip_rate    FLOAT    NOT NULL
	);
	`)
	return err
}
