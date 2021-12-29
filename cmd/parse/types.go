package parse

import (
	"github.com/forbole/soljuno/types/logging"
	"github.com/panjf2000/ants/v2"

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
	Pool     *ants.Pool
}

// NewContext builds a new Context instance
func NewContext(
	proxy client.ClientProxy, db db.Database,
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
