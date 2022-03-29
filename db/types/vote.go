package types

type VoteAccountRow struct {
	Address    string `db:"address"`
	Slot       uint64 `db:"slot"`
	Node       string `db:"node"`
	Voter      string `db:"voter"`
	Withdrawer string `db:"withdrawer"`
	Commission uint8  `db:"commission"`
}

func NewVoteAccountRow(address string, slot uint64, node string, voter string, withdrawer string, commission uint8) VoteAccountRow {
	return VoteAccountRow{
		Address:    address,
		Slot:       slot,
		Node:       node,
		Voter:      voter,
		Withdrawer: withdrawer,
		Commission: commission,
	}
}

//____________________________________________________________________________

type ValidatorSkipRateRow struct {
	Address  string  `db:"address"`
	Epoch    uint64  `db:"epoch"`
	SkipRate float64 `db:"skip_rate"`
	Total    int     `db:"total"`
	Skip     int     `db:"skip"`
}

func NewValidatorSkipRateRow(address string, epoch uint64, skipRate float64, total int, skip int) ValidatorSkipRateRow {
	return ValidatorSkipRateRow{
		Address:  address,
		Epoch:    epoch,
		SkipRate: skipRate,
		Total:    total,
		Skip:     skip,
	}
}

//____________________________________________________________________________

type ValidatorStatusRow struct {
	Address        string `db:"address"`
	Slot           uint64 `db:"slot"`
	ActivatedStake uint64 `db:"activated_stake"`
	LastVote       uint64 `db:"last_vote"`
	RootSlot       uint64 `db:"root_slot"`
	Active         bool   `db:"active"`
}

func NewValidatorStatusRow(
	address string,
	slot uint64,
	activatedStake uint64,
	lastVote uint64,
	rootSlot uint64,
	active bool,
) ValidatorStatusRow {
	return ValidatorStatusRow{
		Address:        address,
		Slot:           slot,
		ActivatedStake: activatedStake,
		LastVote:       lastVote,
		RootSlot:       rootSlot,
		Active:         active,
	}
}
