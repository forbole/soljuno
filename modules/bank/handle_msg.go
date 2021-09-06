package bank

import (
	"github.com/forbole/soljuno/types"
	"github.com/rs/zerolog/log"
)

func HandleMsg(tx types.Tx, msg types.Message) error {
	if !tx.Successful() {
		return nil
	}

	log.Info().Str("module", "bank").Str("message", msg.Value.Type()).Uint64("slot", tx.Slot).
		Msg("handled message")
	return nil
}
