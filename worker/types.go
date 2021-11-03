package worker

import (
	"github.com/panjf2000/ants/v2"

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

	Pool      *ants.Pool
	Queue     types.SlotQueue
	Modules   []modules.Module
	BankTasks types.BankTaskQueue
}

// NewContext allows to build a new Worker Context instance
func NewContext(
	clientProxy client.Proxy,
	db db.Database,
	parser parser.Parser,
	logger logging.Logger,
	pool *ants.Pool,
	queue types.SlotQueue,
	modules []modules.Module,
	bankTasks types.BankTaskQueue,
) *Context {
	return &Context{
		ClientProxy: clientProxy,
		Database:    db,
		Parser:      parser,
		Logger:      logger,

		Pool:      pool,
		Queue:     queue,
		Modules:   modules,
		BankTasks: bankTasks,
	}
}
