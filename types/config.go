package types

import (
	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog"
)

var (
	// Cfg represents the configuration to be used during the execution
	Cfg Config
)

// ConfigParser represents a function that allows to parse a file contents as a Config object
type ConfigParser = func(fileContents []byte) (Config, error)

type configToml struct {
	RPC       *rpcConfig       `toml:"rpc"`
	Chain     *chainConfig     `toml:"chain"`
	Database  *databaseConfig  `toml:"database"`
	Logging   *loggingConfig   `toml:"logging"`
	Parsing   *parsingConfig   `toml:"parsing"`
	Pruning   *pruningConfig   `toml:"pruning"`
	Telemetry *telemetryConfig `toml:"telemetry"`
	Worker    *workerConfig    `toml:"worker"`
}

// DefaultConfigParser attempts to read and parse a Juno config from the given string bytes.
// An error reading or parsing the config results in a panic.
func DefaultConfigParser(configData []byte) (Config, error) {
	var cfg configToml
	err := toml.Unmarshal(configData, &cfg)
	return NewConfig(
		cfg.RPC,
		cfg.Chain,
		cfg.Database,
		cfg.Logging,
		cfg.Parsing,
		cfg.Pruning,
		cfg.Telemetry,
		cfg.Worker,
	), err
}

// ---------------------------------------------------------------------------------------------------------------------

// Config represents the configuration to run Juno
type Config interface {
	GetRPCConfig() RPCConfig
	GetChainConfig() ChainConfig
	GetDatabaseConfig() DatabaseConfig
	GetLoggingConfig() LoggingConfig
	GetParsingConfig() ParsingConfig
	GetPruningConfig() PruningConfig
	GetTelemetryConfig() TelemetryConfig
	GetWorkerConfig() WorkerConfig
}

var _ Config = &config{}

// Config defines all necessary juno configuration parameters.
type config struct {
	RPC       RPCConfig       `toml:"rpc"`
	Chain     ChainConfig     `toml:"chain"`
	Database  DatabaseConfig  `toml:"database"`
	Logging   LoggingConfig   `toml:"logging"`
	Parsing   ParsingConfig   `toml:"parsing"`
	Pruning   PruningConfig   `toml:"pruning"`
	Telemetry TelemetryConfig `toml:"telemetry"`
	Worker    WorkerConfig    `toml:"worker"`
}

// NewConfig builds a new Config instance
func NewConfig(
	rpcConfig RPCConfig,
	chainConfig ChainConfig,
	dbConfig DatabaseConfig,
	loggingConfig LoggingConfig,
	parsingConfig ParsingConfig,
	pruningConfig PruningConfig,
	telemetryConfig TelemetryConfig,
	workerConfig WorkerConfig,
) Config {
	return &config{
		RPC:       rpcConfig,
		Chain:     chainConfig,
		Database:  dbConfig,
		Logging:   loggingConfig,
		Parsing:   parsingConfig,
		Pruning:   pruningConfig,
		Telemetry: telemetryConfig,
		Worker:    workerConfig,
	}
}

// GetRPCConfig implements Config
func (c *config) GetRPCConfig() RPCConfig {
	if c.RPC == nil {
		return DefaultRPCConfig()
	}
	return c.RPC
}

// GetChainConfig implements Config
func (c *config) GetChainConfig() ChainConfig {
	if c.Chain == nil {
		return DefaultChainConfig()
	}
	return c.Chain
}

// GetDatabaseConfig implements Config
func (c *config) GetDatabaseConfig() DatabaseConfig {
	if c.Database == nil {
		return DefaultDatabaseConfig()
	}
	return c.Database
}

// GetLoggingConfig implements Config
func (c *config) GetLoggingConfig() LoggingConfig {
	if c.Logging == nil {
		return DefaultLoggingConfig()
	}
	return c.Logging
}

// GetParsingConfig implements Config
func (c *config) GetParsingConfig() ParsingConfig {
	if c.Parsing == nil {
		return DefaultParsingConfig()
	}
	return c.Parsing
}

// GetPruningConfig implements Config
func (c *config) GetPruningConfig() PruningConfig {
	if c.Pruning == nil {
		return DefaultPruningConfig()
	}
	return c.Pruning
}

// GetTelemetryConfig implements Config
func (c *config) GetTelemetryConfig() TelemetryConfig {
	if c.Telemetry == nil {
		return DefaultTelemetryConfig()
	}
	return c.Telemetry
}

// GetWorkerConfig implements Config
func (c *config) GetWorkerConfig() WorkerConfig {
	if c.Worker == nil {
		return DefaultWorkerConfig()
	}
	return c.Worker
}

// ---------------------------------------------------------------------------------------------------------------------

// RPCConfig contains the configuration of the RPC endpoint
type RPCConfig interface {
	GetClientName() string
	GetAddress() string
}

var _ RPCConfig = &rpcConfig{}

type rpcConfig struct {
	ClientName string `toml:"client_name"`
	Address    string `toml:"address"`
}

// NewRPCConfig allows to build a new RPCConfig instance
func NewRPCConfig(clientName, address string) RPCConfig {
	return &rpcConfig{
		ClientName: clientName,
		Address:    address,
	}
}

// DefaultRPCConfig returns the default instance of RPCConfig
func DefaultRPCConfig() RPCConfig {
	return NewRPCConfig("soljuno", "http://localhost:8899")
}

// GetClientName implements RPCConfig
func (r *rpcConfig) GetClientName() string {
	return r.ClientName
}

// GetAddress implements RPCConfig
func (r *rpcConfig) GetAddress() string {
	return r.Address
}

// ---------------------------------------------------------------------------------------------------------------------

// ChainConfig contains the data to configure the ChainConfig SDK
type ChainConfig interface {
	GetPrefix() string
	GetModules() []string
}

var _ ChainConfig = &chainConfig{}

type chainConfig struct {
	Prefix  string   `toml:"prefix"`
	Modules []string `toml:"modules"`
}

// NewChainConfig returns a new ChainConfig instance
func NewChainConfig(prefix string, modules []string) ChainConfig {
	return &chainConfig{
		Prefix:  prefix,
		Modules: modules,
	}
}

// DefaultChainConfig returns the default instance of ChainConfig
func DefaultChainConfig() ChainConfig {
	return NewChainConfig("", nil)
}

// GetPrefix implements ChainConfig
func (c *chainConfig) GetPrefix() string {
	return c.Prefix
}

// GetModules implements ChainConfig
func (c *chainConfig) GetModules() []string {
	return c.Modules
}

// ---------------------------------------------------------------------------------------------------------------------

// DatabaseConfig represents a generic database configuration
type DatabaseConfig interface {
	GetName() string
	GetHost() string
	GetPort() int64
	GetUser() string
	GetPassword() string
	GetSSLMode() string
	GetSchema() string
	GetMaxOpenConnections() int
	GetMaxIdleConnections() int
}

var _ DatabaseConfig = &databaseConfig{}

type databaseConfig struct {
	Name               string `toml:"name"`
	Host               string `toml:"host"`
	Port               int64  `toml:"port"`
	User               string `toml:"user"`
	Password           string `toml:"password"`
	SSLMode            string `toml:"ssl_mode"`
	Schema             string `toml:"schema"`
	MaxOpenConnections int    `toml:"max_open_connections"`
	MaxIdleConnections int    `toml:"max_idle_connections"`
}

func NewDatabaseConfig(
	name, host string, port int64, user string, password string,
	sslMode string, schema string,
	maxOpenConnections int, maxIdleConnections int,
) DatabaseConfig {
	return &databaseConfig{
		Name:               name,
		Host:               host,
		Port:               port,
		User:               user,
		Password:           password,
		SSLMode:            sslMode,
		Schema:             schema,
		MaxOpenConnections: maxOpenConnections,
		MaxIdleConnections: maxIdleConnections,
	}
}

// DefaultDatabaseConfig returns the default instance of DatabaseConfig
func DefaultDatabaseConfig() DatabaseConfig {
	return NewDatabaseConfig(
		"database-name",
		"localhost",
		5432,
		"user",
		"password",
		"",
		"public",
		1,
		1,
	)
}

// GetName implements DatabaseConfig
func (d *databaseConfig) GetName() string {
	return d.Name
}

// GetHost implements DatabaseConfig
func (d *databaseConfig) GetHost() string {
	return d.Host
}

// GetPort implements DatabaseConfig
func (d *databaseConfig) GetPort() int64 {
	return d.Port
}

// GetUser implements DatabaseConfig
func (d *databaseConfig) GetUser() string {
	return d.User
}

// GetPassword implements DatabaseConfig
func (d *databaseConfig) GetPassword() string {
	return d.Password
}

// GetSSLMode implements DatabaseConfig
func (d *databaseConfig) GetSSLMode() string {
	return d.SSLMode
}

// GetSchema implements DatabaseConfig
func (d *databaseConfig) GetSchema() string {
	return d.Schema
}

// GetMaxOpenConnections implements DatabaseConfig
func (d *databaseConfig) GetMaxOpenConnections() int {
	return d.MaxOpenConnections
}

// GetMaxIdleConnections implements DatabaseConfig
func (d *databaseConfig) GetMaxIdleConnections() int {
	return d.MaxIdleConnections
}

// ---------------------------------------------------------------------------------------------------------------------

// LoggingConfig represents the configuration for the logging part
type LoggingConfig interface {
	GetLogLevel() string
	GetLogFormat() string
}

var _ LoggingConfig = &loggingConfig{}

type loggingConfig struct {
	LogLevel  string `toml:"level"`
	LogFormat string `toml:"format"`
}

// NewLoggingConfig returns a new LoggingConfig instance
func NewLoggingConfig(level, format string) LoggingConfig {
	return &loggingConfig{
		LogLevel:  level,
		LogFormat: format,
	}
}

// DefaultLoggingConfig returns the default LoggingConfig instance
func DefaultLoggingConfig() LoggingConfig {
	return NewLoggingConfig(zerolog.DebugLevel.String(), "text")
}

// GetLogLevel implements LoggingConfig
func (l *loggingConfig) GetLogLevel() string {
	return l.LogLevel
}

// GetLogFormat implements LoggingConfig
func (l *loggingConfig) GetLogFormat() string {
	return l.LogFormat
}

// ---------------------------------------------------------------------------------------------------------------------

// ParsingConfig represents the configuration of the parsing
type ParsingConfig interface {
	GetWorkers() int64
	ShouldParseNewBlocks() bool
	ShouldParseOldBlocks() bool
	ShouldParseGenesis() bool
	GetGenesisFilePath() string
	GetStartSlot() uint64
	UseFastSync() bool
}

var _ ParsingConfig = &parsingConfig{}

type parsingConfig struct {
	Workers         int64  `toml:"workers"`
	ParseNewBlocks  bool   `toml:"listen_new_blocks"`
	ParseOldBlocks  bool   `toml:"parse_old_blocks"`
	GenesisFilePath string `toml:"genesis_file_path"`
	ParseGenesis    bool   `toml:"parse_genesis"`
	StartSlot       uint64 `toml:"start_slot"`
	FastSync        bool   `toml:"fast_sync"`
}

// NewParsingConfig allows to build a new ParsingConfig instance
func NewParsingConfig(
	workers int64,
	parseNewBlocks, parseOldBlocks bool,
	parseGenesis bool, genesisFilePath string, startSlot uint64, fastSync bool,
) ParsingConfig {
	return &parsingConfig{
		Workers:         workers,
		ParseOldBlocks:  parseOldBlocks,
		ParseNewBlocks:  parseNewBlocks,
		ParseGenesis:    parseGenesis,
		GenesisFilePath: genesisFilePath,
		StartSlot:       startSlot,
		FastSync:        fastSync,
	}
}

// DefaultParsingConfig returns the default instance of ParsingConfig
func DefaultParsingConfig() ParsingConfig {
	return NewParsingConfig(
		1,
		true,
		true,
		true,
		"",
		1,
		false,
	)
}

// GetWorkers implements ParsingConfig
func (p *parsingConfig) GetWorkers() int64 {
	return p.Workers
}

// ShouldParseNewBlocks implements ParsingConfig
func (p *parsingConfig) ShouldParseNewBlocks() bool {
	return p.ParseNewBlocks
}

// ShouldParseOldBlocks implements ParsingConfig
func (p *parsingConfig) ShouldParseOldBlocks() bool {
	return p.ParseOldBlocks
}

// ShouldParseGenesis implements ParsingConfig
func (p *parsingConfig) ShouldParseGenesis() bool {
	return p.ParseGenesis
}

func (p *parsingConfig) GetGenesisFilePath() string {
	return p.GenesisFilePath
}

// GetStartHeight implements ParsingConfig
func (p *parsingConfig) GetStartSlot() uint64 {
	return p.StartSlot
}

// UseFastSync implements ParsingConfig
func (p *parsingConfig) UseFastSync() bool {
	return p.FastSync
}

// ---------------------------------------------------------------------------------------------------------------------

// PruningConfig contains the configuration of the pruning strategy
type PruningConfig interface {
	GetKeepRecent() int64
	GetKeepEvery() int64
	GetInterval() int64
}

var _ PruningConfig = &pruningConfig{}

type pruningConfig struct {
	KeepRecent int64 `toml:"keep_recent"`
	KeepEvery  int64 `toml:"keep_every"`
	Interval   int64 `toml:"interval"`
}

// NewPruningConfig returns a new PruningConfig
func NewPruningConfig(keepRecent, keepEvery, interval int64) PruningConfig {
	return &pruningConfig{
		KeepRecent: keepRecent,
		KeepEvery:  keepEvery,
		Interval:   interval,
	}
}

// DefaultPruningConfig returns the default PruningConfig instance
func DefaultPruningConfig() PruningConfig {
	return NewPruningConfig(100, 500, 10)
}

// GetKeepRecent implements PruningConfig
func (p *pruningConfig) GetKeepRecent() int64 {
	return p.KeepRecent
}

// GetKeepEvery implements PruningConfig
func (p *pruningConfig) GetKeepEvery() int64 {
	return p.KeepEvery
}

// GetInterval implements PruningConfig
func (p *pruningConfig) GetInterval() int64 {
	return p.Interval
}

// ---------------------------------------------------------------------------------------------------------------------

// TelemetryConfig contains the configuration of the telemetry strategy
type TelemetryConfig interface {
	IsEnabled() bool
	GetPort() uint
}

var _ TelemetryConfig = &telemetryConfig{}

type telemetryConfig struct {
	Enabled bool `toml:"enabled"`
	Port    uint `toml:"port"`
}

// NewTelemetryConfig allows to build a new TelemetryConfig instance
func NewTelemetryConfig(enabled bool, port uint) TelemetryConfig {
	return &telemetryConfig{
		Enabled: enabled,
		Port:    port,
	}
}

// DefaultTelemetryConfig returns the default TelemetryConfig instance
func DefaultTelemetryConfig() TelemetryConfig {
	return NewTelemetryConfig(false, 500)
}

// IsEnabled implements TelemetryConfig
func (p *telemetryConfig) IsEnabled() bool {
	return p.Enabled
}

// GetPort implements TelemetryConfig
func (p *telemetryConfig) GetPort() uint {
	return p.Port
}

// ---------------------------------------------------------------------------------------------------------------------

// WorkerConfig contains the configuration of the worker strategy
type WorkerConfig interface {
	GetPoolSize() int
}

var _ WorkerConfig = &workerConfig{}

type workerConfig struct {
	PoolSize int `toml:"poll_size"`
}

// NewWorkerConfig allows to build a new WorkerConfig instance
func NewWorkerConfig(poolSize int) WorkerConfig {
	return &workerConfig{
		PoolSize: poolSize,
	}
}

// DefaultWorkerConfig returns the default WorkerConfig instance
func DefaultWorkerConfig() WorkerConfig {
	return NewWorkerConfig(5000)
}

// GetPoolSize implements WorkerConfig
func (w *workerConfig) GetPoolSize() int {
	return w.PoolSize
}
