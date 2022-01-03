package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/forbole/soljuno/types/logging"
	"github.com/jmoiron/sqlx"

	"github.com/lib/pq"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

const MAX_PARAMS_LENGTH = 65535

// Builder creates a database connection with the given database connection info
// from config. It returns a database connection handle or an error if the
// connection fails.
func Builder(ctx *db.Context) (db.Database, error) {
	sslMode := "disable"
	if ctx.Cfg.GetSSLMode() != "" {
		sslMode = ctx.Cfg.GetSSLMode()
	}

	schema := "public"
	if ctx.Cfg.GetSchema() != "" {
		schema = ctx.Cfg.GetSchema()
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s sslmode=%s search_path=%s",
		ctx.Cfg.GetHost(), ctx.Cfg.GetPort(), ctx.Cfg.GetName(), ctx.Cfg.GetUser(), sslMode, schema,
	)

	if ctx.Cfg.GetPassword() != "" {
		connStr += fmt.Sprintf(" password=%s", ctx.Cfg.GetPassword())
	}

	postgresDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Set max open connections
	postgresDb.SetMaxOpenConns(ctx.Cfg.GetMaxOpenConnections())
	postgresDb.SetMaxIdleConns(ctx.Cfg.GetMaxIdleConnections())

	return &Database{
		Sqlx:   sqlx.NewDb(postgresDb, "postgresql"),
		Logger: ctx.Logger,
	}, nil
}

// type check to ensure interface is properly implemented
var _ db.Database = &Database{}

// Database defines a wrapper around a SQL database and implements functionality
// for data aggregation and exporting.
type Database struct {
	Sqlx   *sqlx.DB
	Logger logging.Logger
}

// LastBlockSlot implements db.Database
func (db *Database) LastBlockSlot() (int64, error) {
	var height int64
	err := db.Sqlx.QueryRow(`SELECT coalesce(MAX(slot),0) AS slot FROM block;`).Scan(&height)
	return height, err
}

// HasBlock implements db.Database
func (db *Database) HasBlock(height uint64) (bool, error) {
	var res bool
	err := db.Sqlx.QueryRow(`SELECT EXISTS(SELECT 1 FROM block WHERE slot = $1);`, height).Scan(&res)
	return res, err
}

// SaveBlock implements db.Database
func (db *Database) SaveBlock(block types.Block) error {
	stmt := `
INSERT INTO block (slot, height, hash, proposer, timestamp, num_txs)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING`
	proposer := sql.NullString{Valid: len(block.Proposer) != 0, String: block.Proposer}
	_, err := db.Sqlx.Exec(
		stmt, block.Slot, block.Height, block.Hash, proposer, block.Timestamp, len(block.Txs),
	)
	return err
}

// SaveTxs implements db.Database
func (db *Database) SaveTxs(txs []types.Tx) error {
	if len(txs) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO transaction (hash, slot, error, fee, logs) VALUES`
	paramsStmt := ""
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 5
	for i, tx := range txs {
		bi := i * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)
		params = append(
			params,
			tx.Hash,
			tx.Slot,
			tx.Successful(),
			tx.Fee,
			pq.Array(tx.Logs),
		)
	}
	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}

// SaveMessages implements db.Database
func (db *Database) SaveMessages(msgs []types.Message) error {
	if len(msgs) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO message
	(transaction_hash, slot, index, inner_index, program, type) VALUES`
	paramsStmt := ""
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 6
	for i, msg := range msgs {
		bi := i * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)
		params = append(
			params,
			msg.TxHash,
			msg.Slot,
			msg.Index,
			msg.InnerIndex,
			msg.Program,
			msg.Parsed.Type(),
		)
	}
	err := db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
	if err != nil {
		return nil
	}
	return db.saveMsgAddressIndexes(msgs)
}

// saveMsgAddressIndexes implements db.Database
func (db *Database) saveMsgAddressIndexes(msgs []types.Message) error {
	if len(msgs) == 0 {
		return nil
	}
	var params []interface{}
	paramsNumber := 5
	count := 0
	insertStmt := `INSERT INTO message_by_address
	(address, slot, transaction_hash, index, inner_index) VALUES`
	paramsStmt := ""
	conflictStmt := `ON CONFLICT DO NOTHING`

	for _, msg := range msgs {
		for _, account := range msg.InvolvedAccounts {
			// Excute if the max params length will be reached
			if len(params)+paramsNumber >= MAX_PARAMS_LENGTH {
				err := db.insertWithParams(
					insertStmt,
					paramsStmt[:len(paramsStmt)-1],
					conflictStmt,
					params,
				)
				if err != nil {
					return err
				}
				count = 0
				paramsStmt = ""
				params = params[:0]
			}

			bi := count * paramsNumber
			paramsStmt += getParamsStmt(bi, paramsNumber)
			params = append(
				params,
				account,
				msg.Slot,
				msg.TxHash,
				msg.Index,
				msg.InnerIndex,
			)
			count++
		}
	}
	if count == 0 {
		return nil
	}
	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}

// Close implements db.Database
func (db *Database) Close() {
	err := db.Sqlx.Close()
	if err != nil {
		db.Logger.Error("error while closing connection", "err", err)
	}
}
