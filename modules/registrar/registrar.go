package registrar

import (
	"github.com/forbole/soljuno/solana/parser/manager"
	"github.com/forbole/soljuno/types/logging"
	"github.com/forbole/soljuno/types/pool"

	"github.com/forbole/soljuno/types"

	"github.com/forbole/soljuno/modules/bank"
	"github.com/forbole/soljuno/modules/bpfloader"
	"github.com/forbole/soljuno/modules/config"
	"github.com/forbole/soljuno/modules/consensus"
	"github.com/forbole/soljuno/modules/epoch"
	"github.com/forbole/soljuno/modules/history"
	"github.com/forbole/soljuno/modules/instructions"
	"github.com/forbole/soljuno/modules/pricefeed"
	"github.com/forbole/soljuno/modules/pruning"
	"github.com/forbole/soljuno/modules/stake"
	"github.com/forbole/soljuno/modules/system"
	"github.com/forbole/soljuno/modules/telemetry"
	"github.com/forbole/soljuno/modules/token"
	"github.com/forbole/soljuno/modules/txs"
	"github.com/forbole/soljuno/modules/vote"

	"github.com/forbole/soljuno/modules"

	"github.com/forbole/soljuno/db"
	"github.com/forbole/soljuno/solana/client"
)

// Context represents the context of the modules registrar
type Context struct {
	Config        types.Config
	Database      db.Database
	ParserManager manager.ParserManager
	Proxy         client.ClientProxy
	Logger        logging.Logger
	Pool          pool.Pool
}

// NewContext allows to build a new Context instance
func NewContext(
	config types.Config, database db.Database, proxy client.ClientProxy, logger logging.Logger, pool pool.Pool,
) Context {
	return Context{
		Config:   config,
		Database: database,
		Proxy:    proxy,
		Logger:   logger,
		Pool:     pool,
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
	pruningModule := pruning.NewModule(ctx.Config.GetPruningConfig(), ctx.Database, ctx.Logger)
	txsModule := txs.NewModule(ctx.Database, ctx.Pool)
	instructionsModule := instructions.NewModule(ctx.Database, ctx.Pool)
	pruningModule.RegisterService(txsModule, instructionsModule)

	epochModule := epoch.NewModule(ctx.Database, ctx.Proxy)
	voteModule := vote.NewModule(ctx.Database, ctx.Proxy)
	epochModule.RegisterService(voteModule)

	historyModule := history.NewModule()
	pricefeedModule := pricefeed.NewModule(ctx.Database)
	bankModule := bank.NewModule(ctx.Database)
	historyModule.RegisterService(pricefeedModule, bankModule)

	return modules.Modules{
		telemetry.NewModule(ctx.Config),
		bankModule,
		system.NewModule(ctx.Database, ctx.Proxy),
		stake.NewModule(ctx.Database, ctx.Proxy),
		token.NewModule(ctx.Database, ctx.Proxy),
		config.NewModule(ctx.Database, ctx.Proxy),
		bpfloader.NewModule(ctx.Database, ctx.Proxy),
		consensus.NewModule(ctx.Database),
		pricefeedModule,
		pruningModule,
		txsModule,
		instructionsModule,
		epochModule,
		voteModule,
		historyModule,
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
