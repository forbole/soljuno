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
    hash            TEXT    NOT NULL,
    slot            BIGINT  NOT NULL,
    error           BOOLEAN NOT NULL,
    fee             INT     NOT NULL,
    logs            TEXT[],
    messages        JSON    NOT NULL DEFAULT '{}',
    partition_id    INT     NOT NULL,
    CHECK (slot / 1000 = partition_id)
) PARTITION BY LIST(partition_id);
ALTER TABLE transaction ADD UNIQUE (hash, partition_id);
CREATE INDEX transaction_hash_index ON transaction (hash);
CREATE INDEX transaction_slot_index ON transaction (slot);
