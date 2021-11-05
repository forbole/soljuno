package snapshot

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
)

// GetParsingContext setups all the things that should be later passed to ImportSnapshot in order
// to import the chain data from the given snapshot properly.
func GetParsingContext(parseConfig *Config) (*Context, error) {
	// Get the global config
	cfg := types.Cfg

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

	return NewContext(cp, database, parseConfig.GetLogger()), nil
}
