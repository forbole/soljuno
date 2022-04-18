package system_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules/system"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/types/pool"
)

type ModuleTestSuit struct {
	suite.Suite
	module *system.Module
}

type MockDb struct{}

var _ db.SystemDb = &MockDb{}

type MockClient struct{}

var _ system.ClientProxy = &MockClient{}

func (*MockClient) AccountInfo(address string) {
	return clienttypes.AccountInfo{}
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuit))
}

func (suite *ModuleTestSuit) SetupTest() {
	pool, err := pool.NewDefaultPool(10)
	suite.Require().NoError(err)
	suite.module = system.NewModule(MockDb{}, pool)
}
