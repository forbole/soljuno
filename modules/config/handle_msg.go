package config

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/types"
)

func HandleMsg(msg types.Message, tx types.Tx, client client.Proxy) error {
	configAccount := ""
	if len(msg.InvolvedAccounts) == 1 {
		configAccount = msg.InvolvedAccounts[0]
	}
	if len(msg.InvolvedAccounts) == 2 {
		configAccount = msg.InvolvedAccounts[1]
	}
	_, err := client.AccountInfo(configAccount)
	if err != nil {
		return err
	}

	return nil
}
