package token

import (
	"encoding/base64"
	"fmt"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/account/parser"
	"github.com/forbole/soljuno/solana/client"
)

// updateToken properly stores the authority of mint inside the database
func updateToken(mint string, currentSlot uint64, db db.TokenDb, client client.ClientProxy) error {
	if db.CheckTokenLatest(mint, currentSlot) {
		return nil
	}

	info, err := client.GetAccountInfo(mint)
	if err != nil {
		return err
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	token, ok := parser.Parse(info.Value.Owner, bz).(parser.Token)
	if !ok {
		return fmt.Errorf("failed to parse token:%s", mint)
	}

	return db.SaveToken(
		dbtypes.NewTokenRow(
			mint,
			info.Context.Slot,
			token.Decimals,
			token.MintAuthority.String(),
			token.FreezeAuthority.String(),
		),
	)
}

// updateTokenAccount properly stores the account state inside database
func updateTokenAccount(address string, currentSlot uint64, db db.TokenDb, client client.ClientProxy) error {
	if db.CheckTokenAccountLatest(address, currentSlot) {
		return nil
	}

	info, err := client.GetAccountInfo(address)
	if err != nil {
		return err
	}

	if info.Value == nil {
		return db.DeleteTokenAccount(address)
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	tokenAccount, ok := parser.Parse(info.Value.Owner, bz).(parser.TokenAccount)
	if !ok {
		return db.DeleteTokenAccount(address)
	}

	err = db.SaveTokenAccount(
		dbtypes.NewTokenAccountRow(
			address, info.Context.Slot, tokenAccount.Mint.String(), tokenAccount.Owner.String()),
	)
	if err != nil {
		return err
	}

	if tokenAccount.Delegate.String() == "" {
		return db.DeleteTokenDelegation(address)
	}

	return db.SaveTokenDelegation(
		dbtypes.NewTokenDelegationRow(
			address, tokenAccount.Delegate.String(), info.Context.Slot, tokenAccount.DelegateAmount,
		),
	)
}

// updateTokenSupply properly stores the supply of the given mint inside the database
func updateTokenSupply(mint string, currentSlot uint64, db db.TokenDb, client client.ClientProxy) error {
	if db.CheckTokenSupplyLatest(mint, currentSlot) {
		return nil
	}

	info, err := client.GetAccountInfo(mint)
	if err != nil {
		return err
	}

	bz, err := base64.StdEncoding.DecodeString(info.Value.Data[0])
	if err != nil {
		return err
	}

	token, ok := parser.Parse(info.Value.Owner, bz).(parser.Token)
	if !ok {
		return fmt.Errorf("failed to parse token:%s", mint)
	}
	return db.SaveTokenSupply(dbtypes.NewTokenSupplyRow(mint, info.Context.Slot, token.Supply))
}
