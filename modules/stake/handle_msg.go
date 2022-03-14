package stake

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the stake module
func HandleMsg(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	switch msg.Parsed.Type {
	case "initialize":
		return handleMsgInitialize(msg, tx, db)
	case "authorize":
		return handleMsgAuthorize(msg, tx, db, client)
	case "delegate":
		return handleMsgDelegate(msg, tx, db, client)
	case "split":
		return handleMsgSplit(msg, tx, db, client)
	case "withdraw":
		return handleMsgWithdraw(msg, tx, db, client)
	case "deactivate":
		return handleMsgDeactivate(msg, tx, db, client)
	case "setLockup":
		return handleMsgSetLockup(msg, tx, db, client)
	case "merge":
		return handleMsgMerge(msg, tx, db, client)
	case "authorizeWithSeed":
		return handleMsgAuthorizeWithSeed(msg, tx, db, client)
	case "initializeChecked":
		return handleMsgInitializeChecked(msg, tx, db, client)
	case "authorizeChecked":
		return handleMsgAuthorizeChecked(msg, tx, db, client)
	case "authorizeCheckedWithSeed":
		return handleMsgAuthorizeCheckedWithSeed(msg, tx, db, client)
	case "setLockupChecked":
		return handleMsgSetLockupChecked(msg, tx, db, client)
	}
	return nil
}

// handleMsgInitialize handles a MsgInitialize
func handleMsgInitialize(msg types.Message, tx types.Tx, db db.StakeDb) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedInitialize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initialize", msg.Parsed.Type)

	}
	err := db.SaveStakeAccount(instruction.StakeAccount, tx.Slot, instruction.Authorized.Staker, instruction.Authorized.Withdrawer)
	if err != nil {
		return err
	}
	return db.SaveStakeLockup(instruction.StakeAccount, tx.Slot, instruction.Lockup.Custodian, instruction.Lockup.Epoch, instruction.Lockup.UnixTimestamp)
}

// handleMsgAuthorize handles a MsgAuthorize
func handleMsgAuthorize(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedAuthorize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorize", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgDelegate handles a MsgDelegate
func handleMsgDelegate(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedDelegateStake)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "delegate", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgSplit handles a MsgSplit
func handleMsgSplit(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedSplit)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "split", msg.Parsed.Type)

	}
	err := updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
	if err != nil {
		return nil
	}
	return updateStakeAccount(instruction.NewSplitAccount, tx.Slot, db, client)
}

// handleMsgWithdraw handles a MsgWithdraw
func handleMsgWithdraw(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedWithdraw)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdraw", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgDeactivate handles a MsgDeactivate
func handleMsgDeactivate(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedDeactivate)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "deactivate", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgSetLockup handles a MsgSetLockup
func handleMsgSetLockup(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedSetLockup)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setLockup", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgMerge handles a MsgMerge
func handleMsgMerge(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedMerge)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "merge", msg.Parsed.Type)

	}
	err := updateStakeAccount(instruction.Source, tx.Slot, db, client)
	if err != nil {
		return err
	}
	return updateStakeAccount(instruction.Destination, tx.Slot, db, client)
}

// handleMsgAuthorizeWithSeed handles a MsgAuthorizeWithSeed
func handleMsgAuthorizeWithSeed(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedAuthorizeWithSeed)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeWithSeed", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgInitializeChecked handles a MsgInitializeChecked
func handleMsgInitializeChecked(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedInitializeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeChecked", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeChecked handles a MsgAuthorizeChecked
func handleMsgAuthorizeChecked(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedAuthorizeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeChecked", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgAuthorizeCheckedWithSeed handles a MsgAuthorizeCheckedWithSeed
func handleMsgAuthorizeCheckedWithSeed(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedAuthorizeCheckedWithSeed)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeCheckedWithSeed", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}

// handleMsgSetLockupChecked handles a MsgSetLockupChecked
func handleMsgSetLockupChecked(msg types.Message, tx types.Tx, db db.StakeDb, client client.ClientProxy) error {
	instruction, ok := msg.Parsed.Value.(stake.ParsedSetLockupChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setLockupChecked", msg.Parsed.Type)

	}
	return updateStakeAccount(instruction.StakeAccount, tx.Slot, db, client)
}
