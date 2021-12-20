package snapshot

import (
	"time"
)

func consumeBuffer(ctx *Context, parallelize int) {
	for {
		accounts := getAccounts(ctx, parallelize)
		addresses, balances, err := getBalances(accounts)
		if err != nil {
			ctx.Logger.Error("failed to get balances", "addresses", addresses, "err", err)
		}
		go func() {
			err := ctx.Database.SaveAccountBalances(0, addresses, balances)
			if err != nil {
				ctx.Logger.Error("failed to save balances", "addresses", addresses, "err", err)
			}
		}()
	}
}

func getAccounts(ctx *Context, num int) []Account {
	var accounts []Account
	for {
		select {
		case account := <-ctx.Buffer:
			accounts = append(accounts, account)
			if len(accounts) >= num {
				return accounts
			}
		case <-time.After(100 * time.Millisecond):
			return accounts
		}
	}
}

func getBalances(accounts []Account) ([]string, []uint64, error) {
	addresses := make([]string, len(accounts))
	balaces := make([]uint64, len(accounts))
	for i, account := range accounts {
		address, balance, err := account.ToBalance()
		if err != nil {
			return nil, nil, err
		}
		addresses[i] = address
		balaces[i] = balance
	}
	return addresses, balaces, nil
}
