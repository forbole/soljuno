package parse

import (
	"fmt"

	"github.com/forbole/soljuno/db"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/modules"
	modsregistrar "github.com/forbole/soljuno/modules/registrar"
	"github.com/forbole/soljuno/types"
)

// GetParsingContext setups all the things that should be later passed to StartParsing in order
// to parse the chain data properly.
func GetParsingContext(parseConfig *Config) (*Context, error) {
	// Get the global config
	cfg := types.Cfg

	// Setup the SDK configuration
	sdkConfig := sdk.GetConfig()
	sdkConfig.Seal()

	// Get the database
	databaseCtx := db.NewContext(cfg.GetDatabaseConfig(), parseConfig.GetLogger())
	database, err := parseConfig.GetDBBuilder()(databaseCtx)
	if err != nil {
		return nil, err
	}

	// Init the client
	cp, err := client.NewClientProxy(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to start client: %s", err)
	}

	// Setup the logging
	err = parseConfig.GetLogger().SetLogFormat(cfg.GetLoggingConfig().GetLogFormat())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = parseConfig.GetLogger().SetLogLevel(cfg.GetLoggingConfig().GetLogLevel())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	// Get the modules
	context := modsregistrar.NewContext(cfg, database, cp, parseConfig.GetLogger())
	mods := parseConfig.GetRegistrar().BuildModules(context)
	registeredModules := modsregistrar.GetModules(mods, cfg.GetChainConfig().GetModules(), parseConfig.GetLogger())

	// Run all the additional operations
	for _, module := range registeredModules {
		if module, ok := module.(modules.AdditionalOperationsModule); ok {
			err := module.RunAdditionalOperations()
			if err != nil {
				return nil, err
			}
		}
	}

	return NewContext(cp, database, parseConfig.GetLogger(), registeredModules), nil
}
