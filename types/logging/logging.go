package logging

import (
	"github.com/forbole/soljuno/modules"
	"github.com/forbole/soljuno/types"
)

const (
	LogKeyModule  = "module"
	LogKeySlot    = "slot"
	LogKeyTxHash  = "tx_hash"
	LogKeyProgram = "program"
	LogKeyMsgType = "msg_type"
)

// Logger defines a function that takes an error and logs it.
type Logger interface {
	SetLogLevel(level string) error
	SetLogFormat(format string) error

	Info(msg string, keyvals ...interface{})
	Debug(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})

	GenesisError(module modules.Module, err error)
	BlockError(module modules.Module, block types.Block, err error)
	TxError(module modules.Module, tx types.Tx, err error)
	MsgError(module modules.Module, tx types.Tx, msg types.Message, err error)
}
