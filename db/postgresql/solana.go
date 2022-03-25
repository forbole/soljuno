package postgresql

import (
	"database/sql"
	"fmt"

	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/types/logging"
	"github.com/jmoiron/sqlx"

	"github.com/forbole/soljuno/db"
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
func (db *Database) SaveBlock(block dbtypes.BlockRow) error {
	stmt := `
INSERT INTO block (slot, height, hash, leader, timestamp, num_txs)
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT DO NOTHING`
	leader := sql.NullString{Valid: len(block.Leader) != 0, String: block.Leader}
	_, err := db.Sqlx.Exec(
		stmt, block.Slot, block.Height, block.Hash, leader, block.Timestamp, block.NumTxs,
	)
	return err
}

// createPartition allows to create a partition with the id for the given table name
func (db *Database) createPartition(table string, id int) error {
	stmt := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %v_%d PARTITION OF %v FOR VALUES IN (%d)",
		table,
		id,
		table,
		id,
	)
	_, err := db.Exec(stmt)
	return err
}

// dropPartition allows to drop a partition with the given partition name
func (db *Database) dropPartition(name string) error {
	stmt := fmt.Sprintf(
		"DROP TABLE IF EXISTS %v",
		name,
	)
	_, err := db.Exec(stmt)
	return err
}

// Close implements db.Database
func (db *Database) Close() {
	err := db.Sqlx.Close()
	if err != nil {
		db.Logger.Error("error while closing connection", "err", err)
	}
}
