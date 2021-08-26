package messages

import (
	"fmt"

	"github.com/forbole/soljuno/solana/bincode"
	"github.com/forbole/soljuno/types"
)

// MessageNotSupported returns an error telling that the given message is not supported
func MessageNotSupported(msg types.Instruction) error {
	return fmt.Errorf("message type not supported: %s", msg.Type)
}

// MessageParser represents a function that extracts all the
// involved addresses from a provided message (both accounts and validators)
type MessageParser = func(cdc bincode.Decoder, msg types.Instruction) ([]string, error)

// JoinMessageParsers joins together all the given parsers, calling them in order
func JoinMessageParsers(parsers ...MessageParser) MessageParser {
	return func(cdc bincode.Decoder, msg types.Instruction) ([]string, error) {
		for _, parser := range parsers {
			// Try getting the addresses
			addresses, _ := parser(cdc, msg)

			// If some addresses are found, return them
			if len(addresses) > 0 {
				return addresses, nil
			}
		}
		return nil, MessageNotSupported(msg)
	}
}
