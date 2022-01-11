package postgresql_test

import (
	"time"

	solanatypes "github.com/forbole/soljuno/solana/types"
	"github.com/forbole/soljuno/types"
)

func (suite *DbTestSuite) TestPrune() {
	type Row struct {
		Slot uint64 `db:"slot"`
	}
	var rows []Row

	msg := types.NewMessage("tx", 1, 0, 0, "program", []string{"address1", "address2"}, "", solanatypes.NewParsedInstruction("unknown", nil))
	tx := types.NewTx("tx", 1, nil, 0, nil, []types.Message{msg}, nil, nil, nil)
	err := suite.database.SaveBlock(types.NewBlock(1, 1, "block", "proposer", time.Now(), []types.Tx{tx}))
	suite.Require().NoError(err)
	err = suite.database.SaveTxs([]types.Tx{tx})
	suite.Require().NoError(err)

	// Prune nothing
	err = suite.database.PruneTxsBySlot(0)
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&rows, "SELECT slot FROM transaction")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	rows = []Row{}

	// Prune before slot 1
	err = suite.database.PruneTxsBySlot(1)
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&rows, "SELECT slot FROM transaction")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
	rows = []Row{}

}
