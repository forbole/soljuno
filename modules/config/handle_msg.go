package config

import (
	"encoding/base64"
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
	"github.com/forbole/soljuno/solana/program/config"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the config module
func HandleMsg(msg types.Message, tx types.Tx, db db.ConfigDb, client client.Proxy) error {
	if len(msg.InvolvedAccounts) != 2 {
		return nil
	}
	address := msg.InvolvedAccounts[0]
	info, err := client.AccountInfo(address)
	if err != nil {
		return err
	}
	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}
	configAccount, ok := accountParser.Parse(config.ProgramID, bz).(accountParser.ConfigAccount)
	if !ok {
		return fmt.Errorf("failed to parse config account")
	}

	return db.SaveConfigAccount(address, tx.Slot, configAccount.Keys[1].Pubkey.String(), configAccount.Info)
}
