package postgresql

import (
	"strconv"
	"time"

	"github.com/forbole/soljuno/db"
)

// type check to ensure interface is properly implemented
var _ db.BankDb = &Database{}

func (db *Database) SaveAccountBalances(slot uint64, accounts []string, balances []uint64) error {
	if len(balances) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO account_balance (address, slot, balance) VALUES`
	conflictStmt := `
	ON CONFLICT (address) DO UPDATE
		SET slot = excluded.slot,
			balance = excluded.balance
	WHERE account_balance.slot <= excluded.slot
	`
	var params []interface{}
	paramsNumber := 3
	params = make([]interface{}, 0, paramsNumber*len(balances))
	for i, bal := range balances {
		params = append(params, accounts[i], slot, bal)
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

func (db *Database) SaveAccountTokenBalances(slot uint64, accounts []string, balances []uint64) error {
	if len(balances) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO token_account_balance (address, slot, balance) VALUES`
	conflictStmt := `
ON CONFLICT (address) DO UPDATE
    SET slot = excluded.slot,
        balance = excluded.balance
WHERE token_account_balance.slot <= excluded.slot
	`
	var params []interface{}
	paramsNumber := 3
	params = make([]interface{}, 0, paramsNumber*len(balances))
	for i, bal := range balances {
		params = append(params, accounts[i], slot, strconv.FormatUint(bal, 10))

	}

	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

// ----------------------------------------------------------------

func (db *Database) SaveAccountHistoryBalances(timestamp time.Time, accounts []string, balances []uint64) error {
	if len(balances) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO account_balance_history (address, timestamp, balance) VALUES`
	conflictStmt := ""
	var params []interface{}
	paramsNumber := 3
	params = make([]interface{}, 0, paramsNumber*len(balances))
	for i, bal := range balances {
		params = append(params, accounts[i], timestamp, strconv.FormatUint(bal, 10))
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

func (db *Database) SaveAccountHistoryTokenBalances(timestamp time.Time, accounts []string, balances []uint64) error {
	if len(balances) == 0 {
		return nil
	}
	insertStmt := `INSERT INTO token_account_balance_history (address, timestamp, balance) VALUES`
	conflictStmt := ""
	var params []interface{}
	paramsNumber := 3
	params = make([]interface{}, 0, paramsNumber*len(balances))

	count := 0
	for i, bal := range balances {
		// Excute if the max params length will be reached
		params = append(params, accounts[i], timestamp, strconv.FormatUint(bal, 10))
		count++
	}

	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}
