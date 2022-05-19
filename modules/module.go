package modules

import (
	"github.com/go-co-op/gocron"

	"github.com/forbole/soljuno/types"
)

// Module represents a generic module without any particular handling of data
type Module interface {
	// Name returns the module name
	Name() string
}

// Modules represents a slice of Module objects
type Modules []Module

// FindByName returns the module having the given name inside the m slice.
// If no modules are found, returns nil and false.
func (m Modules) FindByName(name string) (module Module, found bool) {
	for _, m := range m {
		if m.Name() == name {
			return m, true
		}
	}
	return nil, false
}

// --------------------------------------------------------------------------------------------------------------------

type AdditionalOperationsModule interface {
	// RunAdditionalOperations runs all the additional operations required by the module.
	// This is the perfect place where to initialize all the operations that subscribe to websockets or other
	// external sources.
	// NOTE. This method will only be run ONCE before starting the parsing of the blocks.
	RunAdditionalOperations() error
}

type AsyncOperationsModule interface {
	// RunAsyncOperations runs all the async operations associated with a module.
	// This method will be run on a separate goroutine, that will stop only when the user stops the entire process.
	// For this reason, this method cannot return an error, and all raised errors should be signaled by panicking.
	RunAsyncOperations()
}

type PeriodicOperationsModule interface {
	// RegisterPeriodicOperations allows to register all the operations that will be run on a periodic basis.
	// The given scheduler can be used to define the periodicity of each task.
	// NOTE. This method will only be run ONCE during the module initialization.
	RegisterPeriodicOperations(scheduler *gocron.Scheduler) error

	RunPeriodicOperations() error
}

type BlockModule interface {
	// HandleBlock allows to handle a single block.
	// For convenience of use, all the transactions present inside the given block
	// and the currently used database will be passed as well.
	// For each transaction present inside the block, HandleTx will be called as well.
	// NOTE. The returned error will be logged using the logging.LogBlockError method. All other modules' handlers
	// will still be called.
	HandleBlock(block types.Block) error
}

type TransactionModule interface {
	// HandleTx handles a single transaction.
	// For each instruction present inside the transaction, HandleInstruction will be called as well.
	// NOTE. The returned error will be logged using the logging.LogTxError method. All other modules' handlers
	// will still be called.
	HandleTx(tx types.Tx) error
}

type InstructionModule interface {
	// HandleInstruction handles a single instruction.
	// For convenience of use, the index of the instruction inside the transaction and the transaction itself
	// are passed as well.
	// NOTE. The returned error will be logged using the logging.LogInstructionError method. All other modules' handlers
	// will still be called.
	HandleInstruction(instruction types.Instruction, tx types.Tx) error
}
