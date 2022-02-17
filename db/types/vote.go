package types

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
