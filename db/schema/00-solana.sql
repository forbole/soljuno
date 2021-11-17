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
    index               INT     NOT NULL,
    inner_index         INT     NOT NULL,
    program             TEXT    NOT NULL,      
    involved_accounts   TEXT[]  NOT NULL DEFAULT array[]::TEXT[],
    type                TEXT    NOT NULL DEFAULT 'unknown',
    value               JSONB   NOT NULL DEFAULT '{}'::JSONB
);
CREATE INDEX message_transaction_hash_index ON message (transaction_hash);

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
SELECT message.transaction_hash, message.index, message.program, message.involved_accounts, message.type, message.value
FROM message
         JOIN transaction t on message.transaction_hash = t.hash
WHERE (cardinality(types) = 0 OR type = ANY (types))
  AND addresses && involved_accounts
ORDER BY slot DESC
LIMIT "limit" OFFSET "offset"
$$ LANGUAGE sql STABLE;

CREATE TABLE pruning
(
    last_pruned_slot BIGINT NOT NULL
);