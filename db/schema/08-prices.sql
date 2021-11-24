CREATE TABLE token_unit
(
    address    TEXT NOT NULL UNIQUE,
    token_name TEXT NOT NULL UNIQUE
);

CREATE TABLE token_price
(
    unit_name  TEXT                        NOT NULL REFERENCES token_unit (token_name) UNIQUE,
    price      DECIMAL                     NOT NULL,
    market_cap BIGINT                      NOT NULL,
    timestamp  TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX token_price_timestamp_index ON token_price (timestamp);