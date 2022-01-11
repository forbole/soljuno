CREATE TABLE block
(
    slot        BIGINT      PRIMARY KEY,
    height      BIGINT      NOT NULL,
    hash        TEXT        NOT NULL UNIQUE,
    proposer    TEXT        DEFAULT '',
    timestamp   TIMESTAMP   WITHOUT TIME ZONE NOT NULL,
    num_txs     INT         NOT NULL DEFAULT 0
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
    logs       TEXT[],
    messages   JSONB    NOT NULL DEFAULT '{}'
);
CREATE INDEX transaction_slot_index ON transaction (slot);
