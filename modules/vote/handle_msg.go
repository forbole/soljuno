package vote

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the vote module
func HandleMsg(msg types.Message, tx types.Tx, db db.VoteDb, client client.Proxy) error {
	switch msg.Value.Type() {
	case "initialize":
	case "authorize":
	case "withdraw":
	case "updateValidatorIdentity":
	case "updateCommission":
	case "authorizeChecked":
	}
	return nil
}

func handleMsgInitialize(msg types.Message, tx types.Tx, db db.VoteDb) error {
	return nil
}

func handleMsgAuthorize() error {
	return nil
}

func handleMsgWithdraw() error {
	return nil
}

func handleUpdateValidatorIdentity() error {
	return nil
}

func handleMsgUpdateCommission() error {
	return nil
}

func handleMsgAuthorizeChecked() error {
	return nil
}
