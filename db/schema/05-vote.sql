CREATE TABLE vote_account
(
    address     TEXT    NOT NULL PRIMARY KEY,
    slot        BIGINT  NOT NULL,
    node        TEXT    NOT NULL,
    voter       TEXT    NOT NULL,
    withdrawer  TEXT    NOT NULL,
    commission  INT     NOT NULL
);