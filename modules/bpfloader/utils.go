package bpfloader

import (
	"encoding/base64"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
)

// updateBufferAccount properly stores the statement of buffer account inside the database
func updateBufferAccount(address string, db db.BpfLoaderDb, client client.Proxy) error {
	info, err := client.AccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.SaveBufferAccount(
			address,
			info.Context.Slot,
			"",
			"closed",
		)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	bufferAccount, ok := accountParser.Parse(info.Value.Owner, bz).(accountParser.BufferAccount)
	if !ok {
		return db.SaveBufferAccount(
			address,
			info.Context.Slot,
			"",
			"closed",
		)
	}

	return db.SaveBufferAccount(
		address,
		info.Context.Slot,
		bufferAccount.Authority.String(),
		"initialized",
	)
}

// updateProgramAccount properly stores the statement of program account inside the database
func updateProgramAccount(address string, db db.BpfLoaderDb, client client.Proxy) error {
	info, err := client.AccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.SaveProgramAccount(
			address,
			info.Context.Slot,
			"",
			"closed",
		)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	programAccount, ok := accountParser.Parse(info.Value.Owner, bz).(accountParser.ProgramAccount)
	if !ok {
		return db.SaveProgramAccount(
			address,
			info.Context.Slot,
			"",
			"closed",
		)
	}
	return db.SaveProgramAccount(
		address,
		info.Context.Slot,
		programAccount.ProgramDataAccount.String(),
		"initialized",
	)
}

// updateProgramDataAccount properly stores the statement of program data account inside the database
func updateProgramDataAccount(address string, db db.BpfLoaderDb, client client.Proxy) error {
	info, err := client.AccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.SaveProgramDataAccount(
			address,
			info.Context.Slot,
			0,
			"",
			"closed",
		)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	programDataAccount, ok := accountParser.Parse(info.Value.Owner, bz).(accountParser.ProgramDataAccount)
	if !ok {
		return db.SaveProgramDataAccount(
			address,
			info.Context.Slot,
			0,
			"",
			"closed",
		)
	}
	return db.SaveProgramDataAccount(
		address,
		info.Context.Slot,
		programDataAccount.Slot,
		programDataAccount.Authority.String(),
		"initialized",
	)
}