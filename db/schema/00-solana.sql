CREATE TABLE block
(
    slot        BIGINT      PRIMARY KEY,
    height      BIGINT      NOT NULL,
    hash        TEXT        NOT NULL UNIQUE,
    proposer    TEXT        DEFAULT '',
    timestamp   TIMESTAMP   WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX block_hash_index ON block (hash);
CREATE INDEX block_proposer_index ON block (proposer);
CREATE INDEX block_timestamp_index ON block (timestamp);


CREATE TABLE transaction
(
    hash       TEXT     NOT NULL PRIMARY KEY,
    slot       BIGINT   NOT NULL REFERENCES block (slot),
    error      BOOLEAN  NOT NULL,
    fee        INT      NOT NULL,
    logs       TEXT[]
);
CREATE INDEX transaction_slot_index ON transaction (slot);

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
