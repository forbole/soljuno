package vote

import (
	"fmt"

	"github.com/forbole/soljuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	dbtypes "github.com/forbole/soljuno/db/types"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "vote").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Minute().Do(func() {
		utils.WatchMethod(m, m.updateValidatorsStatus)
	}); err != nil {
		return fmt.Errorf("error while setting up vote periodic operation: %s", err)
	}

	return nil
}

// updateValidatorsStatus insert current validators status
func (m *Module) updateValidatorsStatus() error {
	slot, voteAccounts, err := m.client.GetVoteAccountsWithSlot()
	if err != nil {
		return nil
	}

	statuses := make([]dbtypes.ValidatorStatusRow, len(voteAccounts.Current)+len(voteAccounts.Delinquent))
	count := 0
	for _, account := range voteAccounts.Current {
		statuses[count] = dbtypes.NewValidatorStatusRow(
			account.VotePubkey,
			slot,
			account.ActivatedStake,
			account.LastVote,
			account.RootSlot,
			true,
		)
		count++
	}

	for _, account := range voteAccounts.Delinquent {
		statuses[count] = dbtypes.NewValidatorStatusRow(
			account.VotePubkey,
			slot,
			account.ActivatedStake,
			account.LastVote,
			account.RootSlot,
			false,
		)
		count++
	}
	return m.db.SaveValidatorStatuses(statuses)
}
