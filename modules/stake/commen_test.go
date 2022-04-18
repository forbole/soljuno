package stake_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules/stake"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

var _ db.StakeDb = &MockDb{}

type MockDb struct {
	isLatest bool
}

func NewDefaultMockDb() *MockDb {
	return &MockDb{isLatest: true}
}

func (db *MockDb) SaveStakeAccount(account dbtypes.StakeAccountRow) error          { return nil }
func (db *MockDb) DeleteStakeAccount(address string) error                         { return nil }
func (db *MockDb) SaveStakeDelegation(delegation dbtypes.StakeDelegationRow) error { return nil }
func (db *MockDb) SaveStakeLockup(lockup dbtypes.StakeLockupRow) error             { return nil }
func (db *MockDb) DeleteStakeDelegation(address string) error                      { return nil }

func (db *MockDb) CheckStakeAccountLatest(address string, currentSlot uint64) bool {
	return db.isLatest
}

func (m MockDb) GetCached() MockDb {
	return m
}

func (m *MockDb) WithLatest(isLatest bool) {
	m.isLatest = isLatest
}

// ----------------------------------------------------------------

var _ stake.ClientProxy = &MockClient{}

type MockClient struct {
	account clienttypes.AccountInfo
}

func NewDefaultMockClient() *MockClient {
	return &MockClient{}
}

func (m MockClient) GetCached() MockClient {
	return m
}

func (m *MockClient) WithNonceAccount(account clienttypes.AccountInfo) {
	m.account = account
}

func (m *MockClient) GetAccountInfo(address string) (clienttypes.AccountInfo, error) {
	return m.account, nil
}

// ----------------------------------------------------------------

type ModuleTestSuite struct {
	suite.Suite
	module *stake.Module
	db     *MockDb
	client *MockClient
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	suite.module = stake.NewModule(NewDefaultMockDb(), NewDefaultMockClient())
	suite.db = NewDefaultMockDb()
	suite.client = NewDefaultMockClient()
}
