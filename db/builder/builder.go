package builder

import (
	"github.com/forbole/soljuno/db"

	"github.com/forbole/soljuno/db/postgresql"
)

// Builder represents a generic Builder implementation that build the proper database
// instance based on the configuration the user has specified
func Builder(ctx *db.Context) (db.Database, error) {
	return postgresql.Builder(ctx)
}
