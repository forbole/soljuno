package bank

import (
	"github.com/forbole/soljuno/db"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/types"
)

func HandleBlock(block types.Block, db db.BankDb) error {
	var accounts []string
	var balances []uint64
	var tokenAccounts []string
	var tokenBalances []clienttypes.TransactionTokenBalance
	accountMap := make(map[string]int)
	tokenAccountMap := make(map[string]int)

	accountIndex := 0
	// Gather balances from rewards
	for _, reward := range block.Rewards {
		if index, ok := accountMap[reward.Pubkey]; ok {
			balances[index] = reward.PostBalance
			continue
		}
		accounts = append(accounts, reward.Pubkey)
		balances = append(balances, reward.PostBalance)
		accountMap[reward.Pubkey] = accountIndex
		accountIndex++
	}

	tokenAccountIndex := 0
	for _, tx := range block.Txs {
		// Gather balances from tx
		for i, pubkey := range tx.Accounts {
			if index, ok := accountMap[pubkey]; ok {
				balances[index] = tx.PostBalances[i]
				continue
			}
			accounts = append(accounts, pubkey)
			balances = append(balances, tx.PostBalances[i])
			accountMap[pubkey] = accountIndex
			accountIndex++
		}

		// Gather token balances from tx
		for _, balance := range tx.PostTokenBalances {
			pubkey := tx.Accounts[balance.AccountIndex]
			if index, ok := tokenAccountMap[pubkey]; ok {
				tokenBalances[index].UiTokenAmount = balance.UiTokenAmount
				tokenBalances[index].Mint = balance.Mint
				continue
			}
			balance.AccountIndex = uint(tokenAccountIndex)
			tokenBalances = append(tokenBalances, balance)
			tokenAccounts = append(tokenAccounts, pubkey)
			tokenAccountMap[pubkey] = tokenAccountIndex
			tokenAccountIndex++
		}
	}
	err := db.SaveAccountBalances(block.Slot, accounts, balances)
	if err != nil {
		return err
	}
	return db.SaveAccountTokenBalances(block.Slot, tokenAccounts, tokenBalances)
}
