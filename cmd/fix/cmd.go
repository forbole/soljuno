package fix

import (
	cmdtypes "github.com/forbole/soljuno/cmd/types"
	"github.com/spf13/cobra"
)

func FixCmd(cmdCfg *cmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "fix",
		Short:                      "Fix subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	cmd.AddCommand(
		FixMissingBlocksCmd(cmdCfg),
	)
	return cmd
}
