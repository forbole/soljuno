CREATE TABLE nonce
(
    address                 TEXT    NOT NULL PRIMARY KEY,
    slot                    BIGINT  NOT NULL,
    authority               TEXT    NOT NULL,
    blockhash               TEXT    NOT NULL,
    lamports_per_signature  INT     NOT NULL,
    state                   TEXT    NOT NULL
);