CREATE TABLE nonce_account
(
    address                 TEXT    NOT NULL PRIMARY KEY,
    slot                    BIGINT  NOT NULL,
    authority               TEXT    NOT NULL,
    blockhash               TEXT    NOT NULL,
    lamports_per_signature  INT     NOT NULL,
    state                   TEXT    NOT NULL
);
CREATE INDEX nonce_authority_index ON nonce (authority);