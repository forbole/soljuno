CREATE TABLE buffer_account
(
    address     TEXT    NOT NULL PRIMARY KEY,
    slot        BIGINT  NOT NULL,
    authority   TEXT    NOT NULL
);

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
    last_modified_slot      BIGINT  NOT Null,
    update_authority        TEXT    NOT NULL
);

