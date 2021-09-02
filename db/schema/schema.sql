CREATE TABLE validator
(
    vote_pubkey TEXT NOT NULL PRIMARY KEY,
    node_pubkey  TEXT NOT NULL UNIQUE 
);

CREATE TABLE block
(
    slot        BIGINT UNIQUE PRIMARY KEY,
    hash        TEXT   NOT NULL UNIQUE,
    proposer    TEXT   DEFAULT '',
    timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX block_hash_index ON block (hash);
CREATE INDEX block_proposer_index ON block (proposer);


CREATE TABLE transaction
(
    hash       TEXT     NOT NULL UNIQUE PRIMARY KEY,
    slot       BIGINT   NOT NULL REFERENCES block (slot),
    error      BOOLEAN  NOT NULL,
    fee        INT      NOT NULL,
    logs       TEXT[]   NOT NULL
);
CREATE INDEX transaction_hash_index ON transaction (hash);
CREATE INDEX transaction_slot_index ON transaction (slot);

CREATE TABLE message
(
    transaction_hash    TEXT    NOT NULL REFERENCES transaction (hash),
    index               BIGINT  NOT NULL,
    program             TEXT    NOT NULL,      
    involved_accounts   TEXT[]  NOT NULL,
    type                TEXT    NOT NULL DEFAULT 'unknown',
    value               JSONB   NOT NULL DEFAULT '{}'::JSONB
);
CREATE INDEX message_transaction_hash_index ON message (transaction_hash);

CREATE TABLE pruning
(
    last_pruned_slot BIGINT NOT NULL
)