CREATE TABLE token_unit
(
    address     TEXT    NOT NULL PRIMARY KEY,
    price_id    TEXT    NOT NULL UNIQUE,
    unit_name   TEXT    NOT NULL,
    logo_url    TEXT    NOT NULL DEFAULT ''
);

CREATE TABLE token_price
(
    id          TEXT                        NOT NULL REFERENCES token_unit (price_id) PRIMARY KEY,
    price       DECIMAL                     NOT NULL,
    market_cap  BIGINT                      NOT NULL,
    symbol      TEXT                        NOT NULL,
    timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX token_price_timestamp_index ON token_price (timestamp);

