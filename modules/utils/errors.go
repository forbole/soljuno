package utils

import (
	"github.com/forbole/soljuno/modules"
	"github.com/rs/zerolog/log"
)

// WatchMethod allows to watch for a method that returns an error.
// It executes the given method in a goroutine, logging any error that might raise.
func WatchMethod(module modules.Module, method func() error) {
	go func() {
		err := method()
		if err != nil {
			log.Error().Str("module", module.Name()).Err(err).Send()
		}
	}()
}
