CREATE TABLE block
(
    slot        BIGINT      PRIMARY KEY,
    hash        TEXT        NOT NULL UNIQUE,
    proposer    TEXT        DEFAULT '',
    timestamp   TIMESTAMP   WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX block_hash_index ON block (hash);
CREATE INDEX block_proposer_index ON block (proposer);

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
    transaction_hash    TEXT    NOT NULL REFERENCES transaction (hash),
    slot                BIGINT  NOT NULL REFERENCES block (slot),
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


/**
 * This function is used to find all the utils that involve any of the given addresses and have
 * type that is one of the specified types.
 */
CREATE FUNCTION messages_by_address(
    addresses TEXT[],
    types TEXT[],
    "limit" BIGINT = 100,
    "offset" BIGINT = 0)
    RETURNS SETOF message AS
$$
SELECT 
    message.transaction_hash, message.slot, message.index, message.inner_index, message.program, message.involved_accounts, message.raw_data, message.type, message.value
FROM message
WHERE (cardinality(types) = 0 OR type = ANY (types))
  AND addresses && involved_accounts
ORDER BY slot DESC
LIMIT "limit" OFFSET "offset"
$$ LANGUAGE sql STABLE;
