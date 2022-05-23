package vote

import (
	cmdtypes "github.com/forbole/soljuno/cmd/types"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules/vote"
	"github.com/spf13/cobra"
)

func FixVoteAccountsCmd(cmdCfg *cmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "vote-accounts",
		Short:   "Fix vote accounts",
		PreRunE: cmdtypes.ReadConfig(cmdCfg),
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := GetContext(cmdCfg)
			if err != nil {
				return err
			}
			return fixVoteAccounts(ctx)
		},
	}
	return cmd
}

func fixVoteAccounts(ctx *Context) error {
	rows, err := ctx.Database.Query("SELECT * FROM validator")
	if err != nil {
		return err
	}
	voteAccounts := []dbtypes.VoteAccountRow{}
	defer rows.Close()
	for rows.Next() {
		var voteAccount dbtypes.VoteAccountRow
		if err := rows.Scan(
			&voteAccount.Address,
			&voteAccount.Slot,
			&voteAccount.Node,
			&voteAccount.Voter,
			&voteAccount.Withdrawer,
			&voteAccount.Commission,
		); err != nil {
			return err
		}
		voteAccounts = append(voteAccounts, voteAccount)
	}
	slot, err := ctx.Proxy.GetLatestSlot()
	if err != nil {
		return err
	}

	for _, voteAccount := range voteAccounts {
		err := vote.UpdateVoteAccount(voteAccount.Address, slot, ctx.Database, ctx.Proxy)
		if err != nil {
			return err
		}
	}
	return nil
}
