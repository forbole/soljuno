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
    program             TEXT    NOT NULL,      
    involved_accounts   TEXT[]  NOT NULL DEFAULT array[]::TEXT[],
    type                TEXT    NOT NULL DEFAULT 'unknown',
    value               JSONB   NOT NULL DEFAULT '{}'::JSONB
);
CREATE INDEX message_transaction_hash_index ON message (transaction_hash);

CREATE TABLE pruning
(
    last_pruned_slot BIGINT NOT NULL
);