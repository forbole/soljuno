package cmd

import (
	dbcmd "github.com/forbole/soljuno/cmd/db"
	initcmd "github.com/forbole/soljuno/cmd/init"
	parsecmd "github.com/forbole/soljuno/cmd/parse"
	snapshotcmd "github.com/forbole/soljuno/cmd/snapshot"
)

// Config represents the general configuration for the commands
type Config struct {
	name           string
	initConfig     *initcmd.Config
	parseConfig    *parsecmd.Config
	snapshotConfig *snapshotcmd.Config
	dbConfig       *dbcmd.Config
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
