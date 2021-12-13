package types

type EpochRow struct {
	Epoch uint64 `db:"epoch"`
}

func NewEpochRow(epoch uint64) EpochRow {
	return EpochRow{epoch}
}

type EpochSupplyRow struct {
	Epoch          uint64 `db:"epoch"`
	Total          uint64 `db:"total"`
	Circulating    uint64 `db:"circulating"`
	NonCirculating uint64 `db:"non_circulating"`
}

func NewEpochSupplyRow(
	epoch uint64,
	total uint64,
	circulating uint64,
	nonCirculating uint64,
) EpochSupplyRow {
	return EpochSupplyRow{
		epoch,
		total,
		circulating,
		nonCirculating,
	}
}

type EpochInflationRow struct {
	Epoch     uint64  `db:"epoch"`
	Total     float64 `db:"total"`
	Fondation float64 `db:"fondation"`
	Validator float64 `db:"validator"`
}

func NewEpochInflationRow(
	epoch uint64,
	total float64,
	foundation float64,
	validator float64,
) EpochInflationRow {
	return EpochInflationRow{
		epoch,
		total,
		foundation,
		validator,
	}
}
