package stake

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/types"
)

// HandleInstruction allows to handle different instructions types for the stake module
func HandleInstruction(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	switch instruction.Parsed.Type {
	case "initialize":
		return handleInitialize(instruction, db)
	case "authorize":
		return handleAuthorize(instruction, db, client)
	case "delegate":
		return handleDelegate(instruction, db, client)
	case "split":
		return handleSplit(instruction, db, client)
	case "withdraw":
		return handleWithdraw(instruction, db, client)
	case "deactivate":
		return handleDeactivate(instruction, db, client)
	case "setLockup":
		return handleSetLockup(instruction, db, client)
	case "merge":
		return handleMerge(instruction, db, client)
	case "authorizeWithSeed":
		return handleAuthorizeWithSeed(instruction, db, client)
	case "initializeChecked":
		return handleInitializeChecked(instruction, db, client)
	case "authorizeChecked":
		return handleAuthorizeChecked(instruction, db, client)
	case "authorizeCheckedWithSeed":
		return handleAuthorizeCheckedWithSeed(instruction, db, client)
	case "setLockupChecked":
		return handleSetLockupChecked(instruction, db, client)
	}
	return nil
}

// handleInitialize handles a instruction of Initialize
func handleInitialize(instruction types.Instruction, db db.StakeDb) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedInitialize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initialize", instruction.Parsed.Type)

	}
	err := db.SaveStakeAccount(
		dbtypes.NewStakeAccountRow(
			parsed.StakeAccount, instruction.Slot, parsed.Authorized.Staker, parsed.Authorized.Withdrawer,
		),
	)
	if err != nil {
		return err
	}
	return db.SaveStakeLockup(
		dbtypes.NewStakeLockupRow(
			parsed.StakeAccount, instruction.Slot, parsed.Lockup.Custodian, parsed.Lockup.Epoch, parsed.Lockup.UnixTimestamp,
		),
	)
}

// handleAuthorize handles a instruction of Authorize
func handleAuthorize(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedAuthorize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorize", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleDelegate handles a instruction of Delegate
func handleDelegate(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedDelegateStake)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "delegate", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleSplit handles a instruction of Split
func handleSplit(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedSplit)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "split", instruction.Parsed.Type)

	}
	err := UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
	if err != nil {
		return nil
	}
	return UpdateStakeAccount(parsed.NewSplitAccount, instruction.Slot, db, client)
}

// handleWithdraw handles a instruction of Withdraw
func handleWithdraw(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedWithdraw)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdraw", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleDeactivate handles a instruction of Deactivate
func handleDeactivate(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedDeactivate)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "deactivate", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleSetLockup handles a instruction of SetLockup
func handleSetLockup(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedSetLockup)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setLockup", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleMerge handles a instruction of Merge
func handleMerge(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedMerge)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "merge", instruction.Parsed.Type)

	}
	err := UpdateStakeAccount(parsed.Source, instruction.Slot, db, client)
	if err != nil {
		return err
	}
	return UpdateStakeAccount(parsed.Destination, instruction.Slot, db, client)
}

// handleAuthorizeWithSeed handles a instruction of AuthorizeWithSeed
func handleAuthorizeWithSeed(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedAuthorizeWithSeed)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeWithSeed", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleInitializeChecked handles a instruction of InitializeChecked
func handleInitializeChecked(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedInitializeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeChecked", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleAuthorizeChecked handles a instruction of AuthorizeChecked
func handleAuthorizeChecked(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedAuthorizeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeChecked", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleAuthorizeCheckedWithSeed handles a instruction of AuthorizeCheckedWithSeed
func handleAuthorizeCheckedWithSeed(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedAuthorizeCheckedWithSeed)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeCheckedWithSeed", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}

// handleSetLockupChecked handles a instruction of SetLockupChecked
func handleSetLockupChecked(instruction types.Instruction, db db.StakeDb, client ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedSetLockupChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setLockupChecked", instruction.Parsed.Type)

	}
	return UpdateStakeAccount(parsed.StakeAccount, instruction.Slot, db, client)
}
