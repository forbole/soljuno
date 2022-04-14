package bpfloader

import (
	"encoding/base64"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/account/parser"
	"github.com/forbole/soljuno/solana/client"
)

// updateBufferAccount properly stores the statement of buffer account inside the database
func updateBufferAccount(address string, currentSlot uint64, db db.BpfLoaderDb, client client.ClientProxy) error {
	if !db.CheckBufferAccountLatest(address, currentSlot) {
		return nil
	}

	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.DeleteBufferAccount(address)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	bufferAccount, ok := parser.Parse(info.Value.Owner, bz).(parser.BufferAccount)
	if !ok {
		return db.DeleteBufferAccount(address)
	}

	return db.SaveBufferAccount(
		dbtypes.NewBufferAccountRow(
			address,
			info.Context.Slot,
			bufferAccount.Authority.String(),
		),
	)
}

// updateProgramAccount properly stores the statement of program account inside the database
func updateProgramAccount(address string, currentSlot uint64, db db.BpfLoaderDb, client client.ClientProxy) error {
	if !db.CheckProgramAccountLatest(address, currentSlot) {
		return nil
	}

	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.DeleteProgramAccount(address)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	programAccount, ok := parser.Parse(info.Value.Owner, bz).(parser.ProgramAccount)
	if !ok {
		return db.DeleteProgramAccount(address)

	}
	return db.SaveProgramAccount(
		dbtypes.NewProgramAccountRow(
			address,
			info.Context.Slot,
			programAccount.ProgramDataAccount.String(),
		),
	)
}

// updateProgramDataAccount properly stores the statement of program data account inside the database
func updateProgramDataAccount(address string, currentSlot uint64, db db.BpfLoaderDb, client client.ClientProxy) error {
	if !db.CheckProgramDataAccountLatest(address, currentSlot) {
		return nil
	}

	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.DeleteProgramDataAccount(address)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	programDataAccount, ok := parser.Parse(info.Value.Owner, bz).(parser.ProgramDataAccount)
	if !ok {
		return db.DeleteProgramDataAccount(address)
	}
	return db.SaveProgramDataAccount(
		dbtypes.NewProgramDataAccountRow(
			address,
			info.Context.Slot,
			programDataAccount.Slot,
			programDataAccount.Authority.String(),
		),
	)
}
