package instructions_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules/instructions"
	"github.com/forbole/soljuno/types/pool"
)

type ModuleTestSuit struct {
	suite.Suite
	module *instructions.Module
}

type MockDb struct{}

var _ db.InstructionDb = &MockDb{}

func (db MockDb) SaveInstructions(instructions []dbtypes.InstructionRow) error { return nil }
func (db MockDb) CreateInstructionPartition(Id int) error                      { return nil }
func (db MockDb) PruneInstructionsBeforeSlot(slot uint64) error                { return nil }

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuit))
}

func (suite *ModuleTestSuit) SetupTest() {
	pool, err := pool.NewDefaultPool(10)
	suite.Require().NoError(err)
	suite.module = instructions.NewModule(MockDb{}, pool)
}
