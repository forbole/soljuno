package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP TABLE transaction;
	CREATE TABLE transaction
	(
		hash            TEXT    NOT NULL,
		slot            BIGINT  NOT NULL,
		error           BOOLEAN NOT NULL,
		fee             INT     NOT NULL,
		logs            TEXT[],
		messages        JSON    NOT NULL DEFAULT '{}',
		partition_id    INT     NOT NULL,
		CHECK (slot / 1000 = partition_id)
	) PARTITION BY LIST(partition_id);
	ALTER TABLE transaction ADD UNIQUE (hash, partition_id);
	CREATE INDEX transaction_hash_index ON transaction (hash);
	CREATE INDEX transaction_slot_index ON transaction (slot);
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP TABLE transaction;
	CREATE TABLE transaction
	(
		hash       TEXT     NOT NULL PRIMARY KEY,
		slot       BIGINT   NOT NULL REFERENCES block (slot),
		error      BOOLEAN  NOT NULL,
		fee        INT      NOT NULL,
		logs       TEXT[],
		messages   JSONB    NOT NULL DEFAULT '{}'
	);
	CREATE INDEX transaction_slot_index ON transaction (slot);
	`)
	return err
}
