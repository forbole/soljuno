package fix_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules/fix"
	"github.com/forbole/soljuno/types"
)

type ModuleTestSuite struct {
	suite.Suite
	module *fix.Module
}

type MockDb struct {
	height uint64
	start  uint64
	end    uint64
}

func NewMockDb(height, start, end uint64) *MockDb {
	return &MockDb{height: height, start: start, end: end}
}

var _ db.FixMissingBlockDb = &MockDb{}

func (db MockDb) GetMissingHeight(start uint64, end uint64) (height uint64, err error) {
	return db.height, nil
}
func (db MockDb) GetMissingSlotRange(height uint64) (start uint64, end uint64, err error) {
	return db.start, db.end, nil
}

// ----------------------------------------------------------------

var _ fix.ClientProxy = &MockClient{}

type MockClient struct{}

func (client MockClient) GetBlocks(start uint64, end uint64) ([]uint64, error) {
	return []uint64{1}, nil
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	suite.module = fix.NewModule(NewMockDb(1, 0, 1), types.NewQueue(1), MockClient{})
}
