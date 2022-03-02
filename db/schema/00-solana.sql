CREATE TABLE block
(
    slot        BIGINT      PRIMARY KEY,
    height      BIGINT      NOT NULL,
    hash        TEXT        NOT NULL UNIQUE,
    proposer    TEXT        NOT NULL DEFAULT '',
    timestamp   TIMESTAMP   WITHOUT TIME ZONE NOT NULL,
    num_txs     INT         NOT NULL DEFAULT 0
);
CREATE INDEX block_hash_index ON block (hash);
CREATE INDEX block_proposer_index ON block (proposer);
CREATE INDEX block_timestamp_index ON block (timestamp);


CREATE TABLE transaction
(
    hash            TEXT    NOT NULL,
    slot            BIGINT  NOT NULL,
    error           BOOLEAN NOT NULL,
    fee             INT     NOT NULL,
    logs            TEXT[],
    partition_id    INT     NOT NULL,
    CHECK (slot / 1000 = partition_id)
) PARTITION BY LIST(partition_id);
ALTER TABLE transaction ADD UNIQUE (hash, partition_id);
CREATE INDEX transaction_hash_index ON transaction (hash);
CREATE INDEX transaction_slot_index ON transaction (slot DESC);


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
    value               JSON    NOT NULL DEFAULT '{}',
    partition_id        INT     NOT NULL,
    CHECK (slot / 1000 = partition_id)
) PARTITION BY LIST(partition_id);
ALTER TABLE message ADD UNIQUE (transaction_hash, index, inner_index, partition_id);
CREATE INDEX message_transaction_hash_index ON message (transaction_hash);
CREATE INDEX message_slot_index ON message (slot);
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
    message.transaction_hash, message.slot, message.index, message.inner_index, message.program, message.involved_accounts, message.raw_data, message.type, message.value, message.partition_id
FROM message
WHERE (cardinality(programs) = 0 OR program = ANY (programs))
  AND involved_accounts @> addresses
ORDER BY slot DESC,
involved_accounts LIMIT "limit" OFFSET "offset"
$$ LANGUAGE sql STABLE;