package types

type ValidatorSkipRateRow struct {
	Address  string  `db:"address"`
	Epoch    uint64  `db:"epoch"`
	SkipRate float64 `db:"skip_rate"`
}

func NewValidatorSkipRateRow(address string, epoch uint64, skipRate float64) ValidatorSkipRateRow {
	return ValidatorSkipRateRow{
		Address:  address,
		Epoch:    epoch,
		SkipRate: skipRate,
	}
}
