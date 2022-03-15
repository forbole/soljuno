package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP TABLE message CASCADE;
	CREATE TABLE instruction
	(
		tx_signature   			TEXT    NOT NULL,
		slot                	BIGINT  NOT NULL,
		index               	INT     NOT NULL,
		inner_index         	INT     NOT NULL,
		program             	TEXT    NOT NULL,      
		involved_accounts   	TEXT[]  NOT NULL DEFAULT array[]::TEXT[],
		raw_data            	TEXT    NOT NULL,
		type                	TEXT    NOT NULL DEFAULT 'unknown',
		value               	JSON    NOT NULL DEFAULT '{}',
		partition_id        	INT     NOT NULL,
		CHECK (slot / 1000 = partition_id)
	) PARTITION BY LIST(partition_id);
	ALTER TABLE instruction ADD UNIQUE (tx_signature, index, inner_index, partition_id);
	CREATE INDEX instruction_tx_signature_index ON instruction (tx_signature);
	CREATE INDEX instruction_slot_index ON instruction (slot DESC);
	CREATE INDEX instruction_program_index ON instruction (program);
	CREATE INDEX instruction_accounts_index ON instruction USING GIN(involved_accounts);

	CREATE FUNCTION instructions_by_address(
		addresses TEXT[],
		programs TEXT[],
		"limit" BIGINT = 100,
		"offset" BIGINT = 0)
		RETURNS SETOF instruction AS
	$$
	SELECT 
		instruction.tx_signature, instruction.slot, instruction.index, instruction.inner_index, instruction.program, instruction.involved_accounts, instruction.raw_data, instruction.type, instruction.value, instruction.partition_id
	FROM instruction
	WHERE (cardinality(programs) = 0 OR program = ANY (programs))
	  AND involved_accounts @> addresses
	ORDER BY slot DESC,
	involved_accounts LIMIT "limit" OFFSET "offset"
	$$ LANGUAGE sql STABLE;

	ALTER TABLE block RENAME COLUMN proposer TO leader;
	ALTER TABLE transaction RENAME COLUMN hash TO signature;
	ALTER TABLE token RENAME COLUMN address TO mint;
	ALTER TABLE token_supply RENAME COLUMN address TO mint;
	ALTER TABLE token_unit RENAME COLUMN address TO mint;
	ALTER TABLE token_account RENAME COLUMN token TO mint;
	`)
	return err
}

func Down(db db.ExcecutorDb) error {
	_, err := db.Exec(`
	DROP TABLE instruction CASCADE;
	CREATE TABLE message
	(
		tx_signature    	TEXT    NOT NULL,
		slot                BIGINT  NOT NULL,
		index               INT     NOT NULL,
		inner_index         INT     NOT NULL,
		program             TEXT    NOT NULL,      
		involved_accounts   TEXT[]  NOT NULL DEFAULT array[]::TEXT[],
		raw_data            TEXT    NOT NULL,
		type                TEXT    NOT NULL DEFAULT 'unknown',
		value               JSON    NOT NULL DEFAULT '{}',
		partition_id        INT     NOT NULL,
		CHECK (slot / 1000 = partition_id)
	) PARTITION BY LIST(partition_id);
	ALTER TABLE message ADD UNIQUE (tx_signature, index, inner_index, partition_id);
	CREATE INDEX message_tx_signature_index ON message (tx_signature);
	CREATE INDEX message_slot_index ON message (slot DESC);
	CREATE INDEX message_program_index ON message (program);
	CREATE INDEX message_accounts_index ON message USING GIN(involved_accounts);

	/**
	* This function is used to find all the utils that involve any of the given addresses and have
	* type that is one of the specified types.
	*/
	CREATE FUNCTION messages_by_address(
		addresses TEXT[],
		programs TEXT[],
		"limit" BIGINT = 100,
		"offset" BIGINT = 0)
		RETURNS SETOF message AS
	$$
	SELECT 
		message.tx_signature, message.slot, message.index, message.inner_index, message.program, message.involved_accounts, message.raw_data, message.type, message.value, message.partition_id
	FROM message
	WHERE (cardinality(programs) = 0 OR program = ANY (programs))
	AND involved_accounts @> addresses
	ORDER BY slot DESC,
	involved_accounts LIMIT "limit" OFFSET "offset"
	$$ LANGUAGE sql STABLE;

	ALTER TABLE block RENAME COLUMN leader TO proposer;
	ALTER TABLE transaction RENAME COLUMN signature TO hash;
	ALTER TABLE token RENAME COLUMN mint TO address;
	ALTER TABLE token_supply RENAME COLUMN mint TO address;
	ALTER TABLE token_unit RENAME COLUMN mint TO address;
	ALTER TABLE token_account RENAME COLUMN mint TO token;
	`)
	return err
}
