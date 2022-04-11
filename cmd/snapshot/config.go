package snapshot

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
	"github.com/panjf2000/ants/v2"
)

type Config interface {
	GetConfigParser() types.ConfigParser
	GetDBBuilder() db.Builder
	GetLogger() logging.Logger
}

// Context contains the snapshot context
type Context struct {
	Proxy          client.ClientProxy
	Database       db.Database
	Logger         logging.Logger
	Pool           *ants.Pool
	BalancesBuffer chan Account
}

// NewContext builds a new Context instance
func NewContext(
	proxy client.ClientProxy, db db.Database, logger logging.Logger, pool *ants.Pool,
) *Context {
	return &Context{
		Proxy:          proxy,
		Database:       db,
		Logger:         logger,
		Pool:           pool,
		BalancesBuffer: make(chan Account),
	}
}
