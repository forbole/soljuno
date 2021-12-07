package db

import (
	"github.com/forbole/soljuno/apis/keybase"
	"github.com/forbole/soljuno/db/migration"
	"github.com/spf13/cobra"
)

func MigrateCmd(cmdCfg *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "migrate",
		Short:   "Migrate the databse to new schema",
		Args:    cobra.ExactArgs(0),
		PreRunE: ReadConfig(cmdCfg),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := GetDatabaseContext(cmdCfg)
			if err != nil {
				return err
			}
			return migration.Up(ctx.Database, keybase.NewClient())
		},
	}
	return cmd
}
