CREATE TABLE average_slot_time_per_hour
(
    one_row_id      BOOL    NOT NULL DEFAULT TRUE PRIMARY KEY,
    slot            BIGINT  NOT NULL,
    average_time    DECIMAL NOT NULL,
    CHECK (one_row_id)
);;