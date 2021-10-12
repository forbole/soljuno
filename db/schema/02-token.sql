CREATE TABLE token
(
    mint                TEXT    NOT NULL PRIMARY KEY,
    slot                BIGINT  NOT NULL,
    decimals            INT     NOT NULL,
    mint_authority      TEXT,
    freeze_authority    TEXT
);
CREATE INDEX token_authority_index ON token (mint_authority);

CREATE TABLE token_supply
(
    mint    TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    supply  BIGINT  NOT NULL
);

CREATE TABLE token_account
(
    address TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    mint    TEXT    NOT NULL,
    owner   TEXT    NOT NULL,
    state   TEXT    NOT NULL
);
CREATE INDEX token_account_owner_index ON token_account (owner);

CREATE TABLE multisig
(
    address TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    signers TEXT[]  NOT NULL,
    m       INT     NOT NULL
);

CREATE TABLE token_delegate
(
    source_address      TEXT    NOT NULL PRIMARY KEY,
    delegate_address    TEXT    NOT NULL,
    slot                BIGINT  NOT NULL,
    amount              BIGINT  NOT NULL
);