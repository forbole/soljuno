package types

type TokenRow struct {
	Mint            string `db:"mint"`
	Slot            uint64 `db:"slot"`
	Decimals        uint8  `db:"decimals"`
	MintAuthority   string `db:"mint_authority"`
	FreezeAuthority string `db:"freeze_authority"`
}

func NewTokenRow(
	mint string,
	slot uint64,
	decimals uint8,
	mintAuthority string,
	freezeAuthority string,
) TokenRow {
	return TokenRow{
		Mint:            mint,
		Slot:            slot,
		Decimals:        decimals,
		MintAuthority:   mintAuthority,
		FreezeAuthority: freezeAuthority,
	}
}

type TokenAccountRow struct {
	Address string `db:"address"`
	Slot    uint64 `db:"slot"`
	Mint    string `json:"mint"`
	Owner   string `json:"owner"`
}

func NewTokenAccountRow(
	address string, slot uint64, mint string, owner string,
) TokenAccountRow {
	return TokenAccountRow{
		Address: address,
		Slot:    slot,
		Mint:    mint,
		Owner:   owner,
	}
}

type TokenSupplyRow struct {
	Mint   string `db:"mint"`
	Slot   uint64 `db:"slot"`
	Supply uint64 `db:"supply"`
}

func NewTokenSupplyRow(mint string, slot uint64, supply uint64) TokenSupplyRow {
	return TokenSupplyRow{
		Mint:   mint,
		Slot:   slot,
		Supply: supply,
	}
}
