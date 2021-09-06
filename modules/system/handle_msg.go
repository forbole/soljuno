package system

import "github.com/forbole/soljuno/types"

func HandleMsg(msg types.Message) error {
	switch msg.Value.Type() {
	case "createAccount":
	case "assign":
	}
	return nil
}
