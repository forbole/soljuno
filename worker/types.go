package worker

import (
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
	Logger      logging.Logger

	Queue   types.HeightQueue
	Modules []modules.Module
}

// NewContext allows to build a new Worker Context instance
func NewContext(
	clientProxy client.Proxy,
	db db.Database,
	logger logging.Logger,
	queue types.HeightQueue,
	modules []modules.Module,
) *Context {
	return &Context{
		ClientProxy: clientProxy,
		Database:    db,
		Logger:      logger,

		Queue:   queue,
		Modules: modules,
	}
}
