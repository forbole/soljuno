CREATE TABLE token
(
    address             TEXT    NOT NULL PRIMARY KEY,
    slot                BIGINT  NOT NULL,
    decimals            INT     NOT NULL,
    mint_authority      TEXT    NOT NULL,
    freeze_authority    TEXT    NOT NULL
);
CREATE INDEX token_authority_index ON token (mint_authority);

CREATE TABLE token_supply
(
    address TEXT            NOT NULL PRIMARY KEY,
    slot    BIGINT          NOT NULL,
    supply  NUMERIC(20,0)   NOT NULL
);

CREATE TABLE token_account
(
    address TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    token   TEXT    NOT NULL,
    owner   TEXT    NOT NULL
);
CREATE INDEX token_account_owner_index ON token_account (owner);

CREATE TABLE multisig
(
    address TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    signers TEXT[]  NOT NULL,
    m       INT     NOT NULL
);

CREATE TABLE token_delegation
(
    source_address      TEXT    NOT NULL PRIMARY KEY,
    delegate_address    TEXT    NOT NULL,
    slot                BIGINT  NOT NULL,
    amount              BIGINT  NOT NULL
);