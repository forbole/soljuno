package parse

import (
	"github.com/forbole/soljuno/types/logging"
	"github.com/forbole/soljuno/types/pool"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/modules/registrar"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/types"
)

// Config contains all the configuration for the "parse" command
type Config interface {
	GetRegistrar() registrar.Registrar
	GetConfigParser() types.ConfigParser
	GetDBBuilder() db.Builder
	GetLogger() logging.Logger
}

// --------------------------------------------------------------------------------------------------------------------

// Context contains the parsing context
type Context struct {
	Proxy    client.ClientProxy
	Database db.Database
	Logger   logging.Logger
	Modules  []modules.Module
	Pool     pool.Pool
}

// NewContext builds a new Context instance
func NewContext(
	proxy client.ClientProxy, db db.Database,
	logger logging.Logger, modules []modules.Module,
	pool pool.Pool,
) *Context {
	return &Context{
		Proxy:    proxy,
		Database: db,
		Modules:  modules,
		Logger:   logger,
		Pool:     pool,
	}
}
