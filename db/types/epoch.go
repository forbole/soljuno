package types

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
