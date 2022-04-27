package vote

import (
	"encoding/base64"
	"fmt"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/account/parser"
	solanatypes "github.com/forbole/soljuno/solana/types"
)

// UpdateVoteAccount properly stores the statement of vote account inside the database
func UpdateVoteAccount(address string, currentSlot uint64, db db.VoteDb, client ClientProxy) error {
	if db.CheckValidatorLatest(address, currentSlot) {
		return nil
	}
	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}
	if info.Value == nil {
		return nil
	}
	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}
	voteAccount, ok := parser.Parse(info.Value.Owner, bz).(parser.VoteAccount)
	if !ok {
		return nil
	}
	return db.SaveValidator(
		dbtypes.NewVoteAccountRow(
			address,
			info.Context.Slot,
			voteAccount.Node.String(),
			voteAccount.Withdrawer.String(),
			voteAccount.Voters[0].Pubkey.String(),
			voteAccount.Commission,
		),
	)
}

// UpdateValidatorsStatus insert current validators status
func UpdateValidatorsStatus(db db.VoteDb, client ClientProxy) error {
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
func UpdateValidatorSkipRates(lastEpoch uint64, db db.VoteDb, client ClientProxy) error {
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
