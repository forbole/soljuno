package client

import (
	"fmt"

	"github.com/forbole/soljuno/solana/client/types"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	jsonrpc "github.com/ybbus/jsonrpc/v2"
)

type ClientProxy interface {
	// GetLatestSlot returns the latest slot on the active chain. An error
	// is returned if the query fails.
	GetLatestSlot() (uint64, error)

	// GetBlock queries for a block by slot. An error is returned if the query fails.
	GetBlock(uint64) (clienttypes.BlockResult, error)

	// GetBlocks returns the slot of confirmed blocks between the given start and end on the active chain.
	// An error is returned if the query fails.
	GetBlocks(uint64, uint64) ([]uint64, error)

	GetVoteAccounts() (clienttypes.VoteAccounts, error)

	// GetAccountInfo returns the information of the given account in the current chain.
	// An error is returned if the query fails
	GetAccountInfo(string) (clienttypes.AccountInfo, error)

	// GetVoteAccountsWithSlot returns vote accounts of validators in the current chain.
	// An error is returned if the query fails
	GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error)

	// GetLeaderSchedule returns epoch leader schedule of the given slot in the current chain.
	// An error is returned if the query fails
	GetLeaderSchedule(slot uint64) (clienttypes.LeaderSchedule, error)

	// GetSupplyInfo returns supply info in the current chain.
	// An error is returned if the query fails
	GetSupplyInfo() (clienttypes.SupplyWithContext, error)

	// GetInflationRate returns inflation rate in the current chain.
	// An error is returned if the query fails
	GetInflationRate() (clienttypes.InflationRate, error)

	// GetEpochInfo returns epoch info in the current chain.
	// An error is returned if the query fails
	GetEpochInfo() (clienttypes.EpochInfo, error)

	// GetEpochSchedule returns epoch schedule in the current chain.
	// An error is returned if the query fails
	GetEpochSchedule() (clienttypes.EpochSchedule, error)

	// GetInflationGovernor return inflation governor in the current chain
	// An error is returned if the query fails
	GetInflationGovernor() (clienttypes.InflationGovernor, error)
}

type Client struct {
	endpoint  string
	rpcClient jsonrpc.RPCClient
}

func NewClientProxy(endpoint string) ClientProxy {
	rpcClient := jsonrpc.NewClient(endpoint)
	return &Client{
		endpoint:  endpoint,
		rpcClient: rpcClient,
	}
}

func (c *Client) GetBlock(slot uint64) (types.BlockResult, error) {
	var block types.BlockResult
	err := c.rpcClient.CallFor(&block, "getBlock", slot)
	return block, err
}

func (c *Client) GetLatestSlot() (uint64, error) {
	var slot uint64
	err := c.rpcClient.CallFor(&slot, "getSlot")
	return slot, err
}

func (c *Client) GetBlocks(start uint64, end uint64) ([]uint64, error) {
	var slots []uint64
	err := c.rpcClient.CallFor(&slots, "getBlocks", start, end)
	return slots, err
}

func (c *Client) GetVoteAccounts() (types.VoteAccounts, error) {
	var voteAccounts types.VoteAccounts
	err := c.rpcClient.CallFor(&voteAccounts, "getVoteAccounts")
	return voteAccounts, err
}

func (c *Client) GetVoteAccountsWithSlot() (uint64, types.VoteAccounts, error) {
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
	err = res.GetByID(1).GetObject(&voteAccounts)
	return slot, voteAccounts, err
}

func (c *Client) GetAccountInfo(address string) (types.AccountInfo, error) {
	var accountInfo types.AccountInfo
	err := c.rpcClient.CallFor(&accountInfo, "getAccountInfo", address, types.NewAccountInfoOption("base64"))
	return accountInfo, err
}

func (c *Client) GetLeaderSchedule(slot uint64) (types.LeaderSchedule, error) {
	var schedule types.LeaderSchedule
	err := c.rpcClient.CallFor(&schedule, "getLeaderSchedule", slot)
	return schedule, err
}

func (c *Client) GetSupplyInfo() (types.SupplyWithContext, error) {
	var supply types.SupplyWithContext
	err := c.rpcClient.CallFor(&supply, "getSupply", []interface{}{types.NewSupplyConfig(true)})
	return supply, err
}

func (c *Client) GetInflationRate() (types.InflationRate, error) {
	var rate types.InflationRate
	err := c.rpcClient.CallFor(&rate, "getInflationRate")
	if err != nil {
		return rate, err
	}
	return rate, nil
}

func (c *Client) GetEpochInfo() (types.EpochInfo, error) {
	var epochInfo types.EpochInfo
	err := c.rpcClient.CallFor(&epochInfo, "getEpochInfo")
	return epochInfo, err
}

func (c *Client) GetEpochSchedule() (types.EpochSchedule, error) {
	var epochSchedule types.EpochSchedule
	err := c.rpcClient.CallFor(&epochSchedule, "getEpochSchedule")
	return epochSchedule, err
}

func (c *Client) GetInflationGovernor() (types.InflationGovernor, error) {
	var governor types.InflationGovernor
	err := c.rpcClient.CallFor(&governor, "getInflationGovernor")
	return governor, err
}
