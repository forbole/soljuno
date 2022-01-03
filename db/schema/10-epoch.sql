CREATE TABLE supply_info (
    one_row_id      BOOL            NOT NULL DEFAULT TRUE PRIMARY KEY,
    epoch           BIGINT          NOT NULL,
    total           NUMERIC(20,0)   NOT NULL,
    circulating     NUMERIC(20,0)   NOT NULL,
    non_circulating NUMERIC(20,0)   NOT NULL,
    CHECK (one_row_id)
);
