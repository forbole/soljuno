package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/forbole/soljuno/types/logging"

	"github.com/lib/pq"
	_ "github.com/lib/pq" // nolint

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
		Sql:    postgresDb,
		Logger: ctx.Logger,
	}, nil
}

// type check to ensure interface is properly implemented
var _ db.Database = &Database{}

// Database defines a wrapper around a SQL database and implements functionality
// for data aggregation and exporting.
type Database struct {
	Sql    *sql.DB
	Logger logging.Logger
}

// LastBlockSlot implements db.Database
func (db *Database) LastBlockSlot() (int64, error) {
	var height int64
	err := db.Sql.QueryRow(`SELECT coalesce(MAX(slot),0) AS slot FROM block;`).Scan(&height)
	return height, err
}

// HasBlock implements db.Database
func (db *Database) HasBlock(height uint64) (bool, error) {
	var res bool
	err := db.Sql.QueryRow(`SELECT EXISTS(SELECT 1 FROM block WHERE slot = $1);`, height).Scan(&res)
	return res, err
}

// SaveBlock implements db.Database
func (db *Database) SaveBlock(block types.Block) error {
	sqlStatement := `
INSERT INTO block (slot, hash, proposer, timestamp)
VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING`
	proposer := sql.NullString{Valid: len(block.Proposer) != 0, String: block.Proposer}
	_, err := db.Sql.Exec(sqlStatement,
		block.Slot, block.Hash, proposer, block.Timestamp,
	)
	return err
}

// SaveTx implements db.Database
func (db *Database) SaveTx(tx types.Tx) error {
	sqlStatement := `
INSERT INTO transaction 
    (hash, slot, error, fee, logs) 
VALUES ($1, $2, $3, $4, $5) ON CONFLICT DO NOTHING`
	_, err := db.Sql.Exec(sqlStatement,
		tx.Hash,
		tx.Slot,
		tx.Successful(),
		tx.Fee,
		pq.Array(tx.Logs),
	)
	return err
}

// HasValidator implements db.Database
func (db *Database) HasValidator(pubkey string) (bool, error) {
	var res bool
	stmt := `SELECT EXISTS(SELECT 1 FROM validator WHERE vote_pubkey = $1 OR node_pubkey = $1);`
	err := db.Sql.QueryRow(stmt, pubkey).Scan(&res)
	return res, err
}

// SaveValidators implements db.Database
func (db *Database) SaveValidators(validators []types.Validator) error {
	if len(validators) == 0 {
		return nil
	}
	stmt := `INSERT INTO validator (consensus_address, consensus_pubkey) VALUES `

	var vparams []interface{}
	for i, val := range validators {
		vi := i * 2

		stmt += fmt.Sprintf("($%d, $%d),", vi+1, vi+2)
		vparams = append(vparams, val.VotePubkey, val.NodePubKey)
	}

	stmt = stmt[:len(stmt)-1] // Remove trailing ,
	stmt += " ON CONFLICT DO NOTHING"
	_, err := db.Sql.Exec(stmt, vparams...)
	return err
}

// SaveMessage implements db.Database
func (db *Database) SaveMessage(msg types.Message) error {
	stmt := `
INSERT INTO message(transaction_hash, index, program, involved_accounts, type, value) 
VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Sql.Exec(
		stmt,
		msg.TxHash,
		msg.Index,
		msg.Program,
		pq.Array(msg.InvolvedAccounts),
		msg.Value.Type(),
		msg.Value.JSON(),
	)
	return err
}

// Close implements db.Database
func (db *Database) Close() {
	err := db.Sql.Close()
	if err != nil {
		db.Logger.Error("error while closing connection", "err", err)
	}
}
