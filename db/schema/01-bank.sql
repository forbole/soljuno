CREATE TABLE account_balance
(
    address TEXT    NOT NULL PRIMARY KEY,
    slot    BIGINT  NOT NULL,
    balance BIGINT  NOT NULL
);

CREATE TABLE token_account_balance
(
    address TEXT            NOT NULL PRIMARY KEY,
    slot    BIGINT          NOT NULL,
    balance NUMERIC(20,0)   NOT NULL
);

CREATE TABLE account_balance_history
(
    address     TEXT                        NOT NULL,
    timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    balance     BIGINT                      NOT NULL
);
CREATE INDEX account_balance_history_address ON account_balance_history(address);
CREATE INDEX account_balance_history_timestamp ON account_balance_history(timestamp);

CREATE TABLE token_account_balance_history
(
    address     TEXT                        NOT NULL,
    timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    balance     NUMERIC(20,0)               NOT NULL
);
CREATE INDEX token_account_balance_history_address ON token_account_balance_history(address);
CREATE INDEX token_account_balance_history_timestamp ON token_account_balance_history(timestamp);