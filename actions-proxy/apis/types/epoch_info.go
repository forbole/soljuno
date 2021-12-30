package types

import (
	"fmt"

	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

type EpochInfo struct {
	AbsoluteSlot     string `json:"absoluteSlot"`
	BlockHeight      string `json:"blockHeight"`
	Epoch            string `json:"epoch"`
	SlotIndex        string `json:"slotIndex"`
	SlotsInEpoch     string `json:"slotsInEpoch"`
	TransactionCount string `json:"transactionCount"`
}

func NewEpochInfo(info clienttypes.EpochInfo) EpochInfo {
	return EpochInfo{
		AbsoluteSlot:     fmt.Sprintf("%d", info.AbsoluteSlot),
		BlockHeight:      fmt.Sprintf("%d", info.BlockHeight),
		Epoch:            fmt.Sprintf("%d", info.Epoch),
		SlotIndex:        fmt.Sprintf("%d", info.SlotIndex),
		SlotsInEpoch:     fmt.Sprintf("%d", info.SlotsInEpoch),
		TransactionCount: fmt.Sprintf("%d", info.TransactionCount),
	}
}
