package main

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/types"
)

type config struct{}

type rpc struct{}

func (rpc) GetClientName() string { return "" }
func (rpc) GetAddress() string    { return "https://api.mainnet-beta.solana.com" }

func (config) GetRPCConfig() types.RPCConfig { return rpc{} }

func (config) GetGrpcConfig() types.GrpcConfig { return nil }

func (config) GetCosmosConfig() types.CosmosConfig { return nil }

func (config) GetDatabaseConfig() types.DatabaseConfig { return nil }

func (config) GetLoggingConfig() types.LoggingConfig { return nil }

func (config) GetParsingConfig() types.ParsingConfig { return nil }

func (config) GetPruningConfig() types.PruningConfig { return nil }

func (config) GetTelemetryConfig() types.TelemetryConfig { return nil }

func main() {
	var c types.Config = config{}
	client, err := client.NewClientProxy(c)
	slot, err := client.Slots(0, 20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(slot)

}
