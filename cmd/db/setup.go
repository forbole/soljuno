package db

import (
	"fmt"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// GetDatabaseContext setups all the things that should be later passed to InitDatabase in order
// to import the chain data from the given snapshot properly.
func GetDatabaseContext(snapshotConfig *Config) (*Context, error) {
	// Get the global config
	cfg := types.Cfg

	databaseCtx := db.NewContext(cfg.GetDatabaseConfig(), snapshotConfig.GetLogger())
	database, err := snapshotConfig.GetDBBuilder()(databaseCtx)
	if err != nil {
		return nil, err
	}

	// Setup the logging
	err = snapshotConfig.GetLogger().SetLogFormat(cfg.GetLoggingConfig().GetLogFormat())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging format: %s", err)
	}

	err = snapshotConfig.GetLogger().SetLogLevel(cfg.GetLoggingConfig().GetLogLevel())
	if err != nil {
		return nil, fmt.Errorf("error while setting logging level: %s", err)
	}

	return NewContext(database, snapshotConfig.GetLogger()), nil
}
