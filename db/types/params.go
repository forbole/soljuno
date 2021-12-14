package types

type EpochScheduleParamRow struct {
	OneRowID         bool   `db:"one_row_id"`
	Epoch            uint64 `db:"epoch"`
	SlotsPerEpoch    uint64 `db:"slots_per_epoch"`
	FirstNormalEpoch uint64 `db:"first_normal_epoch"`
	FirstNormalSlot  uint64 `db:"first_normal_slot"`
	Warmup           bool   `db:"warmup"`
}

func NewEpochScheduleParamRow(
	epoch uint64,
	slotsPerEpoch uint64,
	firstNormalEpoch uint64,
	firstNormalSlot uint64,
	warmup bool,
) EpochScheduleParamRow {
	return EpochScheduleParamRow{
		true,
		epoch,
		slotsPerEpoch,
		firstNormalEpoch,
		firstNormalSlot,
		warmup,
	}
}

type InflationGovernorParamRow struct {
	OneRowID           bool    `db:"one_row_id"`
	Epoch              uint64  `db:"epoch"`
	Initial            float64 `db:"initial"`
	Terminal           float64 `db:"terminal"`
	Taper              float64 `db:"taper"`
	Foundation         float64 `db:"foundation"`
	FoundationTerminal float64 `db:"foundation_terminal"`
}

func NewInflationGovernorParamRow(
	epoch uint64,
	initial float64,
	terminal float64,
	taper float64,
	foundation float64,
	foundationTerminal float64,
) InflationGovernorParamRow {
	return InflationGovernorParamRow{
		true,
		epoch,
		initial,
		terminal,
		taper,
		foundation,
		foundationTerminal,
	}
}
