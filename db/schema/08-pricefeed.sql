CREATE TABLE token_unit
(
    price_id    TEXT    PRIMARY KEY,
    address     TEXT    NOT NULL UNIQUE,
    unit_name   TEXT    NOT NULL UNIQUE
);

CREATE TABLE token_price
(
    unit_name   TEXT                        NOT NULL REFERENCES token_unit (unit_name) PRIMARY KEY,
    price       DECIMAL                     NOT NULL,
    market_cap  BIGINT                      NOT NULL,
    timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX token_price_timestamp_index ON token_price (timestamp);