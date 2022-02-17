CREATE TABLE validator
(
    address     TEXT    NOT NULL PRIMARY KEY,
    slot        BIGINT  NOT NULL,
    node        TEXT    NOT NULL,
    voter       TEXT    NOT NULL,
    withdrawer  TEXT    NOT NULL,
    commission  INT     NOT NULL
);
CREATE INDEX validator_node_index ON validator (node);

CREATE TABLE validator_status
(
    address         TEXT    NOT NULL PRIMARY KEY,
    slot            BIGINT  NOT NULL,
    activated_stake BIGINT  NOT NULL,
    last_vote       BIGINT  NOT NULL,
    root_slot       BIGINT  NOT NULL,
    active          BOOLEAN NOT NULL
);

CREATE TABLE validator_skip_rate
(
    address      TEXT     NOT NULL PRIMARY KEY,
    epoch        BIGINT   NOT NULL,
    skip_rate    FLOAT    NOT NULL,
    total        INT      NOT NULL,
    skip         INT      NOT NULL
);

CREATE TABLE validator_skip_rate_history
(
    address      TEXT     NOT NULL,
    epoch        BIGINT   NOT NULL,
    skip_rate    FLOAT    NOT NULL,
    total        INT      NOT NULL,
    skip         INT      NOT NULL
);
ALTER TABLE validator_skip_rate_history ADD UNIQUE (address, epoch);
CREATE INDEX validator_skip_rate_history_address ON validator_skip_rate_history(address);
CREATE INDEX validator_skip_rate_history_epoch ON validator_skip_rate_history(epoch);
