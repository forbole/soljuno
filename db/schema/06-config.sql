CREATE TABLE config_account
(
    address     TEXT    NOT NULL PRIMARY KEY,
    slot        BIGINT  NOT NULL,
    owner       TEXT    NOT NULL,
    value       JSONB   NOT NULL DEFAULT '{}'::JSONB
);
CREATE INDEX config_account_owner_index ON config_account (owner);
