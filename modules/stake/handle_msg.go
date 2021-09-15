package stake

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the stake module
func HandleMsg(msg types.Message, tx types.Tx, db db.StakeDb, client client.Proxy) error {
	switch msg.Value.Type() {
	case "initialize":
	case "authorize":
	case "delegate":
	case "split":
	case "withdraw":
	case "deactivate":
	case "setLockup":
	case "merge":
	case "authorizeWithSeed":
	case "initializeChecked":
	case "authorizeChecked":
	case "authorizeCheckedWithSeed":
	case "setLockupChecked":
	}
	return nil
}

// handleMsgInitialize handles a MsgInitialize
func handleMsgInitialize(msg types.Message, tx types.Tx, db db.StakeDb) error {
	return nil
}

// handleMsgAuthorize handles a MsgAuthorize
func handleMsgAuthorize() error {
	return nil
}

// handleMsgDelegate handles a MsgDelegate
func handleMsgDelegate() error {
	return nil
}

// handleMsgSplit handles a MsgSplit
func handleMsgSplit() error {
	return nil
}

// handleMsgWithdraw handles a MsgWithdraw
func handleMsgWithdraw() error {
	return nil
}

// handleMsgDeactivate handles a MsgDeactivate
func handleMsgDeactivate() error {
	return nil
}

// handleMsgSetLockup handles a MsgSetLockup
func handleMsgSetLockup() error {
	return nil
}

// handleMsgMerge handles a MsgMerge
func handleMsgMerge() error {
	return nil
}

// handleMsgAuthorizeWithSeed handles a MsgAuthorizeWithSeed
func handleMsgAuthorizeWithSeed() error {
	return nil
}

// handleMsgInitializeChecked handles a MsgInitializeChecked
func handleMsgInitializeChecked() error {
	return nil
}

// handleMsgAuthorizeChecked handles a MsgAuthorizeChecked
func handleMsgAuthorizeChecked() error {
	return nil
}

// handleMsgAuthorizeCheckedWithSeed handles a MsgAuthorizeCheckedWithSeed
func handleMsgAuthorizeCheckedWithSeed() error {
	return nil
}

// handleMsgSetLockupChecked handles a MsgSetLockupChecked
func handleMsgSetLockupChecked() error {
	return nil
}
