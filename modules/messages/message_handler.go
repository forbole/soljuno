package messages

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/types"
)

// HandleMsg represents a message handler that stores the given message inside the proper database table
func HandleMsg(
	index int, msg types.Instruction, tx types.Tx,
	parser MessageParser, cdc bincode.Decoder, db db.Database,
) error {

	// TODO: save msg

	return nil
}
