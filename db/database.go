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

// TokenDb represents a database that supports token properly
type TokenDb interface {
	// SaveToken allows to store the given token data inside the database
	SaveToken(address string, slot uint64, decimals uint8, mintAuthority string, freezeAuthority string) error

	// SaveTokenAccount allows to store the given token account data inside the database
	SaveTokenAccount(address string, slot uint64, mint, owner, state string) error

	// SaveMultisig allows to store the given multisig data inside the database
	SaveMultisig(address string, slot uint64, singers []string, m uint8) error

	// SaveDelegate allows to store the given approve state inside the database
	SaveTokenDelegate(source string, destination string, slot uint64, amount uint64) error

	// SaveTokenSupply allows to store the given token data inside the database
	SaveTokenSupply(address string, slot uint64, supply uint64) error

	TokenCheckerDb
}

// TokenCheckerDb represents a database that checks if token statement is latest
type TokenCheckerDb interface {
	// CheckTokenLatest checks if the token statement is latest
	CheckTokenLatest(address string, currentSlot uint64) bool

	// CheckTokenAccountLatest checks if the token account statement is latest
	CheckTokenAccountLatest(address string, currentSlot uint64) bool

	// CheckMultisigLatest checks if the multisig statement is latest
	CheckMultisigLatest(address string, currentSlot uint64) bool

	// CheckTokenDelegateLatest checks delegate statement
	CheckTokenDelegateLatest(address string, currentSlot uint64) bool

	// CheckTokenSupplyLatest checks if the token supply statement is latest
	CheckTokenSupplyLatest(address string, currentSlot uint64) bool
}

// SystemDb represents a database that supports system properly
type SystemDb interface {
	// SaveNonceAccount allows to store the given nonce account data inside the database
	SaveNonceAccount(address string, slot uint64, authority string, blockhash string, lamportsPerSignature uint64, state string) error
}

// StakeDb represents a database that supports stake properly
type StakeDb interface {
	// SaveStakeAccount allows to store the given stake account data inside the database
	SaveStakeAccount(address string, slot uint64, staker string, withdrawer string, state string) error

	// SaveStakeLockup allows to store the given stake account lockup state inside the database
	SaveStakeLockup(address string, slot uint64, custodian string, epoch uint64, unixTimestamp int64) error

	// SaveStakeDelegation allows to store the given delegation of stake account inside the database
	SaveStakeDelegation(address string, slot uint64, activationEpoch uint64, deactivationEpoch uint64, stake uint64, voter string, rate float64) error
}

// VoteDb represents a database that supports vote properly
type VoteDb interface {
	// SaveVoteAccount allows to store the given vote account data inside the database
	SaveVoteAccount(address string, slot uint64, node string, voter string, withdrawer string, commission uint8) error
}

// ConfigDb represents a database that supports config properly
type ConfigDb interface {
	// SaveConfigAccount allows to store the given config account data inside the database
	SaveConfigAccount(address string, slot uint64, owner string, data string) error
}

// BpfLoaderDb represents a database that supports bpf loader properly
type BpfLoaderDb interface {
	// SaveBufferAccount allows to store the given buffer account data inside the database
	SaveBufferAccount(address string, slot uint64, authority string, state string) error

	// SaveProgramAccount allows to store the given program account data inside the database
	SaveProgramAccount(address string, slot uint64, programDataAccount string, state string) error

	// SaveProgramDataAccount allows to store the given program data account inside the database
	SaveProgramDataAccount(address string, slot uint64, lastModifiedSlot uint64, updateAuthority string, state string) error
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
