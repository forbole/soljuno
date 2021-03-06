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
	err      error
}

func NewDefaultMockDb() *MockDb {
	return &MockDb{isLatest: true}
}

func (db *MockDb) SaveValidator(account dbtypes.VoteAccountRow) error { return db.err }
func (db *MockDb) DeleteValidator(address string) error               { return db.err }
func (db *MockDb) CheckValidatorLatest(address string, currentSlot uint64) bool {
	return db.isLatest
}

func (m MockDb) GetCached() MockDb {
	return m
}

func (m *MockDb) WithError(err error) {
	m.err = err
}

func (m *MockDb) WithLatest(isLatest bool) {
	m.isLatest = isLatest
}

// ----------------------------------------------------------------

var _ vote.ClientProxy = &MockClient{}

type MockClient struct {
	account clienttypes.AccountInfo
	err     error
}

func NewDefaultMockClient() *MockClient {
	return &MockClient{}
}

func (m MockClient) GetCached() MockClient {
	return m
}

func (m *MockClient) WithError(err error) {
	m.err = err
}

func (m *MockClient) WithAccount(account clienttypes.AccountInfo) {
	m.account = account
}

func (m *MockClient) GetAccountInfo(address string) (clienttypes.AccountInfo, error) {
	return m.account, m.err
}

func (m *MockClient) GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error) {
	return 0, clienttypes.VoteAccounts{
		Current:    []clienttypes.VoteAccount{{VotePubkey: "current"}},
		Delinquent: []clienttypes.VoteAccount{{VotePubkey: "delinquent"}},
	}, m.err
}

func (m *MockClient) GetLeaderSchedule(uint64) (clienttypes.LeaderSchedule, error) {
	return clienttypes.LeaderSchedule{"address": []int{0, 1}}, m.err
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
