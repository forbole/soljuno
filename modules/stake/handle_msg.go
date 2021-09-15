package stake

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/program/stake"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the stake module
func HandleMsg(msg types.Message, tx types.Tx, db db.StakeDb, client client.Proxy) error {
	switch msg.Value.Type() {
	case "initialize":
		return handleMsgInitialize(msg, tx, db)
	case "authorize":
		return handleMsgAuthorize(msg, db, client)
	case "delegate":
		return handleMsgDelegate(msg, db, client)
	case "split":
		return handleMsgSplit(msg, db, client)
	case "withdraw":
		return handleMsgWithdraw(msg, db, client)
	case "deactivate":
		return handleMsgDeactivate(msg, db, client)
	case "setLockup":
		return handleMsgSetLockup(msg, db, client)
	case "merge":
		return handleMsgMerge(msg, db, client)
	case "authorizeWithSeed":
		return handleMsgAuthorizeWithSeed(msg, db, client)
	case "initializeChecked":
		return handleMsgInitializeChecked(msg, db, client)
	case "authorizeChecked":
		return handleMsgAuthorizeChecked(msg, db, client)
	case "authorizeCheckedWithSeed":
		return handleMsgAuthorizeCheckedWithSeed(msg, db, client)
	case "setLockupChecked":
		return handleMsgSetLockupChecked(msg, db, client)
	}
	return nil
}

// handleMsgInitialize handles a MsgInitialize
func handleMsgInitialize(msg types.Message, tx types.Tx, db db.StakeDb) error {
	instruction, ok := msg.Value.Data().(stake.ParsedInitialize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initialize", msg.Value.Type())

	}
	err := db.SaveStake(instruction.StakeAccount, tx.Slot, instruction.Authorized.Staker, instruction.Authorized.Withdrawer, "initialized")
	if err != nil {
		return err
	}
	return db.SaveStakeLockup(instruction.StakeAccount, tx.Slot, instruction.Lockup.Custodian, instruction.Lockup.Epoch, uint64(instruction.Lockup.UnixTimestamp))
}

// handleMsgAuthorize handles a MsgAuthorize
func handleMsgAuthorize(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedAuthorize)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorize", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgDelegate handles a MsgDelegate
func handleMsgDelegate(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedDelegateStake)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "delegate", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgSplit handles a MsgSplit
func handleMsgSplit(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedSplit)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "split", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgWithdraw handles a MsgWithdraw
func handleMsgWithdraw(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedWithdraw)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "withdraw", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgDeactivate handles a MsgDeactivate
func handleMsgDeactivate(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedDeactivate)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "deactivate", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgSetLockup handles a MsgSetLockup
func handleMsgSetLockup(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedSetLockup)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setLockup", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgMerge handles a MsgMerge
func handleMsgMerge(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedMerge)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "merge", msg.Value.Type())

	}
	err := updateStakeAccount(instruction.Source, db, client)
	if err != nil {
		return err
	}
	return updateStakeAccount(instruction.Destination, db, client)
}

// handleMsgAuthorizeWithSeed handles a MsgAuthorizeWithSeed
func handleMsgAuthorizeWithSeed(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedAuthorizeWithSeed)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeWithSeed", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgInitializeChecked handles a MsgInitializeChecked
func handleMsgInitializeChecked(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedInitializeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "initializeChecked", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgAuthorizeChecked handles a MsgAuthorizeChecked
func handleMsgAuthorizeChecked(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedAuthorizeChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeChecked", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgAuthorizeCheckedWithSeed handles a MsgAuthorizeCheckedWithSeed
func handleMsgAuthorizeCheckedWithSeed(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedAuthorizeCheckedWithSeed)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "authorizeCheckedWithSeed", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}

// handleMsgSetLockupChecked handles a MsgSetLockupChecked
func handleMsgSetLockupChecked(msg types.Message, db db.StakeDb, client client.Proxy) error {
	instruction, ok := msg.Value.Data().(stake.ParsedSetLockupChecked)
	if !ok {
		return fmt.Errorf("instruction does not match %s type: %s", "setLockupChecked", msg.Value.Type())

	}
	return updateStakeAccount(instruction.StakeAccount, db, client)
}
