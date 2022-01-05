package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	DROP TABLE validator_status;
	CREATE TABLE validator_status
	(
		address         TEXT    NOT NULL PRIMARY KEY,
		slot            BIGINT  NOT NULL,
		activated_stake BIGINT  NOT NULL,
		last_vote       BIGINT  NOT NULL,
		root_slot       BIGINT  NOT NULL,
		active          BOOLEAN NOT NULL
	);
	`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`
	DROP TABLE validator_status;
	CREATE TABLE validator_status
	(
		address         TEXT    NOT NULL,
		slot            BIGINT  NOT NULL,
		activated_stake BIGINT  NOT NULL,
		last_vote       BIGINT  NOT NULL,
		root_slot       BIGINT  NOT NULL,
		active          BOOLEAN NOT NULL,
		PRIMARY KEY (address, slot)
	);
	CREATE INDEX vote_account_voter_index ON validator_status (address);
	CREATE INDEX vote_account_slot_index ON validator_status (slot);`)
	return err
}
