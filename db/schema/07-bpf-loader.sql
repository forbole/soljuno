CREATE TABLE buffer_account
(
    address     TEXT    NOT NULL PRIMARY KEY,
    slot        BIGINT  NOT NULL,
    authority   TEXT    NOT NULL
);
CREATE INDEX buffer_account_authority_index ON buffer_account (authority);

CREATE TABLE program_account
(
    address                 TEXT    NOT NULL PRIMARY KEY,
    slot                    BIGINT  NOT NULL,
    program_data_account    TEXT    NOT NULL
);

CREATE TABLE program_data_account
(
    address                 TEXT    NOT NULL PRIMARY KEY,
    slot                    BIGINT  NOT NULL,
    last_modified_slot      BIGINT  NOT NULL,
    update_authority        TEXT    NOT NULL
);
CREATE INDEX program_data_account_authority_index ON program_data_account (update_authority);
