package vote

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the vote module
func HandleMsg(msg types.Message, tx types.Tx, db db.Database, client client.Proxy) error {
	switch msg.Value.Type() {
	}
	return nil
}
