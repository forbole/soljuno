package migration

import (
	"time"

	"github.com/forbole/soljuno/apis/keybase"
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	legacydbtypes "github.com/forbole/soljuno/db/types/legacy"
)

func Up(db db.Database, kbClient keybase.Client) error {
	rows, err := db.Query(`SELECT * FROM validator_config`)
	defer rows.Close()
	if err != nil {
		return err
	}

	configs := []legacydbtypes.ValidatorConfigRow{}
	for rows.Next() {
		var config legacydbtypes.ValidatorConfigRow
		err := rows.Scan(
			&config.Address,
			&config.Slot,
			&config.Owner,
			&config.Name,
			&config.KeybaseUsername,
			&config.Website,
			&config.Details,
		)
		if err != nil {
			return err
		}
		configs = append(configs, config)
	}

	_, err = db.Exec(`ALTER TABLE validator_config ADD COLUMN avatar_url TEXT NOT NULL DEFAULT ''`)
	if err != nil {
		return err
	}

	for _, config := range configs {
		url, err := kbClient.GetAvatarURL(config.KeybaseUsername)
		if err != nil {
			return err
		}
		newConfig := dbtypes.NewValidatorConfigRow(
			config.Address,
			config.Slot,
			config.Owner,
			dbtypes.NewParsedValidatorConfig(
				config.Name,
				config.KeybaseUsername,
				config.Website,
				config.Details,
			),
			url,
		)
		err = db.SaveValidatorConfig(newConfig)
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`ALTER TABLE validator_config DROP COLUMN avatar_url`)
	if err != nil {
		return err
	}
	return nil
}
