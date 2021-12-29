package snapshot

import (
	"fmt"
	"os"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/db/builder"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
	"github.com/panjf2000/ants/v2"
	"github.com/spf13/cobra"
)

// ReadConfig parses the configuration file for the executable having the give name using
// the provided configuration parser
func ReadConfig(cfg *Config) types.CobraCmdFunc {
	return func(_ *cobra.Command, _ []string) error {
		file := types.GetConfigFilePath()

		// Make sure the path exists
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("config file does not exist. Make sure you have run the init command")
		}

		// Read the config
		cfg, err := types.Read(file, cfg.GetConfigParser())
		if err != nil {
			return err
		}

		// Set the global configuration
		types.Cfg = cfg
		return nil
	}
}

type Config struct {
	configParser types.ConfigParser
	buildDb      db.Builder
	logger       logging.Logger
}

// NewConfig allows to build a new Config instance
func NewConfig() *Config {
	return &Config{}
}

// WithConfigParser sets the configuration parser to be used
func (config *Config) WithConfigParser(p types.ConfigParser) *Config {
	config.configParser = p
	return config
}

// GetConfigParser returns the configuration parser to be used
func (config *Config) GetConfigParser() types.ConfigParser {
	if config.configParser == nil {
		return types.DefaultConfigParser
	}
	return config.configParser
}

// WithDBBuilder sets the database builder to be used
func (config *Config) WithDBBuilder(b db.Builder) *Config {
	config.buildDb = b
	return config
}

// GetDBBuilder returns the database builder to be used
func (config *Config) GetDBBuilder() db.Builder {
	if config.buildDb == nil {
		return builder.Builder
	}
	return config.buildDb
}

// WithLogger sets the logger to be used while importing the snapshot
func (config *Config) WithLogger(logger logging.Logger) *Config {
	config.logger = logger
	return config
}

// GetLogger returns the logger to be used when importing the snapshot
func (config *Config) GetLogger() logging.Logger {
	if config.logger == nil {
		return logging.DefaultLogger()
	}
	return config.logger
}

// --------------------------------------------------------------------------------------------------------------------

// Context contains the snapshot context
type Context struct {
	Proxy    client.ClientProxy
	Database db.Database
	Logger   logging.Logger
	Pool     *ants.Pool
	Buffer   chan Account
}

// NewContext builds a new Context instance
func NewContext(
	proxy client.ClientProxy, db db.Database, logger logging.Logger, pool *ants.Pool,
) *Context {
	return &Context{
		Proxy:    proxy,
		Database: db,
		Logger:   logger,
		Pool:     pool,
		Buffer:   make(chan Account),
	}
}
