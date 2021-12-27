package cmd

import (
	"fmt"
	"os"

	dbcmd "github.com/forbole/soljuno/cmd/db"
	initcmd "github.com/forbole/soljuno/cmd/init"
	parsecmd "github.com/forbole/soljuno/cmd/parse"
	snapshotcmd "github.com/forbole/soljuno/cmd/snapshot"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/db/builder"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
	"github.com/spf13/cobra"
)

// Config represents the general configuration for the commands
type Config struct {
	name           string
	initConfig     *initcmd.Config
	parseConfig    *parsecmd.Config
	snapshotConfig *snapshotcmd.Config
	dbConfig       *dbcmd.Config

	configParser types.ConfigParser
	buildDb      db.Builder
	logger       logging.Logger
}

// NewConfig allows to build a new Config instance
func NewConfig(name string) *Config {
	return &Config{
		name: name,
	}
}

// GetName returns the name of the root command
func (c *Config) GetName() string {
	return c.name
}

// WithInitConfig sets cfg as the parse command configuration
func (c *Config) WithInitConfig(cfg *initcmd.Config) *Config {
	c.initConfig = cfg
	return c
}

// GetInitConfig returns the currently set parse configuration
func (c *Config) GetInitConfig() *initcmd.Config {
	if c.initConfig == nil {
		return initcmd.NewConfig()
	}
	return c.initConfig
}

// WithParseConfig sets cfg as the parse command configuration
func (c *Config) WithParseConfig(cfg *parsecmd.Config) *Config {
	c.parseConfig = cfg
	return c
}

// GetParseConfig returns the currently set parse configuration
func (c *Config) GetParseConfig() *parsecmd.Config {
	if c.parseConfig == nil {
		return parsecmd.NewConfig()
	}
	return c.parseConfig
}

// WithSnapshotConfig set cfg as the snapshot command configuration
func (c *Config) WithSnapshotConfig(cfg *snapshotcmd.Config) *Config {
	c.snapshotConfig = cfg
	return c
}

// GetSnapshotConfig returns the currently set snapshot configuration
func (c *Config) GetSnapshotConfig() *snapshotcmd.Config {
	if c.snapshotConfig == nil {
		return snapshotcmd.NewConfig()
	}
	return c.snapshotConfig
}

func (c *Config) WithDbConfig(cfg *dbcmd.Config) *Config {
	c.dbConfig = cfg
	return c
}

func (c *Config) GetDbConfig() *dbcmd.Config {
	if c.dbConfig == nil {
		return dbcmd.NewConfig()
	}
	return c.dbConfig
}

// WithConfigParser sets the configuration parser to be used
func (config *Config) WithConfigParser(p types.ConfigParser) *Config {
	config.configParser = p
	return config
}

// GetConfigParser returns the configuration parser to be used
func (c *Config) GetConfigParser() types.ConfigParser {
	if c.configParser == nil {
		return types.DefaultConfigParser
	}
	return c.configParser
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
