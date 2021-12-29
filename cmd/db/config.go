package db

import (
	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/types"
	"github.com/forbole/soljuno/types/logging"
)

type Config interface {
	GetConfigParser() types.ConfigParser
	GetDBBuilder() db.Builder
	GetLogger() logging.Logger
}

// --------------------------------------------------------------------------------------------------------------------

// Context contains the snapshot context
type Context struct {
	Database db.Database
	Logger   logging.Logger
}

// NewContext builds a new Context instance
func NewContext(
	db db.Database, logger logging.Logger,
) *Context {
	return &Context{
		Database: db,
		Logger:   logger,
	}
}
