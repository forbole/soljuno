CREATE TABLE token_unit
(
    mint        TEXT    NOT NULL PRIMARY KEY,
    price_id    TEXT    NOT NULL DEFAULT '',
    unit_name   TEXT    NOT NULL DEFAULT '',
    logo_uri    TEXT    NOT NULL DEFAULT '',
    description TEXT    NOT NULL DEFAULT '',
    website     TEXT    NOT NULL DEFAULT ''
);
CREATE INDEX token_unit_price_id_index ON token_unit (price_id);

CREATE TABLE token_price
(
    id          TEXT                        NOT NULL PRIMARY KEY,
    price       DECIMAL                     NOT NULL,
    market_cap  BIGINT                      NOT NULL,
    symbol      TEXT                        NOT NULL,
    volume      BIGINT                      NOT NULL DEFAULT 0,
    timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE token_price_history
(
    id          TEXT                        NOT NULL,
    price       DECIMAL                     NOT NULL,
    market_cap  BIGINT                      NOT NULL,
    symbol      TEXT                        NOT NULL,
    volume      BIGINT                      NOT NULL DEFAULT 0,
    timestamp   TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE INDEX token_price_history_id_index ON token_price_history (id);
CREATE INDEX token_price_history_timestamp_index ON token_price_history (timestamp DESC);

