package postgresql

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

// type check to ensure interface is properly implemented
var _ db.BankDb = &Database{}

func (db *Database) SaveAccountBalances(slot uint64, accounts []string, balances []uint64) error {
	// Store up-to-date data
	err := db.saveUpToDateBalances(3, slot, accounts, balances)
	if err != nil {
		return fmt.Errorf("error while storing up-to-date balances: %s", err)
	}
	return nil
}

func (db *Database) saveUpToDateBalances(paramsNumber int, slot uint64, accounts []string, balances []uint64) error {
	if len(balances) == 0 {
		return nil
	}

	stmt := `INSERT INTO account_balance (address, slot, balance) VALUES`
	var params []interface{}

	for i, bal := range balances {
		bi := i * paramsNumber
		stmt += fmt.Sprintf("($%d, $%d, $%d),", bi+1, bi+2, bi+3)
		params = append(params, accounts[i], slot, bal)

	}

	stmt = stmt[:len(stmt)-1]

	stmt += `
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
        balance = excluded.balance
WHERE account_balance.slot <= excluded.slot
`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) SaveAccountTokenBalances(slot uint64, accounts []string, balances []clienttypes.TransactionTokenBalance) error {
	err := db.saveUpToDateTokenBalances(3, slot, accounts, balances)
	if err != nil {
		return fmt.Errorf("error while storing up-to-date token balances: %s", err)
	}
	return nil
}

func (db *Database) saveUpToDateTokenBalances(paramsNumber int, slot uint64, accounts []string, balances []clienttypes.TransactionTokenBalance) error {
	if len(balances) == 0 {
		return nil
	}

	stmt := `INSERT INTO token_account_balance (address, slot, balance) VALUES`
	var params []interface{}

	for i, bal := range balances {
		bi := i * paramsNumber
		stmt += fmt.Sprintf("($%d, $%d, $%d),", bi+1, bi+2, bi+3)
		params = append(params, accounts[bal.AccountIndex], slot, bal.UiTokenAmount.Amount)

	}

	stmt = stmt[:len(stmt)-1]

	stmt += `
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
        balance = excluded.balance
WHERE account_balance.slot <= excluded.slot
`

	_, err := db.Sqlx.Exec(stmt, params...)
	if err != nil {
		return err
	}
	return nil
}
