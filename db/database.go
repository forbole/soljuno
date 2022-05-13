package db

import (
	"database/sql"
	"time"

	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
)

// Database represents an abstract database that can be used to save data inside it
type Database interface {
	BlockDb

	TxDb

	InstructionDb

	ExcecutorDb

	BankDb

	TokenDb

	SystemDb

	StakeDb

	VoteDb

	VoteStatusDb

	BpfLoaderDb

	ConfigDb

	PriceDb

	ConsensusDb

	EpochDb

	FixMissingBlockDb

	// Close closes the connection to the database
	Close()
}

type BlockDb interface {
	// HasBlock tells whether or not the database has already stored the block having the given height.
	// An error is returned if the operation fails.
	HasBlock(slot uint64) (bool, error)

	// SaveBlock will be called when a new block is parsed, passing the block itself
	// and the transactions contained inside that block.
	// An error is returned if the operation fails.
	SaveBlock(block dbtypes.BlockRow) error
}

type TxDb interface {
	// SaveTxs stores a batch of transactions.
	// An error is returned if the operation fails.
	SaveTxs(txs []dbtypes.TxRow) error

	// CreateTxPartition allows to create a new transaction table partition with the given partition id
	CreateTxPartition(ID int) error

	// PruneTxsBeforeSlot allows to prune the txs before the given slot
	PruneTxsBeforeSlot(slot uint64) error
}

type InstructionDb interface {
	// SaveInstructions stores a batch of instructions.
	// An error is returned if the operation fails.
	SaveInstructions(instructions []dbtypes.InstructionRow) error

	// CreateInstructionPartition allows to create a instruction partition
	CreateInstructionPartition(Id int) error

	// PruneInstructionsBeforeSlot allows to prune the instructions before the given slot
	PruneInstructionsBeforeSlot(slot uint64) error
}

// ExcecutorDb represents an abstract database that can excute a raw sql
type ExcecutorDb interface {
	// Exec will run the given raw sql
	Exec(sql string, args ...interface{}) (sql.Result, error)

	// Query will run the given query sql
	Query(sql string, args ...interface{}) (*sql.Rows, error)
}

// BankDb represents a database that supports bank properly
type BankDb interface {
	// SaveAccountBalances allows to store the given native balance data inside the database
	SaveAccountBalances(slot uint64, accounts []string, balances []uint64) error

	// SaveAccountBalances allows to store the given token balance data inside the database
	SaveAccountTokenBalances(slot uint64, accounts []string, balances []uint64) error

	// SaveAccountHistoryBalances allows to store the given historical native balance data inside the database
	SaveAccountHistoryBalances(time time.Time, accounts []string, balances []uint64) error

	// SaveAccountHistoryTokenBalances allows to store the given historical token balance data inside the database
	SaveAccountHistoryTokenBalances(time time.Time, accounts []string, balances []uint64) error
}

// TokenDb represents a database that supports token properly
type TokenDb interface {
	// SaveToken allows to store the given token data inside the database
	SaveToken(token dbtypes.TokenRow) error

	// SaveTokenAccount allows to store the given token account data inside the database
	SaveTokenAccount(account dbtypes.TokenAccountRow) error

	// DeleteTokenAccount allows to delete the given address of the token account inside the database
	DeleteTokenAccount(address string) error

	// SaveMultisig allows to store the given multisig data inside the database
	SaveMultisig(multisig dbtypes.MultisigRow) error

	// SaveDelegate allows to store the given approve state inside the database
	SaveTokenDelegation(delegation dbtypes.TokenDelegationRow) error

	// DeleteTokenDelegation allows to delete the given address of the token delegation inside the database
	DeleteTokenDelegation(address string) error

	// SaveTokenSupply allows to store the given token data inside the database
	SaveTokenSupply(supply dbtypes.TokenSupplyRow) error

	TokenCheckerDb
}

// TokenCheckerDb represents a database that checks account statement of token properly
type TokenCheckerDb interface {
	// CheckTokenLatest checks if the token statement is latest
	CheckTokenLatest(mint string, currentSlot uint64) bool

	// CheckTokenAccountLatest checks if the token account statement is latest
	CheckTokenAccountLatest(address string, currentSlot uint64) bool

	// CheckMultisigLatest checks if the multisig statement is latest
	CheckMultisigLatest(address string, currentSlot uint64) bool

	// CheckTokenDelegateLatest checks delegate statement
	CheckTokenDelegateLatest(address string, currentSlot uint64) bool

	// CheckTokenSupplyLatest checks if the token supply statement is latest
	CheckTokenSupplyLatest(mint string, currentSlot uint64) bool
}

// SystemDb represents a database that checks account statement of system properly
type SystemDb interface {
	// SaveNonceAccount allows to store the given nonce account data inside the database
	SaveNonceAccount(nonce dbtypes.NonceAccountRow) error

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
	SaveStakeAccount(account dbtypes.StakeAccountRow) error

	// DeleteStakeAccount allows to delete the given address of the stake account inside the database
	DeleteStakeAccount(address string) error

	// SaveStakeLockup allows to store the given stake account lockup state inside the database
	SaveStakeLockup(lockup dbtypes.StakeLockupRow) error

	// SaveStakeDelegation allows to store the given delegation of stake account inside the database
	SaveStakeDelegation(delegation dbtypes.StakeDelegationRow) error

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
	// SaveValidator allows to store the given vote account data inside the database
	SaveValidator(account dbtypes.VoteAccountRow) error

	VoteCheckerDb
}

// VoteCheckerDb represents a database that checks account statement of vote properly
type VoteCheckerDb interface {
	// CheckValidatorLatest checks if the vote account statement is latest
	CheckValidatorLatest(address string, currentSlot uint64) bool
}

// VoteDb represents a database that supports vote status properly
type VoteStatusDb interface {
	// SaveValidatorStatuses allows to store the given current validator statuses inside the database
	SaveValidatorStatuses(statuses []dbtypes.ValidatorStatusRow) error

	// GetEpochProducedBlocks allows to get the slots in a epoch inside the database
	// It is for calculating validator skip rates
	GetEpochProducedBlocks(epoch uint64) ([]uint64, error)

	// SaveValidatorSkipRates allows to store the historical validator skip rates of the given epoch inside the database
	SaveValidatorSkipRates(skipRates []dbtypes.ValidatorSkipRateRow) error

	// SaveValidatorSkipRates allows to store the historical validator skip rates of the given epoch inside the database
	SaveHistoryValidatorSkipRates(skipRates []dbtypes.ValidatorSkipRateRow) error
}

// ConfigDb represents a database that supports config properly
type ConfigDb interface {
	// SaveValidatorConfig allows to store the given config account data inside the database
	SaveValidatorConfig(row dbtypes.ValidatorConfigRow) error
}

// BpfLoaderDb represents a database that supports bpf loader properly
type BpfLoaderDb interface {
	// SaveBufferAccount allows to store the given buffer account data inside the database
	SaveBufferAccount(account dbtypes.BufferAccountRow) error

	// DeleteBufferAccount allows to delete the given address of the buffer account inside the database
	DeleteBufferAccount(address string) error

	// SaveProgramAccount allows to store the given program account data inside the database
	SaveProgramAccount(account dbtypes.ProgramAccountRow) error

	// DeleteBufferAccount allows to delete the given address of the program account inside the database
	DeleteProgramAccount(address string) error

	// SaveProgramDataAccount allows to store the given program data account inside the database
	SaveProgramDataAccount(account dbtypes.ProgramDataAccountRow) error

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
	GetTokenUnits() ([]dbtypes.TokenUnitRow, error)

	// SaveTokenUnit allows to save the given token unit details
	SaveTokenUnits(units []dbtypes.TokenUnitRow) error

	// SaveTokensPrices allows to store the token prices inside the database
	SaveTokenPrices(prices []dbtypes.TokenPriceRow) error

	// SaveHistoryTokensPrices allows to store the token prices history inside the database
	SaveHistoryTokenPrices(prices []dbtypes.TokenPriceRow) error
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

type EpochDb interface {
	// SaveSupplyInfo allows to store the current supply info inside the database
	SaveSupplyInfo(dbtypes.SupplyInfoRow) error
}

// FixMissingBlockDb represents a database that supports to get missing blocks info properly
type FixMissingBlockDb interface {

	// GetMissingHeight returns the height of the earliest missing block in the given range
	GetMissingHeight(start uint64, end uint64) (height uint64, err error)

	// GetMissingSlotRange returns the smallest slot range containing the given height
	GetMissingSlotRange(height uint64) (start uint64, end uint64, err error)
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
