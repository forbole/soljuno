package client

import (
	"fmt"

	"github.com/forbole/soljuno/solana/client/types"
	jsonrpc "github.com/ybbus/jsonrpc/v2"
)

type Client interface {
	GetSlot() (uint64, error)
	GetBlock(uint64) (types.BlockResult, error)
	GetBlocks(uint64, uint64) ([]uint64, error)
	GetVoteAccounts() (types.VoteAccounts, error)
	GetAccountInfo(string) (types.AccountInfo, error)
	GetVoteAccountsWithSlot() (uint64, types.VoteAccounts, error)
}

type client struct {
	endpoint  string
	rpcClient jsonrpc.RPCClient
}

func NewClient(endpoint string) Client {
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
	err := c.rpcClient.CallFor(&slots, "getConfirmedBlocks", start, end)
	if err != nil {
		return slots, err
	}
	return slots, nil
}

func (c *client) GetVoteAccounts() (types.VoteAccounts, error) {
	var voteAccounts types.VoteAccounts
	err := c.rpcClient.CallFor(&voteAccounts, "getVoteAccounts")
	if err != nil {
		return voteAccounts, err
	}
	return voteAccounts, nil
}

func (c *client) GetVoteAccountsWithSlot() (uint64, types.VoteAccounts, error) {
	var slot uint64
	var voteAccounts types.VoteAccounts

	slotReq := jsonrpc.NewRequest("getSlot")
	voteAccountsReq := jsonrpc.NewRequest("getVoteAccounts")
	res, err := c.rpcClient.CallBatch(
		[]*jsonrpc.RPCRequest{
			slotReq,
			voteAccountsReq,
		},
	)
	if err != nil {
		return slot, voteAccounts, err
	}

	if res.HasError() {
		return slot, voteAccounts, fmt.Errorf("failed to get vote accounts or slot")
	}

	if err := res.GetByID(0).GetObject(&slot); err != nil {
		return slot, voteAccounts, err
	}

	if err := res.GetByID(1).GetObject(&voteAccounts); err != nil {
		return slot, voteAccounts, err
	}
	return slot, voteAccounts, nil
}

func (c *client) GetAccountInfo(address string) (types.AccountInfo, error) {
	var accountInfo types.AccountInfo
	err := c.rpcClient.CallFor(&accountInfo, "getAccountInfo", address, types.NewAccountInfoOption("base64"))
	if err != nil {
		return accountInfo, err
	}
	return accountInfo, nil
}
