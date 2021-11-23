CREATE TABLE vote_account
(
    address     TEXT    NOT NULL PRIMARY KEY,
    slot        BIGINT  NOT NULL,
    node        TEXT    NOT NULL,
    voter       TEXT    NOT NULL,
    withdrawer  TEXT    NOT NULL,
    commission  INT     NOT NULL
);
CREATE INDEX vote_account_node_index ON vote_account (node);

CREATE TABLE validator_status
(
    address         TEXT    NOT NULL,
    slot            BIGINT  NOT NULL,
    activated_stake BIGINT  NOT NULL,
    last_vote       BIGINT  NOT NULL,
    root_slot       BIGINT  NOT NULL,
    active          BOOLEAN NOT NULL,
    PRIMARY KEY (address, slot)
);
CREATE INDEX vote_account_voter_index ON validator_status (address);
CREATE INDEX vote_account_slot_index ON validator_status (slot);
