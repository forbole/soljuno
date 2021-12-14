package epoch

import (
	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/db/types"
	dbtypes "github.com/forbole/soljuno/db/types"
)

func updateInflationRate(epoch uint64, db db.EpochDb, client client.Proxy) error {
	inflation, err := client.InflationRate()
	if err != nil {
		return err
	}
	return db.SaveInflationRate(
		dbtypes.NewInflationRateRow(
			epoch,
			inflation.Total,
			inflation.Foundation,
			inflation.Validator,
		),
	)
}

func updateSupplyInfo(epoch uint64, db db.EpochDb, client client.Proxy) error {
	supply, err := client.Supply()
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

func updateInflationGovernorParam(epoch uint64, db db.EpochDb, client client.Proxy) error {
	governor, err := client.InflationGovernor()
	if err != nil {
		return err
	}
	return db.SaveInflationGovernorParam(
		types.NewInflationGovernorParamRow(
			epoch,
			governor.Initial,
			governor.Terminal,
			governor.Taper,
			governor.Foundation,
			governor.FoundationTerminal,
		),
	)
}

func updateEpochScheduleParam(epoch uint64, db db.EpochDb, client client.Proxy) error {
	schedule, err := client.EpochSchedule()
	if err != nil {
		return err
	}
	return db.SaveEpochScheduleParam(
		dbtypes.NewEpochScheduleParamRow(
			epoch,
			schedule.SlotsPerEpoch,
			schedule.FirstNormalEpoch,
			schedule.FirstNormalSlot,
			schedule.Warmup,
		),
	)
}