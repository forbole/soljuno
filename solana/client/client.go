package client

import (
	"fmt"

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

	// GetSignaturesForAddress returns the metadata of transactions with the given address and config
	// An error is returned if the query fails
	GetSignaturesForAddress(
		address string,
		config clienttypes.GetSignaturesForAddressConfig,
	) ([]clienttypes.ConfirmedTransactionStatusWithSignature, error)

	// GetTransaction returns the transaction of the given hash
	// An error is returned if the query fails
	GetTransaction(hash string) (clienttypes.EncodedConfirmedTransactionWithStatusMeta, error)
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

func (c *Client) GetBlock(slot uint64) (clienttypes.BlockResult, error) {
	var block clienttypes.BlockResult
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

func (c *Client) GetVoteAccounts() (clienttypes.VoteAccounts, error) {
	var voteAccounts clienttypes.VoteAccounts
	err := c.rpcClient.CallFor(&voteAccounts, "getVoteAccounts")
	return voteAccounts, err
}

func (c *Client) GetVoteAccountsWithSlot() (uint64, clienttypes.VoteAccounts, error) {
	var slot uint64
	var voteAccounts clienttypes.VoteAccounts

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

func (c *Client) GetAccountInfo(address string) (clienttypes.AccountInfo, error) {
	var accountInfo clienttypes.AccountInfo
	err := c.rpcClient.CallFor(&accountInfo, "getAccountInfo", address, clienttypes.NewAccountInfoOption("base64"))
	return accountInfo, err
}

func (c *Client) GetLeaderSchedule(slot uint64) (clienttypes.LeaderSchedule, error) {
	var schedule clienttypes.LeaderSchedule
	err := c.rpcClient.CallFor(&schedule, "getLeaderSchedule", slot)
	return schedule, err
}

func (c *Client) GetSupplyInfo() (clienttypes.SupplyWithContext, error) {
	var supply clienttypes.SupplyWithContext
	err := c.rpcClient.CallFor(&supply, "getSupply", []interface{}{clienttypes.NewSupplyConfig(true)})
	return supply, err
}

func (c *Client) GetInflationRate() (clienttypes.InflationRate, error) {
	var rate clienttypes.InflationRate
	err := c.rpcClient.CallFor(&rate, "getInflationRate")
	if err != nil {
		return rate, err
	}
	return rate, nil
}

func (c *Client) GetEpochInfo() (clienttypes.EpochInfo, error) {
	var epochInfo clienttypes.EpochInfo
	err := c.rpcClient.CallFor(&epochInfo, "getEpochInfo")
	return epochInfo, err
}

func (c *Client) GetEpochSchedule() (clienttypes.EpochSchedule, error) {
	var epochSchedule clienttypes.EpochSchedule
	err := c.rpcClient.CallFor(&epochSchedule, "getEpochSchedule")
	return epochSchedule, err
}

func (c *Client) GetInflationGovernor() (clienttypes.InflationGovernor, error) {
	var governor clienttypes.InflationGovernor
	err := c.rpcClient.CallFor(&governor, "getInflationGovernor")
	return governor, err
}

func (c *Client) GetSignaturesForAddress(
	address string,
	config clienttypes.GetSignaturesForAddressConfig,
) ([]clienttypes.ConfirmedTransactionStatusWithSignature, error) {
	var sig []clienttypes.ConfirmedTransactionStatusWithSignature
	err := c.rpcClient.CallFor(&sig, "getSignaturesForAddress", address, config)
	return sig, err
}

func (c *Client) GetTransaction(hash string) (clienttypes.EncodedConfirmedTransactionWithStatusMeta, error) {
	var tx clienttypes.EncodedConfirmedTransactionWithStatusMeta
	err := c.rpcClient.CallFor(&tx, "getTransaction", hash, "json")
	return tx, err
}
