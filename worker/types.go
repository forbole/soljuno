package worker

import (
	"github.com/forbole/soljuno/solana/parser"
	"github.com/forbole/soljuno/types/logging"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

// Context represents the context that is shared among different workers
type Context struct {
	ClientProxy client.Proxy
	Database    db.Database
	Parser      parser.Parser
	Logger      logging.Logger

	Queue   types.SlotQueue
	Modules []modules.Module
}

// NewContext allows to build a new Worker Context instance
func NewContext(
	clientProxy client.Proxy,
	db db.Database,
	parser parser.Parser,
	logger logging.Logger,
	queue types.SlotQueue,
	modules []modules.Module,
) *Context {
	return &Context{
		ClientProxy: clientProxy,
		Database:    db,
		Parser:      parser,
		Logger:      logger,

		Queue:   queue,
		Modules: modules,
	}
}
