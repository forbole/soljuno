package client

import (
	"context"

	"github.com/forbole/soljuno/types"

	client "github.com/forbole/soljuno/solana/client"
	clienttype "github.com/forbole/soljuno/solana/client/types"
)

type Proxy interface {
	// Block queries for a block by slot. An error is returned if the query fails.
	Block(uint64) (clienttype.BlockResult, error)

	// LatestSlot returns the latest slot on the active chain. An error
	// is returned if the query fails.
	LatestSlot() (uint64, error)

	// Slots returns the slot of confirmed blocks between the given start and end on the active chain.
	// An error is returned if the query fails.
	Slots(uint64, uint64) ([]uint64, error)

	// Validators returns vote accounts of validators in the current chain.
	// An error is returned if the query fails
	Validators() (clienttype.VoteAccounts, error)
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
	slot, err := cp.rpcClient.GetSlot()
	if err != nil {
		return 0, err
	}

	return slot, nil
}

func (cp *proxy) Slots(start uint64, end uint64) ([]uint64, error) {
	slots, err := cp.rpcClient.GetBlocks(start, end)
	if err != nil {
		return []uint64{}, err
	}

	return slots, nil
}

func (cp *proxy) Block(slot uint64) (clienttype.BlockResult, error) {
	return cp.rpcClient.GetBlock(slot)
}

func (cp *proxy) Validators() (clienttype.VoteAccounts, error) {
	validators, err := cp.rpcClient.GetVoteAccounts()
	if err != nil {
		return clienttype.VoteAccounts{}, err
	}
	return validators, nil
}
