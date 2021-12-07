package db

import (
	"database/sql"
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
)

// Database represents an abstract database that can be used to save data inside it
type Database interface {
	BasicDb

	ExceutorDb

	PruningDb

	BankDb

	TokenDb

	SystemDb

	StakeDb

	VoteDb

	BpfLoaderDb

	ConfigDb

	PriceDb

	ConsensusDb

	// Close closes the connection to the database
	Close()
}

type BasicDb interface {
	// HasBlock tells whether or not the database has already stored the block having the given height.
	// An error is returned if the operation fails.
	HasBlock(slot uint64) (bool, error)

	// SaveBlock will be called when a new block is parsed, passing the block itself
	// and the transactions contained inside that block.
	// An error is returned if the operation fails.
	SaveBlock(block types.Block) error

	// SaveTx will be called to save each transaction contained inside a block.
	// An error is returned if the operation fails.
	SaveTx(tx types.Tx) error

	// SaveMessage stores a single message.
	// An error is returned if the operation fails.
	SaveMessage(msg types.Message) error
}

// ExceutorDb represents an abstract database that can excute a raw sql
type ExceutorDb interface {
	// Exec will run the given raw sql
	Exec(string) (sql.Result, error)
}

// PruningDb represents a database that supports pruning properly
type PruningDb interface {
	// Prune prunes the data before the given slot, returning any error
	Prune(slot uint64) error
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
	SaveTokenAccount(address string, slot uint64, mint, owner string) error

	// DeleteTokenAccount allows to delete the given address of the token account inside the database
	DeleteTokenAccount(address string) error

	// SaveMultisig allows to store the given multisig data inside the database
	SaveMultisig(address string, slot uint64, singers []string, m uint8) error

	// SaveDelegate allows to store the given approve state inside the database
	SaveTokenDelegation(source string, destination string, slot uint64, amount uint64) error

	// DeleteTokenDelegation allows to delete the given address of the token delegation inside the database
	DeleteTokenDelegation(address string) error

	// SaveTokenSupply allows to store the given token data inside the database
	SaveTokenSupply(address string, slot uint64, supply uint64) error

	TokenCheckerDb
}

// TokenCheckerDb represents a database that checks account statement of token properly
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

// SystemDb represents a database that checks account statement of system properly
type SystemDb interface {
	// SaveNonceAccount allows to store the given nonce account data inside the database
	SaveNonceAccount(address string, slot uint64, authority string, blockhash string, lamportsPerSignature uint64) error

	// DeleteNonceAccount allows to delete the given address of the nonce account inside the database
	DeleteNonceAccount(address string) error

	SystemCheckerDb
}

// SystemCheckerDb represents a database that checks account statement of system properly
type SystemCheckerDb interface {
	// CheckNonceAccountLatest checks if the nonce account statement is latest
	CheckNonceAccountLatest(address string, currentSlot uint64) bool
}

// StakeDb represents a database that supports stake properly
type StakeDb interface {
	// SaveStakeAccount allows to store the given stake account data inside the database
	SaveStakeAccount(address string, slot uint64, staker string, withdrawer string) error

	// DeleteStakeAccount allows to delete the given address of the stake account inside the database
	DeleteStakeAccount(address string) error

	// SaveStakeLockup allows to store the given stake account lockup state inside the database
	SaveStakeLockup(address string, slot uint64, custodian string, epoch uint64, unixTimestamp int64) error

	// SaveStakeDelegation allows to store the given delegation of stake account inside the database
	SaveStakeDelegation(address string, slot uint64, activationEpoch uint64, deactivationEpoch uint64, stake uint64, voter string, rate float64) error

	// DeleteStakeDelegation allows to delete the given address of the stake delegation inside the database
	DeleteStakeDelegation(address string) error

	StakeCheckerDb
}

// SystemCheckerDb represents a database that checks account statement of system properly
type StakeCheckerDb interface {
	// CheckStakeAccountLatest checks if the stake account statement is latest
	CheckStakeAccountLatest(address string, currentSlot uint64) bool
}

// VoteDb represents a database that supports vote properly
type VoteDb interface {
	// SaveVoteAccount allows to store the given vote account data inside the database
	SaveVoteAccount(address string, slot uint64, node string, voter string, withdrawer string, commission uint8) error

	// SaveValidatorStatus allows to store the given current validator status inside the database
	SaveValidatorStatus(address string, slot uint64, activatedStake uint64, lastVote uint64, rootSlot uint64, active bool) error

	// PruneValidatorStatus allows to delete validator statuses before the given slot
	PruneValidatorStatus(slot uint64) error

	VoteCheckerDb
}

// VoteCheckerDb represents a database that checks account statement of vote properly
type VoteCheckerDb interface {
	// CheckVoteAccountLatest checks if the vote account statement is latest
	CheckVoteAccountLatest(address string, currentSlot uint64) bool
}

// ConfigDb represents a database that supports config properly
type ConfigDb interface {
	// SaveValidatorConfig allows to store the given config account data inside the database
	SaveValidatorConfig(row dbtypes.ValidatorConfigRow) error
}

// BpfLoaderDb represents a database that supports bpf loader properly
type BpfLoaderDb interface {
	// SaveBufferAccount allows to store the given buffer account data inside the database
	SaveBufferAccount(address string, slot uint64, authority string) error

	// DeleteBufferAccount allows to delete the given address of the buffer account inside the database
	DeleteBufferAccount(address string) error

	// SaveProgramAccount allows to store the given program account data inside the database
	SaveProgramAccount(address string, slot uint64, programDataAccount string) error

	// DeleteBufferAccount allows to delete the given address of the program account inside the database
	DeleteProgramAccount(address string) error

	// SaveProgramDataAccount allows to store the given program data account inside the database
	SaveProgramDataAccount(address string, slot uint64, lastModifiedSlot uint64, updateAuthority string) error

	// DeleteBufferAccount allows to delete the given address of the program data account inside the database
	DeleteProgramDataAccount(address string) error

	BpfLoaderCheckerDb
}

// BpfLoaderCheckerDb represents a database that checks account statement of bpf loader properly
type BpfLoaderCheckerDb interface {
	// CheckBufferAccountLatest checks if the buffer account statement is latest
	CheckBufferAccountLatest(address string, currentSlot uint64) bool

	// CheckProgramAccountLatest checks if the program account statement is latest
	CheckProgramAccountLatest(address string, currentSlot uint64) bool

	// CheckProgramDataAccountLatest checks if the program data account statement is latest
	CheckProgramDataAccountLatest(address string, currentSlot uint64) bool
}

// PricesDb represents a database that supports pricefeed properly
type PriceDb interface {
	// GetTokenUnits returns the slice of all the names of the different tokens units
	GetTokenUnits() ([]types.TokenUnit, error)

	// SaveTokenUnit allows to save the given token unit details
	SaveTokenUnit(unit types.TokenUnit) error

	// SaveTokensPrices allows to store the token prices inside the database
	SaveTokensPrices(prices []types.TokenPrice) error
}

// ConsensusDb represents a database that supports consesus properly
type ConsensusDb interface {
	// GetLastBlock allows to get the last block
	GetLastBlock() (dbtypes.BlockRow, error)

	// GetBlockHourAgo allows to get the latest block before a hour ago inside the database
	GetBlockHourAgo(now time.Time) (dbtypes.BlockRow, error)

	// SaveAverageSlotTimePerHour allows to store the average slot time inside the database
	SaveAverageSlotTimePerHour(slot uint64, averageTime float64) error
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
