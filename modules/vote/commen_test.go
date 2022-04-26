package vote_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/forbole/soljuno/db"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules/vote"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

var _ db.VoteDb = &MockDb{}

type MockDb struct {
	isLatest bool
}

func NewDefaultMockDb() *MockDb {
	return &MockDb{isLatest: true}
}

func (db *MockDb) SaveValidator(account dbtypes.VoteAccountRow) error                    { return nil }
func (db *MockDb) SaveValidatorStatuses(statuses []dbtypes.ValidatorStatusRow) error     { return nil }
func (db *MockDb) GetEpochProducedBlocks(epoch uint64) ([]uint64, error)                 { return []uint64{0}, nil }
func (db *MockDb) SaveValidatorSkipRates(skipRates []dbtypes.ValidatorSkipRateRow) error { return nil }
func (db *MockDb) SaveHistoryValidatorSkipRates(skipRates []dbtypes.ValidatorSkipRateRow) error {
	return nil
}

func (db *MockDb) CheckValidatorLatest(address string, currentSlot uint64) bool {
	return db.isLatest
}

func (m MockDb) GetCached() MockDb {
	return m
}

func (m *MockDb) WithLatest(isLatest bool) {
	m.isLatest = isLatest
}

// ----------------------------------------------------------------

var _ vote.ClientProxy = &MockClient{}

type MockClient struct {
	account clienttypes.AccountInfo
}

func NewDefaultMockClient() *MockClient {
	return &MockClient{}
}

func (m MockClient) GetCached() MockClient {
	return m
}

func (m *MockClient) WithAccount(account clienttypes.AccountInfo) {
	m.account = account
}

func (m *MockClient) GetAccountInfo(address string) (clienttypes.AccountInfo, error) {
	return m.account, nil
}

func (m *MockClient) GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error) {
	return 0, clienttypes.VoteAccounts{}, nil
}

func (m *MockClient) GetLeaderSchedule(uint64) (clienttypes.LeaderSchedule, error) {
	return clienttypes.LeaderSchedule{}, nil
}

// ----------------------------------------------------------------

type ModuleTestSuite struct {
	suite.Suite
	module *vote.Module
	db     *MockDb
	client *MockClient
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	suite.module = vote.NewModule(NewDefaultMockDb(), NewDefaultMockClient())
	suite.db = NewDefaultMockDb()
	suite.client = NewDefaultMockClient()
}
