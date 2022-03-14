package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/forbole/soljuno/apis/keybase"
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	accountParser "github.com/forbole/soljuno/solana/account"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/types"
)

// HandleMsg allows to handle different messages types for the config module
func HandleMsg(msg types.Message, tx types.Tx, db db.ConfigDb, client client.ClientProxy) error {
	if len(msg.InvolvedAccounts) != 2 {
		return nil
	}
	address := msg.InvolvedAccounts[0]
	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}
	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}
	configAccount, ok := accountParser.Parse(info.Value.Owner, bz).(accountParser.ValidatorConfig)
	if !ok {
		return fmt.Errorf("failed to parse config account")
	}

	var parsedConfig dbtypes.ParsedValidatorConfig
	err = json.Unmarshal([]byte(configAccount.Info), &parsedConfig)
	if err != nil {
		return err
	}

	kbClient := keybase.NewClient()
	avatarUrl, err := kbClient.GetAvatarURL(parsedConfig.KeybaseUsername)
	if err != nil {
		avatarUrl = ""
	}

	row := dbtypes.NewValidatorConfigRow(
		address, tx.Slot, configAccount.Keys[1].Pubkey.String(), parsedConfig, avatarUrl,
	)

	err = db.SaveValidatorConfig(row)
	if err != nil {
		return err
	}
	return nil
}
