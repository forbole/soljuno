package types

type BufferAccountRow struct {
	Address   string `db:"address"`
	Slot      uint64 `db:"slot"`
	Authority string `db:"authority"`
}

func NewBufferAccountRow(address string, slot uint64, authority string) BufferAccountRow {
	return BufferAccountRow{
		Address:   address,
		Slot:      slot,
		Authority: authority,
	}
}

//____________________________________________________________________________

type ProgramAccountRow struct {
	Address            string `db:"address"`
	Slot               uint64 `db:"slot"`
	ProgramDataAccount string `db:"program_data_account"`
}

func NewProgramAccountRow(address string, slot uint64, programDataAccount string) ProgramAccountRow {
	return ProgramAccountRow{
		Address:            address,
		Slot:               slot,
		ProgramDataAccount: programDataAccount,
	}
}

//____________________________________________________________________________

type ProgramDataAccountRow struct {
	Address          string `db:"address"`
	Slot             uint64 `db:"slot"`
	LastModifiedSlot uint64 `db:"last_modified_slot"`
	UpdateAuthority  string `db:"update_authority"`
}

func NewProgramDataAccountRow(address string, slot uint64, lastModifiedSlot uint64, updateAuthority string) ProgramDataAccountRow {
	return ProgramDataAccountRow{
		Address:          address,
		Slot:             slot,
		LastModifiedSlot: lastModifiedSlot,
		UpdateAuthority:  updateAuthority,
	}
}
