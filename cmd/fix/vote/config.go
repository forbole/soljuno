package vote

import (
	"github.com/forbole/soljuno/types/logging"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/types"
)

// Config contains all the configuration for the "parse" command
type Config interface {
	GetConfigParser() types.ConfigParser
	GetDBBuilder() db.Builder
	GetLogger() logging.Logger
}

// --------------------------------------------------------------------------------------------------------------------

// Context contains the parsing context
type Context struct {
	GlobalCfg types.Config
	Proxy     client.ClientProxy
	Database  db.Database
	Logger    logging.Logger
}

// NewContext builds a new Context instance
func NewContext(
	globalCfg types.Config, proxy client.ClientProxy, db db.Database, logger logging.Logger,
) *Context {
	return &Context{
		GlobalCfg: globalCfg,
		Proxy:     proxy,
		Database:  db,
		Logger:    logger,
	}
}
