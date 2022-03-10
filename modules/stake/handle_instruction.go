package stake

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/types"
)

// HandleInstruction allows to handle different instructions types for the stake module
func HandleInstruction(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	switch instruction.Parsed.Type {
	case "initialize":
		return handleInitialize(instruction, tx, db)
	case "authorize":
		return handleAuthorize(instruction, tx, db, client)
	case "delegate":
		return handleDelegate(instruction, tx, db, client)
	case "split":
		return handleSplit(instruction, tx, db, client)
	case "withdraw":
		return handleWithdraw(instruction, tx, db, client)
	case "deactivate":
		return handleDeactivate(instruction, tx, db, client)
	case "setLockup":
		return handleSetLockup(instruction, tx, db, client)
	case "merge":
		return handleMerge(instruction, tx, db, client)
	case "authorizeWithSeed":
		return handleAuthorizeWithSeed(instruction, tx, db, client)
	case "initializeChecked":
		return handleInitializeChecked(instruction, tx, db, client)
	case "authorizeChecked":
		return handleAuthorizeChecked(instruction, tx, db, client)
	case "authorizeCheckedWithSeed":
		return handleAuthorizeCheckedWithSeed(instruction, tx, db, client)
	case "setLockupChecked":
		return handleSetLockupChecked(instruction, tx, db, client)
	}
	return nil
}

// handleInitialize handles a MsgInitialize
func handleInitialize(instruction types.Instruction, tx types.Tx, db db.StakeDb) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedInitialize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initialize", instruction.Parsed.Type)

	}
	err := db.SaveStakeAccount(parsed.StakeAccount, tx.Slot, parsed.Authorized.Staker, parsed.Authorized.Withdrawer)
	if err != nil {
		return err
	}
	return db.SaveStakeLockup(parsed.StakeAccount, tx.Slot, parsed.Lockup.Custodian, parsed.Lockup.Epoch, parsed.Lockup.UnixTimestamp)
}

// handleAuthorize handles a MsgAuthorize
func handleAuthorize(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedAuthorize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorize", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleDelegate handles a MsgDelegate
func handleDelegate(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedDelegateStake)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "delegate", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleSplit handles a MsgSplit
func handleSplit(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedSplit)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "split", instruction.Parsed.Type)

	}
	err := updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
	if err != nil {
		return nil
	}
	return updateStakeAccount(parsed.NewSplitAccount, tx.Slot, db, client)
}

// handleWithdraw handles a MsgWithdraw
func handleWithdraw(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedWithdraw)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdraw", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleDeactivate handles a MsgDeactivate
func handleDeactivate(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedDeactivate)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "deactivate", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleSetLockup handles a MsgSetLockup
func handleSetLockup(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedSetLockup)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setLockup", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleMerge handles a MsgMerge
func handleMerge(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedMerge)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "merge", instruction.Parsed.Type)

	}
	err := updateStakeAccount(parsed.Source, tx.Slot, db, client)
	if err != nil {
		return err
	}
	return updateStakeAccount(parsed.Destination, tx.Slot, db, client)
}

// handleAuthorizeWithSeed handles a MsgAuthorizeWithSeed
func handleAuthorizeWithSeed(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedAuthorizeWithSeed)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeWithSeed", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleInitializeChecked handles a MsgInitializeChecked
func handleInitializeChecked(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedInitializeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeChecked", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleAuthorizeChecked handles a MsgAuthorizeChecked
func handleAuthorizeChecked(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedAuthorizeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeChecked", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleAuthorizeCheckedWithSeed handles a MsgAuthorizeCheckedWithSeed
func handleAuthorizeCheckedWithSeed(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedAuthorizeCheckedWithSeed)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeCheckedWithSeed", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}

// handleSetLockupChecked handles a MsgSetLockupChecked
func handleSetLockupChecked(instruction types.Instruction, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	parsed, ok := instruction.Parsed.Value.(stake.ParsedSetLockupChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setLockupChecked", instruction.Parsed.Type)

	}
	return updateStakeAccount(parsed.StakeAccount, tx.Slot, db, client)
}
