package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`
	DROP TABLE message CASCADE;
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

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`
	DROP TABLE message;
	DROP TABLE message_by_address;
	CREATE TABLE message
	(
		transaction_hash    TEXT    NOT NULL,
		slot                BIGINT  NOT NULL,
		index               INT     NOT NULL,
		inner_index         INT     NOT NULL,
		program             TEXT    NOT NULL,      
		involved_accounts   TEXT[]  NOT NULL DEFAULT array[]::TEXT[],
		raw_data            TEXT    NOT NULL,
		type                TEXT    NOT NULL DEFAULT 'unknown',
		value               JSONB   NOT NULL DEFAULT '{}'::JSONB
	);
	CREATE INDEX message_transaction_hash_index ON message (transaction_hash);
	CREATE INDEX message_slot_index ON message (slot);
	CREATE INDEX message_program_index ON message (program);
	CREATE INDEX message_accounts_index ON message USING GIN(involved_accounts);
	ALTER TABLE message ALTER COLUMN involved_accounts SET STATISTICS 1000;
	ANALYZE message;

	/**
	* This function is used to find all the utils that involve any of the given addresses and have
	* type that is one of the specified types.
	*/
	CREATE FUNCTION messages_by_address(
		addresses TEXT[],
		types TEXT[],
		programs TEXT[],
		"limit" BIGINT = 100,
		"offset" BIGINT = 0)
		RETURNS SETOF message AS
	$$
	SELECT 
		message.transaction_hash, message.slot, message.index, message.inner_index, message.program, message.involved_accounts, message.raw_data, message.type, message.value
	FROM message
	WHERE (cardinality(types) = 0 OR type = ANY (types))
	AND (cardinality(programs) = 0 OR program = ANY (programs))
	AND involved_accounts @> addresses
	ORDER BY slot DESC,
	involved_accounts LIMIT "limit" OFFSET "offset"
	$$ LANGUAGE sql STABLE;
	`)
	return err
}
