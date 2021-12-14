package types

type EpochRow struct {
	Epoch uint64 `db:"epoch"`
}

func NewEpochRow(epoch uint64) EpochRow {
	return EpochRow{epoch}
}

type SupplyInfoRow struct {
	Epoch          uint64 `db:"epoch"`
	Total          uint64 `db:"total"`
	Circulating    uint64 `db:"circulating"`
	NonCirculating uint64 `db:"non_circulating"`
}

func NewSupplyInfoRow(
	epoch uint64,
	total uint64,
	circulating uint64,
	nonCirculating uint64,
) SupplyInfoRow {
	return SupplyInfoRow{
		epoch,
		total,
		circulating,
		nonCirculating,
	}
}

type InflationRateRow struct {
	Epoch      uint64  `db:"epoch"`
	Total      float64 `db:"total"`
	Foundation float64 `db:"foundation"`
	Validator  float64 `db:"validator"`
}

func NewInflationRateRow(
	epoch uint64,
	total float64,
	foundation float64,
	validator float64,
) InflationRateRow {
	return InflationRateRow{
		epoch,
		total,
		foundation,
		validator,
	}
}
