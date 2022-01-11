package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	DROP TABLE message;
	DROP TABLE message_by_address;
	`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`
	CREATE TABLE message
	(
		transaction_hash    TEXT    NOT NULL,
		slot                BIGINT  NOT NULL,
		index               INT     NOT NULL,
		inner_index         INT     NOT NULL,
		program             TEXT    NOT NULL,      
		raw_data            TEXT    NOT NULL,
		type                TEXT    NOT NULL DEFAULT 'unknown',
		value               JSONB   NOT NULL DEFAULT '{}'::JSONB
	);
	CREATE INDEX message_transaction_hash_index ON message (transaction_hash);
	CREATE INDEX message_slot_index ON message (slot);
	CREATE INDEX message_program_index ON message (program);

	CREATE TABLE message_by_address
	(
		address             TEXT    NOT NULL,
		slot                BIGINT  NOT NULL,
		transaction_hash    TEXT    NOT NULL,
		index               INT     NOT NULL,
		inner_index         INT     NOT NULL
	);
	CREATE INDEX message_by_address_address_index ON message_by_address (address);
	CREATE INDEX message_by_address_slot_index ON message_by_address (slot);
	CREATE INDEX message_by_transaction_hash_slot_index ON message_by_address (transaction_hash);
	`)
	return err
}
