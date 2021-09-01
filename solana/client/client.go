package client

import (
	"github.com/forbole/soljuno/solana/client/types"
	jsonrpc "github.com/ybbus/jsonrpc/v2"
)

type Client interface {
	GetSlot() (uint64, error)
	GetBlock(uint64) (types.BlockResult, error)
	GetBlocks(uint64, uint64) ([]uint64, error)
	GetVoteAccounts() (types.VoteAccounts, error)
}

type client struct {
	endpoint  string
	rpcClient jsonrpc.RPCClient
}

func NewClient(endpoint string) *client {
	rpcClient := jsonrpc.NewClient(endpoint)
	return &client{
		endpoint:  endpoint,
		rpcClient: rpcClient,
	}
}

func (c *client) GetBlock(slot uint64) (types.BlockResult, error) {
	var block types.BlockResult
	err := c.rpcClient.CallFor(&block, "getBlock", slot)
	if err != nil {
		return block, err
	}
	return block, nil
}

func (c *client) GetSlot() (uint64, error) {
	var slot uint64
	err := c.rpcClient.CallFor(&slot, "getSlot")
	if err != nil {
		return slot, err
	}
	return slot, nil
}

func (c *client) GetBlocks(start uint64, end uint64) ([]uint64, error) {
	var slots []uint64
	err := c.rpcClient.CallFor(&slots, "getBlocks", start, end)
	if err != nil {
		return slots, err
	}
	return slots, nil
}

func (c *client) GetVoteAccounts() (types.VoteAccounts, error) {
	var validators types.VoteAccounts
	err := c.rpcClient.CallFor(&validators, "getVoteAccounts")
	if err != nil {
		return validators, err
	}
	return validators, nil
}
