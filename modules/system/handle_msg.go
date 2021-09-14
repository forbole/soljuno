package system

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the token module
func HandleMsg(msg types.Message, tx types.Tx, db db.SystemDb, client client.Proxy) error {
	switch msg.Value.Type() {
	case "advanceNonce":
	case "withdrawFromNonce":
	case "initializeNonce":
	case "authorizeNonce":
	}
	return nil
}

func handleMsgAdvaceNonce() error {
	return nil
}

func handleMsgWithdrawFromNonce() error {
	return nil
}

func handleMsgInitializeNonce() error {
	return nil
}

func handleMsgAuthorizeNonce() error {
	return nil
}
