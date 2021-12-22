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
INSERT INTO block (slot, height, hash, proposer, timestamp, num_tx)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING`
	proposer := sql.NullString{Valid: len(block.Proposer) != 0, String: block.Proposer}
	_, err := db.Sqlx.Exec(
		stmt, block.Slot, block.Height, block.Hash, proposer, block.Timestamp, len(block.Txs),
	)
	return err
}

// SaveTx implements db.Database
func (db *Database) SaveTx(tx types.Tx) error {
	stmt := `
INSERT INTO transaction (hash, slot, error, fee, logs)
VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	_, err := db.Sqlx.Exec(
		stmt,
		tx.Hash,
		tx.Slot,
		tx.Successful(),
		tx.Fee,
		pq.Array(tx.Logs),
	)
	return err
}

// SaveTxs implements db.Database
func (db *Database) SaveTxs(txs []types.Tx) error {
	if len(txs) == 0 {
		return nil
	}
	stmt := `INSERT INTO transaction (hash, slot, error, fee, logs) VALUES`
	var params []interface{}
	paramsNumber := 5
	for i, tx := range txs {
		bi := i * paramsNumber
		stmt += fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d),",
			bi+1, bi+2, bi+3, bi+4, bi+5,
		)
		params = append(
			params,
			tx.Hash,
			tx.Slot,
			tx.Successful(),
			tx.Fee,
			pq.Array(tx.Logs),
		)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += `ON CONFLICT DO NOTHING`
	_, err := db.Sqlx.Exec(stmt, params...)
	return err
}

// SaveMessage implements db.Database
func (db *Database) SaveMessage(msg types.Message) error {
	stmt := `
INSERT INTO message(transaction_hash, slot, index, inner_index, program, involved_accounts, raw_data, type, value) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT DO NOTHING`
	if msg.InvolvedAccounts == nil {
		msg.InvolvedAccounts = []string{}
	}
	_, err := db.Sqlx.Exec(
		stmt,
		msg.TxHash,
		msg.Slot,
		msg.Index,
		msg.InnerIndex,
		msg.Program,
		pq.Array(msg.InvolvedAccounts),
		msg.RawData,
		msg.Parsed.Type(),
		msg.Parsed.JSON(),
	)
	return err
}

// SaveMessages implements db.Database
func (db *Database) SaveMessages(msgs []types.Message) error {
	if len(msgs) == 0 {
		return nil
	}
	stmt := `INSERT INTO message
	(transaction_hash, slot, index, inner_index, program, involved_accounts, raw_data, type, value) VALUES`

	var params []interface{}
	paramsNumber := 9
	for i, msg := range msgs {
		bi := i * paramsNumber
		stmt += fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d),",
			bi+1, bi+2, bi+3, bi+4, bi+5, bi+6, bi+7, bi+8, bi+9,
		)
		if msg.InvolvedAccounts == nil {
			msg.InvolvedAccounts = []string{}
		}
		params = append(
			params,
			msg.TxHash,
			msg.Slot,
			msg.Index,
			msg.InnerIndex,
			msg.Program,
			pq.Array(msg.InvolvedAccounts),
			msg.RawData,
			msg.Parsed.Type(),
			msg.Parsed.JSON(),
		)
	}

	stmt = stmt[:len(stmt)-1]
	stmt += `ON CONFLICT DO NOTHING`
	_, err := db.Sqlx.Exec(stmt, params...)
	return err
}

// Close implements db.Database
func (db *Database) Close() {
	err := db.Sqlx.Close()
	if err != nil {
		db.Logger.Error("error while closing connection", "err", err)
	}
}
