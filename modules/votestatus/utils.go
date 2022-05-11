package votestatus

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	solanatypes "github.com/forbole/soljuno/solana/types"
)

// UpdateValidatorsStatus insert current validators status
func UpdateValidatorsStatus(db db.VoteStatusDb, client ClientProxy) error {
	slot, voteAccounts, err := client.GetVoteAccountsWithSlot()
	if err != nil {
		return err
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
	return db.SaveValidatorStatuses(statuses)
}

// UpdateValidatorSkipRates properly stores the skip rates of all validators inside the database
func UpdateValidatorSkipRates(lastEpoch uint64, db db.VoteStatusDb, client ClientProxy) error {
	slots, err := db.GetEpochProducedBlocks(lastEpoch)
	if err != nil {
		return err
	}
	if len(slots) == 0 {
		return fmt.Errorf("%d epoch blocks does not exist", lastEpoch)
	}

	endSlot := slots[len(slots)-1]
	schedules, err := client.GetLeaderSchedule(lastEpoch * solanatypes.SlotsInEpoch)
	if err != nil {
		return err
	}

	produced := make(map[int]bool)
	for _, slot := range slots {
		produced[int(slot%solanatypes.SlotsInEpoch)] = true
	}

	skipRateRows := make([]dbtypes.ValidatorSkipRateRow, len(schedules))
	count := 0
	end := int(endSlot % solanatypes.SlotsInEpoch)
	for validator, schedule := range schedules {
		total, skip := GetSkipRateReference(end, produced, schedule)
		skipRate := float64(skip) / float64(total)
		skipRateRows[count] = dbtypes.NewValidatorSkipRateRow(validator, lastEpoch, skipRate, total, skip)
		count++
	}

	err = db.SaveValidatorSkipRates(skipRateRows)
	if err != nil {
		return err
	}
	return db.SaveHistoryValidatorSkipRates(skipRateRows)
}

// GetSkipRateReference returns the total and skip amount in a epoch of the validator from the given produced map and the validator schedule
func GetSkipRateReference(end int, produced map[int]bool, schedule []int) (int, int) {
	var skip int = 0
	var total int = 0
	for _, slotInEpoch := range schedule {
		total++
		if slotInEpoch > end {
			break
		}
		if ok := produced[slotInEpoch]; !ok {
			skip++
		}
	}
	return total, skip
}
