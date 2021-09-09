package db

import (
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
)

// Database represents an abstract database that can be used to save data inside it
type Database interface {
	// HasBlock tells whether or not the database has already stored the block having the given height.
	// An error is returned if the operation fails.
	HasBlock(slot uint64) (bool, error)

	// SaveBlock will be called when a new block is parsed, passing the block itself
	// and the transactions contained inside that block.
	// An error is returned if the operation fails.
	// NOTE. For each transaction inside txs, SaveTx will be called as well.
	SaveBlock(block types.Block) error

	// SaveTx will be called to save each transaction contained inside a block.
	// An error is returned if the operation fails.
	SaveTx(tx types.Tx) error

	// SaveMessage stores a single message.
	// An error is returned if the operation fails.
	SaveMessage(msg types.Message) error

	// Close closes the connection to the database
	Close()
}

// PruningDb represents a database that supports pruning properly
type PruningDb interface {
	// Prune prunes the data for the given slot, returning any error
	Prune(slot uint64) error

	// StoreLastPruned saves the last slot at which the database was pruned
	StoreLastPruned(slot uint64) error

	// GetLastPruned returns the last slot at which the database was pruned
	GetLastPruned() (uint64, error)
}

// BankDb represents a database that supports bank properly
type BankDb interface {
	// SaveAccountBalances allows to store the given native balance data inside the database
	SaveAccountBalances(slot uint64, accounts []string, balances []uint64) error

	// SaveAccountBalances allows to store the given token balance data inside the database
	SaveAccountTokenBalances(slot uint64, accounts []string, balances []clienttypes.TransactionTokenBalance) error
}

type TokenDb interface {
	SaveToken(mint string, slot uint64, decimals uint8, mintAuthority string, freezeAuthority string) error

	SaveTokenAccount(address string, slot uint64, mint, owner string) error
}

// Context contains the data that might be used to build a Database instance
type Context struct {
	Cfg    types.DatabaseConfig
	Logger logging.Logger
}

// NewContext allows to build a new Context instance
func NewContext(cfg types.DatabaseConfig, logger logging.Logger) *Context {
	return &Context{
		Cfg:    cfg,
		Logger: logger,
	}
}

// Builder represents a method that allows to build any database from a given codec and configuration
type Builder func(ctx *Context) (Database, error)
