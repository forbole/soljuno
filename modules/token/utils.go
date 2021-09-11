package token

import (
	"encoding/base64"
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	accountParser "github.com/forbole/soljuno/solana/account"
	"github.com/forbole/soljuno/solana/program/token"
)

// updateDelegation properly stores the statement of delegation inside the database
func updateDelegation(source string, db db.TokenDb, client client.Proxy) error {
	info, err := client.AccountInfo(source)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.SaveDelegate(source, "", info.Context.Slot, 0)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	tokenAccount, ok := accountParser.Parse(token.ProgramID, bz).(accountParser.TokenAccount)
	if !ok {
		return db.SaveDelegate(source, "", info.Context.Slot, 0)
	}

	if !tokenAccount.Delegate.Option.Bool() {
		return db.SaveDelegate(source, "", info.Context.Slot, 0)
	}

	return db.SaveDelegate(source, tokenAccount.Delegate.Value.String(), info.Context.Slot, tokenAccount.DelegateAmount)
}

// updateTokenSupply properly stores the supply of the given mint inside the database
func updateTokenSupply(mint string, db db.TokenDb, client client.Proxy) error {
	info, err := client.AccountInfo(mint)
	if err != nil {
		return err
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	token, ok := accountParser.Parse(token.ProgramID, bz).(accountParser.TokenMint)
	if !ok {
		return fmt.Errorf("failed to parse token:%s", mint)
	}
	return db.SaveTokenSupply(mint, info.Context.Slot, token.Supply)
}

// updateAccountState properly stores the account state inside database
func updateAccountState(address string, db db.TokenDb, client client.Proxy) error {
	info, err := client.AccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.SaveTokenAccount(address, info.Context.Slot, "", "", "closed")
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	tokenAccount, ok := accountParser.Parse(token.ProgramID, bz).(accountParser.TokenAccount)
	if !ok {
		return db.SaveTokenAccount(address, info.Context.Slot, "", "", "closed")
	}
	return db.SaveTokenAccount(address, info.Context.Slot, tokenAccount.Mint.String(), tokenAccount.Owner.String(), tokenAccount.State.String())
}
