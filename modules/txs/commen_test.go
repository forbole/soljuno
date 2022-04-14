package txs_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules/txs"
	"github.com/forbole/soljuno/types/pool"
)

type ModuleTestSuit struct {
	suite.Suite
	module *txs.Module
}

type MockDb struct{}

var _ db.TxDb = &MockDb{}

func (db MockDb) SaveTxs(txs []dbtypes.TxRow) error    { return nil }
func (db MockDb) CreateTxPartition(ID int) error       { return nil }
func (db MockDb) PruneTxsBeforeSlot(slot uint64) error { return nil }

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuit))
}

func (suite *ModuleTestSuit) SetupTest() {
	pool, err := pool.NewDefaultPool(10)
	suite.Require().NoError(err)
	suite.module = txs.NewModule(MockDb{}, pool)
}
