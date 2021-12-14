CREATE TABLE epoch_info (
    one_row_id          BOOL            NOT NULL DEFAULT TRUE PRIMARY KEY,
    epoch               BIGINT          NOT NULL,
    transaction_count   NUMERIC(20, 0)  NOT NULL,
    CHECK (one_row_id)
);

CREATE TABLE supply_info (
    one_row_id      BOOL            NOT NULL DEFAULT TRUE PRIMARY KEY,
    epoch           BIGINT          NOT NULL,
    total           NUMERIC(20,0)   NOT NULL,
    circulating     NUMERIC(20,0)   NOT NULL,
    non_circulating NUMERIC(20,0)   NOT NULL,
    CHECK (one_row_id)
);

CREATE TABLE inflation_rate (
    one_row_id  BOOL    NOT NULL DEFAULT TRUE PRIMARY KEY,
    epoch       BIGINT  NOT NULL,
    total       FLOAT   NOT NULL,
    foundation  FLOAT   NOT NULL,
    validator   FLOAT   NOT NULL,
    CHECK (one_row_id)
);

CREATE TABLE epoch_schedule_param (
    one_row_id          BOOL    NOT NULL DEFAULT TRUE PRIMARY KEY,
    epoch               BIGINT  NOT NULL,
    slots_per_epoch     INT     NOT NULL,
    first_normal_epoch  INT     NOT NULL,
    first_normal_slot   INT     NOT NULL,
    warmup              BOOL    NOT NULL,
    CHECK (one_row_id)
);

CREATE TABLE inflation_governor_param (
    one_row_id          BOOL    NOT NULL DEFAULT TRUE PRIMARY KEY,
    epoch               BIGINT  NOT NULL,
    initial             FLOAT   NOT NULL,
    terminal            FLOAT   NOT NULL,
    taper               FLOAT   NOT NULL,
    foundation          FLOAT   NOT NULL,
    foundation_terminal FLOAT   NOT NULL,
    CHECK (one_row_id)
);
