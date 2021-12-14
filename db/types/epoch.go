package types

type EpochInfoRow struct {
	OneRowID         bool   `db:"one_row_id"`
	Epoch            uint64 `db:"epoch"`
	TransactionCount uint64 `db:"transaction_count"`
}

func NewEpochInfoRow(epoch uint64, count uint64) EpochInfoRow {
	return EpochInfoRow{
		true,
		epoch,
		count,
	}
}

type SupplyInfoRow struct {
	OneRowID       bool   `db:"one_row_id"`
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
		true,
		epoch,
		total,
		circulating,
		nonCirculating,
	}
}

type InflationRateRow struct {
	OneRowID   bool    `db:"one_row_id"`
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
		true,
		epoch,
		total,
		foundation,
		validator,
	}
}
