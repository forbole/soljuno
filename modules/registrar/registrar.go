package registrar

import (
	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/types/logging"

	"github.com/forbole/soljuno/types"

	"github.com/forbole/soljuno/modules/pruning"

	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/modules/messages"

	"github.com/forbole/soljuno/client"
	"github.com/forbole/soljuno/db"
)

// Context represents the context of the modules registrar
type Context struct {
	ParsingConfig types.Config
	Database      db.Database
	Decoder       bincode.Decoder
	Proxy         *client.Proxy
	Logger        logging.Logger
}

// NewContext allows to build a new Context instance
func NewContext(
	parsingConfig types.Config, database db.Database, proxy *client.Proxy, logger logging.Logger,
) Context {
	return Context{
		ParsingConfig: parsingConfig,
		Database:      database,
		Proxy:         proxy,
		Logger:        logger,
	}
}

// Registrar represents a modules registrar. This allows to build a list of modules that can later be used by
// specifying their names inside the TOML configuration file.
type Registrar interface {
	BuildModules(context Context) modules.Modules
}

// ------------------------------------------------------------------------------------------------------------------

var (
	_ Registrar = &EmptyRegistrar{}
)

// EmptyRegistrar represents a Registrar which does not register any custom module
type EmptyRegistrar struct{}

// BuildModules implements Registrar
func (*EmptyRegistrar) BuildModules(_ Context) modules.Modules {
	return nil
}

// ------------------------------------------------------------------------------------------------------------------

var (
	_ Registrar = &DefaultRegistrar{}
)

// DefaultRegistrar represents a registrar that allows to handle the default Juno modules
type DefaultRegistrar struct {
}

// NewDefaultRegistrar builds a new DefaultRegistrar
func NewDefaultRegistrar() *DefaultRegistrar {
	return &DefaultRegistrar{}
}

// BuildModules implements Registrar
func (r *DefaultRegistrar) BuildModules(ctx Context) modules.Modules {
	return modules.Modules{
		pruning.NewModule(ctx.ParsingConfig.GetPruningConfig(), ctx.Database, ctx.Logger),
		messages.NewModule(ctx.Decoder, ctx.Database),
	}
}

// ------------------------------------------------------------------------------------------------------------------

// GetModules returns the list of module implementations based on the given module names.
// For each module name that is specified but not found, a warning log is printed.
func GetModules(mods modules.Modules, names []string, logger logging.Logger) []modules.Module {
	var modulesImpls []modules.Module
	for _, name := range names {
		module, found := mods.FindByName(name)
		if found {
			modulesImpls = append(modulesImpls, module)
		} else {
			logger.Error("Module is required but not registered. Be sure to register it using registrar.RegisterModule", "module", name)
		}
	}
	return modulesImpls
}
