package types

type NonceAccountRow struct {
	Address              string `db:"address"`
	Slot                 uint64 `db:"slot"`
	Authority            string `db:"authority"`
	Blockhash            string `db:"blockhash"`
	LamportsPerSignature uint64 `db:"lamports_per_signature"`
}

func NewNonceAccountRow(
	address string, slot uint64, authority string, blockhash string, lamportsPerSignature uint64,
) NonceAccountRow {
	return NonceAccountRow{
		Address:              address,
		Slot:                 slot,
		Authority:            authority,
		Blockhash:            blockhash,
		LamportsPerSignature: lamportsPerSignature,
	}
}
