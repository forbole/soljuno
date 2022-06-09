package block

import (
	"sync"
	"time"

	"github.com/forbole/soljuno/types"

	client "github.com/forbole/soljuno/solana/client"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

type proxy struct {
	mtx       sync.Mutex
	rpcClient client.ClientProxy
}

// NewClientProxy allows to build a new Proxy instance
func NewClientProxy(cfg types.Config) (client.ClientProxy, error) {
	rpcClient := client.NewClientProxy(cfg.GetRPCConfig().GetAddress())
	return &proxy{
		rpcClient: rpcClient,
	}, nil
}

// Set delay in order to parse from rate limited node
func delay() {
	time.Sleep(100 * time.Millisecond)
}

func (cp *proxy) GetLatestSlot() (uint64, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetLatestSlot()
}

func (cp *proxy) GetBlocks(start uint64, end uint64) ([]uint64, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetBlocks(start, end)
}

func (cp *proxy) GetBlock(slot uint64) (clienttypes.BlockResult, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetBlock(slot)
}

func (cp *proxy) GetVoteAccounts() (clienttypes.VoteAccounts, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetVoteAccounts()
}

func (cp *proxy) GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetVoteAccountsWithSlot()
}

func (cp *proxy) GetAccountInfo(address string) (clienttypes.AccountInfo, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetAccountInfo(address)
}

func (cp *proxy) GetSupplyInfo() (clienttypes.SupplyWithContext, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetSupplyInfo()
}

func (cp *proxy) GetInflationRate() (clienttypes.InflationRate, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetInflationRate()
}

func (cp *proxy) GetLeaderSchedule(slot uint64) (clienttypes.LeaderSchedule, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetLeaderSchedule(slot)
}

func (cp *proxy) GetEpochInfo() (clienttypes.EpochInfo, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetEpochInfo()
}

func (cp *proxy) GetEpochSchedule() (clienttypes.EpochSchedule, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetEpochSchedule()
}

func (cp *proxy) GetInflationGovernor() (clienttypes.InflationGovernor, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetInflationGovernor()
}

func (cp *proxy) GetSignaturesForAddress(
	address string,
	config clienttypes.GetSignaturesForAddressConfig,
) ([]clienttypes.ConfirmedTransactionStatusWithSignature, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetSignaturesForAddress(address, config)
}

func (cp *proxy) GetTransaction(signature string) (clienttypes.EncodedConfirmedTransactionWithStatusMeta, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetTransaction(signature)
}

func (cp *proxy) GetSlotLeaders(slot uint64, limit uint64) ([]string, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	defer delay()
	return cp.rpcClient.GetSlotLeaders(slot, limit)
}
