package client

import (
	"context"

	"github.com/forbole/soljuno/types"

	client "github.com/forbole/soljuno/solana/client"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

type Proxy interface {
	// Block queries for a block by slot. An error is returned if the query fails.
	Block(uint64) (clienttypes.BlockResult, error)

	// LatestSlot returns the latest slot on the active chain. An error
	// is returned if the query fails.
	LatestSlot() (uint64, error)

	// Slots returns the slot of confirmed blocks between the given start and end on the active chain.
	// An error is returned if the query fails.
	Slots(uint64, uint64) ([]uint64, error)

	// AccountInfo returns the information of the given account in the current chain.
	// An error is returned if the query fails
	AccountInfo(address string) (clienttypes.AccountInfo, error)

	// Validators returns vote accounts of validators in the current chain.
	// An error is returned if the query fails
	ValidatorsWithSlot() (uint64, clienttypes.VoteAccounts, error)

	// Supply returns supply info in the current chain.
	// An error is returned if the query fails
	Supply() (clienttypes.SupplyWithContext, error)

	// InflationRate returns inflation rate in the current chain.
	// An error is returned if the query fails
	InflationRate() (clienttypes.InflationRate, error)

	// GetLeaderSchedule returns epoch leader schedule of the given slot in the current chain.
	// An error is returned if the query fails
	GetLeaderSchedule(slot uint64) (clienttypes.LeaderSchedule, error)
}

// proxy implements a wrapper around both a Tendermint RPC client and a
// Cosmos Sdk REST client that allows for essential data queries.
type proxy struct {
	ctx context.Context

	rpcClient client.Client
}

// NewClientProxy allows to build a new Proxy instance
func NewClientProxy(cfg types.Config) (Proxy, error) {
	rpcClient := client.NewClient(cfg.GetRPCConfig().GetAddress())

	return &proxy{
		ctx:       context.Background(),
		rpcClient: rpcClient,
	}, nil
}

func (cp *proxy) LatestSlot() (uint64, error) {
	return cp.rpcClient.GetSlot()
}

func (cp *proxy) Slots(start uint64, end uint64) ([]uint64, error) {
	return cp.rpcClient.GetBlocks(start, end)
}

func (cp *proxy) Block(slot uint64) (clienttypes.BlockResult, error) {
	return cp.rpcClient.GetBlock(slot)
}

func (cp *proxy) Validators() (clienttypes.VoteAccounts, error) {
	return cp.rpcClient.GetVoteAccounts()
}

func (cp *proxy) ValidatorsWithSlot() (uint64, clienttypes.VoteAccounts, error) {
	return cp.rpcClient.GetVoteAccountsWithSlot()
}

func (cp *proxy) AccountInfo(address string) (clienttypes.AccountInfo, error) {
	return cp.rpcClient.GetAccountInfo(address)
}

func (cp *proxy) Supply() (clienttypes.SupplyWithContext, error) {
	return cp.rpcClient.GetSupplyInfo()
}

func (cp *proxy) InflationRate() (clienttypes.InflationRate, error) {
	return cp.rpcClient.GetInflationRate()
}

func (cp *proxy) GetLeaderSchedule(slot uint64) (clienttypes.LeaderSchedule, error) {
	return cp.rpcClient.GetLeaderSchedule(slot)
}
