CREATE TABLE validator_config
(
    address             TEXT    NOT NULL PRIMARY KEY,
    slot                BIGINT  NOT NULL,
    owner               TEXT    NOT NULL,
    name                TEXT    NOT NULL DEFAULT '',
    keybase_username    TEXT    NOT NULL DEFAULT '',
    website             TEXT    NOT NULL DEFAULT '',
    details             TEXT    NOT NULL DEFAULT ''

);
CREATE INDEX validator_config_owner_index ON validator_config (owner);
