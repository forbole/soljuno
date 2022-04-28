package bank_test

import (
	"time"

	"github.com/forbole/soljuno/modules/bank"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/types"
)

func (suite *ModuleTestSuite) TestModule_Name() {
	suite.Require().Equal("bank", suite.module.Name())
}

func (suite *ModuleTestSuite) TestModule_HandleBlock() {
	block := types.NewBlock(
		0,
		0,
		"hash",
		"leader",
		[]clienttypes.Reward{clienttypes.NewReward("address", 1, 1, clienttypes.RewardStaking, 0)},
		time.Date(2022, 04, 14, 0, 0, 0, 0, time.UTC),
		[]types.Tx{
			types.NewTx(
				"sig",
				0,
				0,
				nil,
				0,
				nil,
				nil,
				[]string{"address"},
				[]uint64{2},
				[]clienttypes.TransactionTokenBalance{
					{
						AccountIndex:  0,
						Mint:          "mint",
						UiTokenAmount: clienttypes.UiTokenAmount{Amount: "1"},
					}},
			),
			types.NewTx(
				"sig",
				0,
				1,
				nil,
				0,
				nil,
				nil,
				[]string{"address"},
				[]uint64{3},
				[]clienttypes.TransactionTokenBalance{
					{
						AccountIndex:  0,
						Mint:          "mint",
						UiTokenAmount: clienttypes.UiTokenAmount{Amount: "2"},
					}},
			),
		},
	)
	err := suite.module.HandleBlock(block)
	suite.Require().NoError(err)
	suite.Require().Equal(bank.NewAccountBalanceEntries(0, []string{"address"}, []uint64{3}), suite.module.BalanceEntries)
	suite.Require().Equal(bank.NewAccountBalanceEntries(0, []string{"address"}, []uint64{3}), suite.module.HistoryBalanceEntries)

	suite.Require().Equal(
		bank.NewTokenAccountBalanceEntries(
			0,
			[]string{"address"},
			[]clienttypes.TransactionTokenBalance{
				{
					AccountIndex:  0,
					Mint:          "mint",
					UiTokenAmount: clienttypes.UiTokenAmount{Amount: "2"},
				},
			},
		), suite.module.TokenBalanceEntries)
	suite.Require().Equal(
		bank.NewTokenAccountBalanceEntries(
			0,
			[]string{"address"},
			[]clienttypes.TransactionTokenBalance{
				{
					AccountIndex:  0,
					Mint:          "mint",
					UiTokenAmount: clienttypes.UiTokenAmount{Amount: "2"},
				},
			},
		), suite.module.HistoryTokenBalanceEntries)
}
