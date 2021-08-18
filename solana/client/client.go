package client

import (
	"github.com/desmos-labs/soljuno/solana/client/types"
	jsonrpc "github.com/ybbus/jsonrpc/v2"
)

type Client struct {
	endpoint  string
	rpcClient jsonrpc.RPCClient
}

func NewClient(endpoint string) *Client {
	rpcClient := jsonrpc.NewClient(endpoint)
	return &Client{
		endpoint:  endpoint,
		rpcClient: rpcClient,
	}
}

func (c *Client) Block(slot uint64) (types.Block, error) {
	var block types.Block
	err := c.rpcClient.CallFor(&block, "getBlock", slot)
	if err != nil {
		return block, err
	}
	return block, nil
}
