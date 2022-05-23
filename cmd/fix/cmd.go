package fix

import (
	blockfixer "github.com/forbole/soljuno/cmd/fix/block"
	votefixer "github.com/forbole/soljuno/cmd/fix/vote"
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
		blockfixer.FixMissingBlocksCmd(cmdCfg),
		votefixer.FixVoteAccountsCmd(cmdCfg),
	)
	return cmd
}
