package bpfloader

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the bpfloader module
func HandleMsg(msg types.Message, tx types.Tx, db db.BpfLoaderDb) error {
	switch msg.Value.Type() {
	case "initializeBuffer":
		return handleMsgInitializeBuffer(msg, tx, db)
	case "deployWithMaxDataLen":
		return handleMsgDeployWithMaxDataLen(msg, tx, db)
	case "upgrade":
		return handleMsgUpgrade(msg, tx, db)
	case "setAuthority":
		return handleMsgSetAuthority(msg, tx, db)
	case "close":
		return handleMsgClose(msg, tx, db)
	}
	return nil
}

func handleMsgInitializeBuffer(msg types.Message, tx types.Tx, db db.BpfLoaderDb) error {
	return nil
}

func handleMsgDeployWithMaxDataLen(msg types.Message, tx types.Tx, db db.BpfLoaderDb) error {
	return nil
}

func handleMsgUpgrade(msg types.Message, tx types.Tx, db db.BpfLoaderDb) error {
	return nil
}

func handleMsgSetAuthority(msg types.Message, tx types.Tx, db db.BpfLoaderDb) error {
	return nil
}

func handleMsgClose(msg types.Message, tx types.Tx, db db.BpfLoaderDb) error {
	return nil
}
