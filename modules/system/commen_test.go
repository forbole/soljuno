package system_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules/system"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

var _ db.SystemDb = &MockDb{}

type MockDb struct {
	isLatest bool
	isBroken bool
}

func NewDefaultMockDb() *MockDb {
	return &MockDb{isLatest: true}
}

func (db *MockDb) SaveNonceAccount(nonce dbtypes.NonceAccountRow) error { return nil }
func (db *MockDb) DeleteNonceAccount(address string) error              { return nil }
func (db *MockDb) CheckNonceAccountLatest(address string, currentSlot uint64) bool {
	return db.isLatest
}

func (m MockDb) GetCached() MockDb {
	return m
}

func (m *MockDb) WithLatest(isLatest bool) {
	m.isLatest = isLatest
}

func (m *MockDb) WithBroken(isBroken bool) {
	m.isBroken = isBroken
}

// ----------------------------------------------------------------

var _ system.ClientProxy = &MockClient{}

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
	module *system.Module
	db     *MockDb
	client *MockClient
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	suite.module = system.NewModule(NewDefaultMockDb(), NewDefaultMockClient())
	suite.db = NewDefaultMockDb()
	suite.client = NewDefaultMockClient()
}
