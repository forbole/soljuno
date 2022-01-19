package postgresql

import (
	"time"

	"github.com/forbole/soljuno/db"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

// type check to ensure interface is properly implemented
var _ db.BankDb = &Database{}

func (db *Database) SaveAccountBalances(slot uint64, accounts []string, balances []uint64) error {
	if len(balances) == 0 {
		return nil
	}
	paramsNumber := 3
	insertStmt := `INSERT INTO account_balance (address, slot, balance) VALUES`
	paramsStmt := ""
	conflictStmt := `
	ON CONFLICT (address) DO UPDATE
		SET slot = excluded.slot,
			balance = excluded.balance
	WHERE account_balance.slot <= excluded.slot
	`
	var params []interface{}
	for i, bal := range balances {
		bi := i * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)
		params = append(params, accounts[i], slot, bal)
	}
	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}

func (db *Database) SaveAccountTokenBalances(slot uint64, accounts []string, balances []clienttypes.TransactionTokenBalance) error {
	if len(balances) == 0 {
		return nil
	}
	paramsNumber := 3
	insertStmt := `INSERT INTO token_account_balance (address, slot, balance) VALUES`
	paramsStmt := ""
	conflictStmt := `
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
        balance = excluded.balance
WHERE token_account_balance.slot <= excluded.slot
	`
	var params []interface{}

	for i, bal := range balances {
		bi := i * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)
		params = append(params, accounts[bal.AccountIndex], slot, bal.UiTokenAmount.Amount)

	}

	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}

// ----------------------------------------------------------------

func (db *Database) SaveAccountHistoryBalances(timestamp time.Time, accounts []string, balances []uint64) error {
	if len(balances) == 0 {
		return nil
	}
	paramsNumber := 3
	insertStmt := `INSERT INTO account_balance_history (address, timestamp, balance) VALUES`
	paramsStmt := ""
	conflictStmt := ""
	var params []interface{}

	count := 0
	for i, bal := range balances {
		// Excute if the max params length will be reached
		if len(params)+paramsNumber >= MAX_PARAMS_LENGTH {
			err := db.insertWithParams(
				insertStmt,
				paramsStmt[:len(paramsStmt)-1],
				conflictStmt,
				params,
			)
			if err != nil {
				return err
			}
			count = 0
			paramsStmt = ""
			params = params[:0]
		}
		bi := count * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)
		params = append(params, accounts[i], timestamp, bal)
		count++
	}
	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}

func (db *Database) SaveAccountHistoryTokenBalances(timestamp time.Time, accounts []string, balances []clienttypes.TransactionTokenBalance) error {
	if len(balances) == 0 {
		return nil
	}
	paramsNumber := 3
	insertStmt := `INSERT INTO token_account_balance_history (address, timestamp, balance) VALUES`
	paramsStmt := ""
	conflictStmt := ""
	var params []interface{}

	count := 0
	for i, bal := range balances {
		// Excute if the max params length will be reached
		if len(params)+paramsNumber >= MAX_PARAMS_LENGTH {
			err := db.insertWithParams(
				insertStmt,
				paramsStmt[:len(paramsStmt)-1],
				conflictStmt,
				params,
			)
			if err != nil {
				return err
			}
			count = 0
			paramsStmt = ""
			params = params[:0]
		}
		bi := count * paramsNumber
		paramsStmt += getParamsStmt(bi, paramsNumber)
		params = append(params, accounts[i], timestamp, bal)
		count++
	}

	return db.insertWithParams(
		insertStmt,
		paramsStmt[:len(paramsStmt)-1],
		conflictStmt,
		params,
	)
}
