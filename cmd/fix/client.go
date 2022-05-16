package fix

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

func delay() {
	time.Sleep(100 * time.Millisecond)
}

func (cp *proxy) GetLatestSlot() (uint64, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetLatestSlot()
}

func (cp *proxy) GetBlocks(start uint64, end uint64) ([]uint64, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetBlocks(start, end)
}

func (cp *proxy) GetBlock(slot uint64) (clienttypes.BlockResult, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetBlock(slot)
}

func (cp *proxy) GetVoteAccounts() (clienttypes.VoteAccounts, error) {
	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetVoteAccounts()
}

func (cp *proxy) GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetVoteAccountsWithSlot()
}

func (cp *proxy) GetAccountInfo(address string) (clienttypes.AccountInfo, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetAccountInfo(address)
}

func (cp *proxy) GetSupplyInfo() (clienttypes.SupplyWithContext, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetSupplyInfo()
}

func (cp *proxy) GetInflationRate() (clienttypes.InflationRate, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetInflationRate()
}

func (cp *proxy) GetLeaderSchedule(slot uint64) (clienttypes.LeaderSchedule, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetLeaderSchedule(slot)
}

func (cp *proxy) GetEpochInfo() (clienttypes.EpochInfo, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetEpochInfo()
}

func (cp *proxy) GetEpochSchedule() (clienttypes.EpochSchedule, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetEpochSchedule()
}

func (cp *proxy) GetInflationGovernor() (clienttypes.InflationGovernor, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetInflationGovernor()
}

func (cp *proxy) GetSignaturesForAddress(
	address string,
	config clienttypes.GetSignaturesForAddressConfig,
) ([]clienttypes.ConfirmedTransactionStatusWithSignature, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetSignaturesForAddress(address, config)
}

func (cp *proxy) GetTransaction(signature string) (clienttypes.EncodedConfirmedTransactionWithStatusMeta, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetTransaction(signature)
}

func (cp *proxy) GetSlotLeaders(slot uint64, limit uint64) ([]string, error) {
	defer delay()

	cp.mtx.Lock()
	defer cp.mtx.Unlock()
	return cp.rpcClient.GetSlotLeaders(slot, limit)
}
