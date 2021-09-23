package bpfloader

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the bpfloader module
func HandleMsg(msg types.Message, tx types.Tx, db db.ConfigDb) error {
	return nil
}
