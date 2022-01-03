package epoch

import (
	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/solana/client"
)

func updateSupplyInfo(epoch uint64, db db.EpochDb, client client.ClientProxy) error {
	supply, err := client.GetSupplyInfo()
	if err != nil {
		return err
	}
	return db.SaveSupplyInfo(
		dbtypes.NewSupplyInfoRow(
			epoch,
			supply.Value.Total,
			supply.Value.Circulating,
			supply.Value.NonCirculating,
		),
	)
}
