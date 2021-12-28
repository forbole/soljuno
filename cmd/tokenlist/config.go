package tokenlist

import (
	"github.com/forbole/soljuno/types/logging"

	"github.com/forbole/soljuno/db"
)

// Config contains all the configuration for the "parse" command
type Config interface {
	GetDBBuilder() db.Builder
	GetLogger() logging.Logger
}

// --------------------------------------------------------------------------------------------------------------------

// Context contains the parsing context
type Context struct {
	Database db.Database
	Logger   logging.Logger
}

// NewContext builds a new Context instance
func NewContext(
	db db.Database,
	logger logging.Logger,
) *Context {
	return &Context{
		Database: db,
		Logger:   logger,
	}
}
