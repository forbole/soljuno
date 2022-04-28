package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/lib/pq"
)

var _ db.TxDb = &Database{}

// SaveTxs implements db.TxDb
func (db *Database) SaveTxs(txs []dbtypes.TxRow) error {
	if len(txs) == 0 {
		return nil
	}
	err := db.saveTxs(txs)
	if err != nil {
		return err
	}
	return db.saveTxByAddressIndexes(txs)
}

func (db *Database) saveTxs(txs []dbtypes.TxRow) error {
	insertStmt := `INSERT INTO transaction (signature, slot, index, involved_accounts, success, fee, logs, num_instructions, partition_id) VALUES`
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 9
	params = make([]interface{}, 0, paramsNumber*len(txs))
	for _, tx := range txs {
		params = append(
			params,
			tx.Signature,
			tx.Slot,
			tx.Index,
			tx.InvolvedAccounts,
			tx.Success,
			tx.Fee,
			pq.Array(tx.Logs),
			tx.NumInstructions,
			tx.PartitionId,
		)
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

func (db *Database) saveTxByAddressIndexes(txs []dbtypes.TxRow) error {
	insertStmt := `INSERT INTO transaction_by_address (address, slot, signature, index, partition_id) VALUES`
	conflictStmt := `ON CONFLICT DO NOTHING`

	var params []interface{}
	paramsNumber := 5
	params = make([]interface{}, 0, paramsNumber*len(txs))
	for _, tx := range txs {
		for _, address := range tx.InvolvedAccounts {
			params = append(
				params,
				address,
				tx.Slot,
				tx.Signature,
				tx.Index,
				tx.PartitionId,
			)
		}
	}
	return db.InsertBatch(
		insertStmt,
		conflictStmt,
		params,
		paramsNumber,
	)
}

// CreateTxPartition implements db.TxDb
func (db *Database) CreateTxPartition(Id int) error {
	err := db.createPartition("transaction", Id)
	if err != nil {
		return err
	}
	return db.createPartition("transaction_by_address", Id)
}

// PruneTxsBeforeSlot implements db.TxDb
func (db *Database) PruneTxsBeforeSlot(slot uint64) error {
	for {
		name, err := db.getOldestPartitionBeforeSlot("transaction", slot)
		if err != nil {
			return err
		}
		if name == "" {
			break
		}

		err = db.dropPartition(name)
		if err != nil {
			return err
		}
	}

	for {
		name, err := db.getOldestPartitionBeforeSlot("transaction_by_address", slot)
		if err != nil {
			return err
		}
		if name == "" {
			break
		}

		err = db.dropPartition(name)
		if err != nil {
			return err
		}
	}
	return nil
}
