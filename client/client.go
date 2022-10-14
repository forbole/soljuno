package client

import (
	"context"

	"github.com/forbole/soljuno/types"

	client "github.com/forbole/soljuno/solana/client"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

type proxy struct {
	ctx context.Context

	rpcClient client.ClientProxy
}

// NewClientProxy allows to build a new Proxy instance
func NewClientProxy(cfg types.Config) (client.ClientProxy, error) {
	rpcClient := client.NewClientProxy(cfg.GetRPCConfig().GetAddress())
	return &proxy{
		ctx:       context.Background(),
		rpcClient: rpcClient,
	}, nil
}

func (cp *proxy) GetLatestSlot() (uint64, error) {
	return cp.rpcClient.GetLatestSlot()
}

func (cp *proxy) GetBlocks(start uint64, end uint64) ([]uint64, error) {
	return cp.rpcClient.GetBlocks(start, end)
}

func (cp *proxy) GetBlock(slot uint64, config clienttypes.BlockConfig) (clienttypes.BlockResult, error) {
	return cp.rpcClient.GetBlock(slot, config)
}

func (cp *proxy) GetVoteAccounts() (clienttypes.VoteAccounts, error) {
	return cp.rpcClient.GetVoteAccounts()
}

func (cp *proxy) GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error) {
	return cp.rpcClient.GetVoteAccountsWithSlot()
}

func (cp *proxy) GetAccountInfo(address string) (clienttypes.AccountInfo, error) {
	return cp.rpcClient.GetAccountInfo(address)
}

func (cp *proxy) GetSupplyInfo() (clienttypes.SupplyWithContext, error) {
	return cp.rpcClient.GetSupplyInfo()
}

func (cp *proxy) GetInflationRate() (clienttypes.InflationRate, error) {
	return cp.rpcClient.GetInflationRate()
}

func (cp *proxy) GetLeaderSchedule(slot uint64) (clienttypes.LeaderSchedule, error) {
	return cp.rpcClient.GetLeaderSchedule(slot)
}

func (cp *proxy) GetEpochInfo() (clienttypes.EpochInfo, error) {
	return cp.rpcClient.GetEpochInfo()
}

func (cp *proxy) GetEpochSchedule() (clienttypes.EpochSchedule, error) {
	return cp.rpcClient.GetEpochSchedule()
}

func (cp *proxy) GetInflationGovernor() (clienttypes.InflationGovernor, error) {
	return cp.rpcClient.GetInflationGovernor()
}

func (cp *proxy) GetSignaturesForAddress(
	address string,
	config clienttypes.GetSignaturesForAddressConfig,
) ([]clienttypes.ConfirmedTransactionStatusWithSignature, error) {
	return cp.rpcClient.GetSignaturesForAddress(address, config)
}

func (cp *proxy) GetTransaction(signature string) (clienttypes.EncodedConfirmedTransactionWithStatusMeta, error) {
	return cp.rpcClient.GetTransaction(signature)
}

func (cp *proxy) GetSlotLeaders(slot uint64, limit uint64) ([]string, error) {
	return cp.rpcClient.GetSlotLeaders(slot, limit)
}
