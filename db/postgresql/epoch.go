package postgresql

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	solanatypes "github.com/forbole/soljuno/solana/types"
)

var _ db.EpochDb = &Database{}

func (db *Database) SaveSupplyInfo(supply dbtypes.SupplyInfoRow) error {
	stmt := `INSERT INTO supply_info 
	(epoch, total, circulating, non_circulating)
VALUES ($1, $2, $3, $4)
ON CONFLICT (one_row_id) DO UPDATE
	SET epoch = excluded.epoch,
		total = excluded.total,
		circulating = excluded.circulating,
		non_circulating = excluded.non_circulating
	WHERE supply_info.epoch <= excluded.epoch`
	_, err := db.Sqlx.Exec(
		stmt,
		supply.Epoch,
		supply.Total,
		supply.Circulating,
		supply.NonCirculating,
	)
	return err
}

func (db *Database) GetEpochBlocks(epoch uint64) ([]uint64, error) {
	start := epoch * solanatypes.SlotsInEpoch
	end := (epoch+1)*solanatypes.SlotsInEpoch - 1
	var blocks []dbtypes.BlockRow
	err := db.Sqlx.Select(&blocks, `SELECT slot FROM block WHERE slot >= $1 AND slot <= $2`, start, end)
	if err != nil {
		return nil, err
	}
	res := make([]uint64, len(blocks))
	for i, block := range blocks {
		res[i] = block.Slot
	}
	return res, nil
}
