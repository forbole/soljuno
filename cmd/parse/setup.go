package parse

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types/pool"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/modules"
	modsregistrar "github.com/forbole/soljuno/modules/registrar"
	"github.com/forbole/soljuno/types"
)

// GetParsingContext setups all the things that should be later passed to StartParsing in order
// to parse the chain data properly.
func GetParsingContext(config Config) (*Context, error) {
	// Get the global config
	cfg := types.Cfg

	// Get the database
	databaseCtx := db.NewContext(cfg.GetDatabaseConfig(), config.GetLogger())
	database, err := config.GetDBBuilder()(databaseCtx)
	if err != nil {
		return nil, err
	}

	// Init the client
	cp, err := client.NewClientProxy(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to start client: %s", err)
	}

	pool, err := pool.NewDefaultPool(cfg.GetWorkerConfig().GetPoolSize())
	if err != nil {
		return nil, err
	}

	// Setup the logging
	err = config.GetLogger().SetLogFormat(cfg.GetLoggingConfig().GetLogFormat())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = config.GetLogger().SetLogLevel(cfg.GetLoggingConfig().GetLogLevel())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	// Create a queue that will collect, aggregate, and export blocks and metadata
	slotQueue := types.NewQueue(25)

	// Get the modules
	context := modsregistrar.NewContext(cfg, database, cp, config.GetLogger(), pool, slotQueue)
	mods := config.GetRegistrar().BuildModules(context)
	registeredModules := modsregistrar.GetModules(mods, cfg.GetChainConfig().GetModules(), config.GetLogger())

	// Run all the additional operations
	for _, module := range registeredModules {
		if module, ok := module.(modules.AdditionalOperationsModule); ok {
			err := module.RunAdditionalOperations()
			if err != nil {
				return nil, err
			}
		}
	}

	return NewContext(cfg, cp, database, config.GetLogger(), registeredModules, pool, slotQueue), nil
}
