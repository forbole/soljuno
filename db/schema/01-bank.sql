CREATE TABLE account_balance
(
    address TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    balance BIGINT  NOT NULL
);

CREATE TABLE token_account_balance
(
    address TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    balance BIGINT  NOT NULL
);