package snapshot

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
	"github.com/panjf2000/ants/v2"
)

// GetSnapshotContext setups all the things that should be later passed to ImportSnapshot in order
// to import the chain data from the given snapshot properly.
func GetSnapshotContext(config Config) (*Context, error) {
	// Get the global config
	cfg := types.Cfg

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

	// Setup the logging
	err = config.GetLogger().SetLogFormat(cfg.GetLoggingConfig().GetLogFormat())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = config.GetLogger().SetLogLevel(cfg.GetLoggingConfig().GetLogLevel())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	pool, err := ants.NewPool(cfg.GetWorkerConfig().GetPoolSize())
	if err != nil {
		return nil, err
	}
	return NewContext(cp, database, config.GetLogger(), pool), nil
}
