package db

import (
	cmdtypes "github.com/forbole/soljuno/cmd/types"
	"github.com/spf13/cobra"
)

func DbCmd(cmdCfg *cmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "db",
		Short:                      "Database subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	cmd.AddCommand(
		InitCmd(cmdCfg),
		MigrateCmd(cmdCfg),
	)
	return cmd
}
