package parse

import (
	"github.com/forbole/soljuno/types/logging"
	"github.com/panjf2000/ants/v2"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/db/builder"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/modules/registrar"
	"github.com/forbole/soljuno/types"
)

// Config contains all the configuration for the "parse" command
type Config struct {
	registrar    registrar.Registrar
	configParser types.ConfigParser
	buildDb      db.Builder
	logger       logging.Logger
}

// NewConfig allows to build a new Config instance
func NewConfig() *Config {
	return &Config{}
}

// WithRegistrar sets the modules registrar to be used
func (config *Config) WithRegistrar(r registrar.Registrar) *Config {
	config.registrar = r
	return config
}

// GetRegistrar returns the modules registrar to be used
func (config *Config) GetRegistrar() registrar.Registrar {
	if config.registrar == nil {
		return &registrar.EmptyRegistrar{}
	}
	return config.registrar
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

// WithLogger sets the logger to be used while parsing the data
func (config *Config) WithLogger(logger logging.Logger) *Config {
	config.logger = logger
	return config
}

// GetLogger returns the logger to be used when parsing the data
func (config *Config) GetLogger() logging.Logger {
	if config.logger == nil {
		return logging.DefaultLogger()
	}
	return config.logger
}

// --------------------------------------------------------------------------------------------------------------------

// Context contains the parsing context
type Context struct {
	Proxy    client.Proxy
	Database db.Database
	Logger   logging.Logger
	Modules  []modules.Module
	Pool     *ants.Pool
}

// NewContext builds a new Context instance
func NewContext(
	proxy client.Proxy, db db.Database,
	logger logging.Logger, modules []modules.Module,
	pool *ants.Pool,
) *Context {
	return &Context{
		Proxy:    proxy,
		Database: db,
		Modules:  modules,
		Logger:   logger,
		Pool:     pool,
	}
}
