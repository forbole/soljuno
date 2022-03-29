package types

import "github.com/lib/pq"

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

//____________________________________________________________________________

type TokenAccountRow struct {
	Address string `db:"address"`
	Slot    uint64 `db:"slot"`
	Mint    string `db:"mint"`
	Owner   string `db:"owner"`
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

//____________________________________________________________________________

type MultisigRow struct {
	Address string         `db:"address"`
	Slot    uint64         `db:"slot"`
	Signers pq.StringArray `db:"signers"`
	Minimum uint8          `db:"minimum"`
}

func NewMultisigRow(address string, slot uint64, signers []string, m uint8) MultisigRow {
	return MultisigRow{
		Address: address,
		Slot:    slot,
		Signers: *pq.Array(signers).(*pq.StringArray),
		Minimum: m,
	}
}

//____________________________________________________________________________

type TokenDelegationRow struct {
	Source      string `db:"source_address"`
	Destination string `db:"delegate_address"`
	Slot        uint64 `db:"slot"`
	Amount      uint64 `db:"amount"`
}

func NewTokenDelegationRow(source string, destination string, slot uint64, amount uint64) TokenDelegationRow {
	return TokenDelegationRow{
		Source:      source,
		Destination: destination,
		Slot:        slot,
		Amount:      amount,
	}
}

//____________________________________________________________________________

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
