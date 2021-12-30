package proxy

import (
	"fmt"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/types"
)

func GetProxyContext(config Config) (*Context, error) {
	// Get the global config
	cfg := types.Cfg

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

	return NewContext(cp, config.GetLogger()), nil
}
