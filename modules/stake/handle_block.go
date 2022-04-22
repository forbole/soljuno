package stake

import (
	"github.com/forbole/soljuno/db"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/types"
)

func HandleBlock(block types.Block, db db.StakeDb, client ClientProxy) error {
	for _, reward := range block.Rewards {
		if reward.RewardType != clienttypes.RewardStaking {
			continue
		}
		err := UpdateStakeAccount(reward.Pubkey, block.Slot, db, client)
		if err != nil {
			return err
		}
	}
	return nil
}
