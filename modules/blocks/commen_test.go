package blocks_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules/blocks"
)

type ModuleTestSuite struct {
	suite.Suite
	module *blocks.Module
}

type MockDb struct{}

var _ db.BlockDb = &MockDb{}

func (db MockDb) HasBlock(slot uint64) (bool, error)     { return false, nil }
func (db MockDb) SaveBlock(block dbtypes.BlockRow) error { return nil }

func TestModuleTestSuitee(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	suite.module = blocks.NewModule(MockDb{})
}
