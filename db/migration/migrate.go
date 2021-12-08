package migration

import (
	"github.com/forbole/soljuno/db"
)

func Up(db db.Database) error {
	_, err := db.Exec(`DROP INDEX vote_account_node_index`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`ALTER TABLE vote_account RENAME TO validator`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE INDEX validator_node_index ON validator (node)`)
	return err
}

func Down(db db.ExceutorDb) error {
	_, err := db.Exec(`DROP INDEX validator_node_index`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`ALTER TABLE validator RENAME TO vote_account`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE INDEX vote_account_node_index ON vote_account (node)`)

	return err
}
