package proxy

import (
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
)

type Config interface {
	GetConfigParser() types.ConfigParser
	GetLogger() logging.Logger
}

// --------------------------------------------------------------------------------------------------------------------

// Context contains the snapshot context
type Context struct {
	Proxy  client.ClientProxy
	Logger logging.Logger
	Port   int
}

// NewContext builds a new Context instance
func NewContext(
	proxy client.ClientProxy,
	logger logging.Logger,
) *Context {
	return &Context{
		Proxy:  proxy,
		Logger: logger,
	}
}

func (c *Context) WithPort(port int) *Context {
	c.Port = port
	return c
}
