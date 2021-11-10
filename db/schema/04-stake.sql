CREATE TABLE stake_account
(
    address     TEXT    NOT NULL PRIMARY KEY,
    slot        BIGINT  NOT NULL,
    staker      TEXT    NOT NULL,
    withdrawer  TEXT    NOT NULL,
    state       TEXT    NOT NULL
);
CREATE INDEX stake_staker_index ON stake_account (withdrawer);

CREATE TABLE stake_lockup
(
    address         TEXT        NOT NULL PRIMARY KEY REFERENCES stake_account (address) ON DELETE CASCADE,
    slot            BIGINT      NOT NULL,
    custodian       TEXT        NOT NULL,
    epoch           BIGINT      NOT NULL,
    unix_timestamp  TIMESTAMP   WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE stake_delegation
(
    address                 TEXT            NOT NULL PRIMARY KEY REFERENCES stake_account (address) ON DELETE CASCADE,
    slot                    BIGINT          NOT NULL,
    activation_epoch        NUMERIC(20,0)   NOT NULL,
    deactivation_epoch      NUMERIC(20,0)   NOT NULL,
    stake                   BIGINT          NOT NULL,
    voter                   TEXT            NOT NULL,
    warmup_cooldown_rate    FLOAT           NOT NULL
);