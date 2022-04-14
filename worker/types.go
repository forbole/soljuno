package worker

import (
	"github.com/forbole/soljuno/types/logging"
	"github.com/forbole/soljuno/types/pool"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/solana/client"
	"github.com/forbole/soljuno/solana/program/parser/manager"
	"github.com/forbole/soljuno/types"
)

// Context represents the context that is shared among different workers
type Context struct {
	ClientProxy   client.ClientProxy
	Database      db.Database
	ParserManager manager.ParserManager
	Logger        logging.Logger

	Pool    pool.Pool
	Queue   types.SlotQueue
	Modules []modules.Module
}

// NewContext allows to build a new Worker Context instance
func NewContext(
	clientProxy client.ClientProxy,
	db db.Database,
	parser manager.ParserManager,
	logger logging.Logger,
	pool pool.Pool,
	queue types.SlotQueue,
	modules []modules.Module,
) *Context {
	return &Context{
		ClientProxy:   clientProxy,
		Database:      db,
		ParserManager: parser,
		Logger:        logger,

		Pool:    pool,
		Queue:   queue,
		Modules: modules,
	}
}
