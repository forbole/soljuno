CREATE TABLE block
(
    slot        BIGINT      PRIMARY KEY,
    height      BIGINT      NOT NULL,
    hash        TEXT        NOT NULL UNIQUE,
    leader    TEXT        NOT NULL DEFAULT '',
    timestamp   TIMESTAMP   WITHOUT TIME ZONE NOT NULL,
    num_txs     INT         NOT NULL DEFAULT 0
);
CREATE INDEX block_hash_index ON block (hash);
CREATE INDEX block_leader_index ON block (leader);
CREATE INDEX block_timestamp_index ON block (timestamp DESC);


CREATE TABLE transaction
(
    signature           TEXT    NOT NULL,
    slot                BIGINT  NOT NULL,
    index               INT     NOT NULL DEFAULT 0,
    involved_accounts   TEXT[]  NOT NULL DEFAULT array[]::TEXT[],
    success             BOOLEAN NOT NULL,
    fee                 INT     NOT NULL,
    logs                TEXT[],
    num_instructions    INT     NOT NULL DEFAULT 0,
    partition_id        INT     NOT NULL,
    CHECK (slot / 1000 = partition_id)
) PARTITION BY LIST(partition_id);
ALTER TABLE transaction ADD UNIQUE (signature, partition_id);
CREATE INDEX transaction_signature_index ON transaction (signature);
CREATE INDEX transaction_slot_index ON transaction (slot DESC);
CREATE INDEX transaction_accounts_index ON transaction USING GIN(involved_accounts);

CREATE TABLE transaction_by_address
(
    address         TEXT    NOT NULL,
    slot            BIGINT  NOT NULL, 
    signature       TEXT    NOT NULL,
    index           INT     NOT NULL DEFAULT 0,
    partition_id    INT     NOT NULL
) PARTITION BY LIST(partition_id);
ALTER TABLE transaction ADD UNIQUE (address, signature, partition_id);
CREATE INDEX transaction_by_address_index ON transaction_by_address(address);
CREATE INDEX transaction_by_address_signature_index ON transaction_by_address(signature);
CREATE INDEX transaction_by_address_slot_index ON transaction_by_address (slot DESC);

CREATE TABLE instruction
(
    tx_signature        TEXT    NOT NULL,
    slot                BIGINT  NOT NULL,
    index               INT     NOT NULL,
    inner_index         INT     NOT NULL,
    program             TEXT    NOT NULL,      
    involved_accounts   TEXT[]  NOT NULL DEFAULT array[]::TEXT[],
    raw_data            TEXT    NOT NULL,
    type                TEXT    NOT NULL DEFAULT 'unknown',
    value               JSON    NOT NULL DEFAULT '{}',
    partition_id        INT     NOT NULL,
    CHECK (slot / 1000 = partition_id)
) PARTITION BY LIST(partition_id);
ALTER TABLE instruction ADD UNIQUE (tx_signature, index, inner_index, partition_id);
CREATE INDEX instruction_tx_signature_index ON instruction (tx_signature);
CREATE INDEX instruction_slot_index ON instruction (slot DESC);
CREATE INDEX instruction_program_index ON instruction (program);
CREATE INDEX instruction_accounts_index ON instruction USING GIN(involved_accounts);

/**
 * This function is used to find all the utils that involve any of the given addresses and have
 * type that is one of the specified types.
 */
CREATE FUNCTION transactions_by_address(
    addresses TEXT[],
    "start_slot" BIGINT = 0,
    "end_slot" BIGINT = 0
    )
    RETURNS SETOF transaction AS
$$
    SELECT * FROM transaction WHERE 
    (slot <= "end_slot" AND slot >= "start_slot") AND
    involved_accounts @> addresses ORDER BY slot+0 DESC;
$$ LANGUAGE sql STABLE;

CREATE OR REPLACE FUNCTION transactions_by_address_2(
    "target"    TEXT,
    "current"   TEXT = '',
    "limit"     INT = 10
    )
    RETURNS SETOF transaction AS
$$
    BEGIN
        IF "current" = '' THEN
            RETURN QUERY SELECT * FROM transaction WHERE
            EXISTS (SELECT 1 FROM transaction_by_address WHERE address = "target") AND
            signature IN ( 
                SELECT signature FROM transaction_by_address WHERE address = "target" ORDER BY slot DESC, index DESC LIMIT "limit"
            );
        ELSE    
            RETURN QUERY WITH slot_getter AS (
                /* slot_filter returns the tx current slot */
                SELECT DISTINCT(slot) FROM transaction_by_address 
                    WHERE signature = "current"
                ), 
                /* index_getter returns the current tx index */
                index_getter AS (
                    SELECT DISTINCT(index) FROM transaction_by_address 
                    WHERE signature = "current"
                ),
                /* account_checker returns if the signature exists behind the current tx slot and index */
                account_checker AS (
                    SELECT signature FROM transaction_by_address WHERE slot <= (
                        SELECT slot FROM slot_getter
                    ) AND 
                    index < (
                        SELECT index FROM index_getter
                    ) AND address = "target"
                )
                ,
                /* account_signatures_getter returns the signatures filtered by the account behind the current tx slot and index */
                account_signatures_getter AS (
                    SELECT signature FROM transaction_by_address WHERE slot <= (
                        SELECT slot FROM slot_getter
                    ) AND 
                    index < (
                        SELECT index FROM index_getter
                    ) AND address = "target" ORDER BY slot DESC, index DESC LIMIT "limit"
                )
                /* main query */  
                SELECT * FROM transaction WHERE EXISTS (
                    SELECT 1 FROM account_checker
                ) AND
                signature IN ( 
                    SELECT signature FROM account_signatures_getter
                );
        END IF;
    END;
$$ LANGUAGE plpgsql;

/**
 * This function is used to find all the utils that involve any of the given addresses and have
 * type that is one of the specified types.
 */
CREATE FUNCTION instructions_by_address(
    addresses TEXT[],
    programs TEXT[],
    "start_slot" BIGINT = 0,
    "end_slot" BIGINT = 0
    )
    RETURNS SETOF instruction AS
$$
SELECT 
    instruction.tx_signature, instruction.slot, instruction.index, instruction.inner_index, instruction.program, instruction.involved_accounts, instruction.raw_data, instruction.type, instruction.value, instruction.partition_id
FROM (
    SELECT * FROM instruction WHERE 
    (slot <= "end_slot" AND slot >= "start_slot") AND
    (cardinality(programs) = 0 OR program = ANY (programs)) AND 
    involved_accounts @> addresses
    ) as instruction 
$$ LANGUAGE sql STABLE;

