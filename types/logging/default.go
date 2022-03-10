package logging

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

var (
	_ Logger = &defaultLogger{}
)

// defaultLogger represents the default logger for any kind of error
type defaultLogger struct {
	Logger zerolog.Logger
}

// DefaultLogger allows to build a new defaultLogger instance
func DefaultLogger() Logger {
	return &defaultLogger{
		Logger: log.Logger,
	}
}

// SetLogLevel implements Logger
func (d *defaultLogger) SetLogLevel(level string) error {
	logLvl, err := zerolog.ParseLevel(level)
	if err != nil {
		return err
	}

	zerolog.SetGlobalLevel(logLvl)
	return nil
}

// SetLogFormat implements Logger
func (d *defaultLogger) SetLogFormat(format string) error {
	switch format {
	case "json":
		// JSON is the default logging format
		break

	case "text":
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	default:
		return fmt.Errorf("invalid logging format: %s", format)
	}

	return nil
}

// Info implements Logger
func (d *defaultLogger) Info(msg string, keyVals ...interface{}) {
	d.Logger.Info().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Debug implements Logger
func (d *defaultLogger) Debug(msg string, keyVals ...interface{}) {
	d.Logger.Debug().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Error implements Logger
func (d *defaultLogger) Error(msg string, keyVals ...interface{}) {
	ErrorCount.Inc()
	d.Logger.Error().Fields(getLogFields(keyVals...)).Msg(msg)
}

// GenesisError implements Logger
func (d *defaultLogger) GenesisError(module modules.Module, err error) {
	d.Error("error while handling genesis",
		"err", err,
		LogKeyModule, module.Name(),
	)
}

// BlockError implements Logger
func (d *defaultLogger) BlockError(module modules.Module, block types.Block, err error) {
	d.Error("error while handling block",
		"err", err,
		LogKeyModule, module.Name(),
		LogKeySlot, block.Slot,
	)
}

// TxError implements Logger
func (d *defaultLogger) TxError(module modules.Module, tx types.Tx, err error) {
	d.Error("error while handling transaction",
		"err", err,
		LogKeyModule, module.Name(),
		LogKeySlot, tx.Slot,
		LogKeyTxHash, tx.Hash,
	)
}

// InstructionError implements Logger
func (d *defaultLogger) InstructionError(module modules.Module, tx types.Tx, instruction types.Instruction, err error) {
	d.Error("error while handling instruction",
		"err", err,
		LogKeyModule, module.Name(),
		LogKeySlot, tx.Slot,
		LogKeyTxHash, tx.Hash,
		LogKeyProgram, instruction.Program,
		LogKeyInstructionType, instruction.Parsed.Type,
	)
}

func getLogFields(keyVals ...interface{}) map[string]interface{} {
	if len(keyVals)%2 != 0 {
		return nil
	}

	fields := make(map[string]interface{})
	for i := 0; i < len(keyVals); i += 2 {
		fields[keyVals[i].(string)] = keyVals[i+1]
	}

	return fields
}
