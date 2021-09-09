CREATE TABLE token
(
    mint                TEXT    NOT NULL PRIMARY KEY,
    slot                BIGINT  NOT NULL,
    decimals            INT     NOT NULL,
    supply              BIGINT  NOT NULL,
    mint_authority      TEXT,
    freeze_authority    TEXT
)
CREATE INDEX token_slot_index ON token (slot);

CREATE TABLE token_account
(
    address TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    mint    TEXT    NOT NULL,
    owner   TEXT    NOT NULL,
    balance BIGINT  NOT NULL
);
CREATE INDEX token_account_slot_index ON token_account (slot);

CREATE TABLE multisig 
(
    address TEXT    NOT NULL PRIMARY KEY,
    signers TEXT[]  NOT NULL,
    m       INT     NOT NULL
);

CREATE TABLE token_delegate
(
    source_address      TEXT    NOT NULL,
    delegate_address    TEXT    NOT NULL,
    amount              BIGINT  NOT NULL
);