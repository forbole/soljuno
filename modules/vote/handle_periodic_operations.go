package vote

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules/utils"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "vote").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Minute().Do(func() {
		utils.WatchMethod(m.updateValidatorsStatus)
	}); err != nil {
		return fmt.Errorf("error while setting up vote periodic operation: %s", err)
	}

	return nil
}

// updateValidatorsStatus insert current validators status
func (m *Module) updateValidatorsStatus() error {
	voteDb, ok := m.db.(db.VoteDb)
	if !ok {
		return fmt.Errorf("vote is enabled, but your database does not implement VoteDb")
	}

	slot, voteAccounts, err := m.client.ValidatorsWithSlot()
	if err != nil {
		return nil
	}

	for _, account := range voteAccounts.Current {
		if err := voteDb.SaveValidatorStatus(
			account.VotePubkey,
			slot,
			account.ActivatedStake,
			account.LastVote,
			account.RootSlot,
			true,
		); err != nil {
			return err
		}
	}

	for _, account := range voteAccounts.Delinquent {
		if err := voteDb.SaveValidatorStatus(
			account.VotePubkey,
			slot,
			account.ActivatedStake,
			account.LastVote,
			account.RootSlot,
			false,
		); err != nil {
			return err
		}
	}
	return nil
}
