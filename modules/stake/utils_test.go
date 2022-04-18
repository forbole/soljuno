package stake_test

import (
	"github.com/forbole/soljuno/modules/stake"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	stakeProgram "github.com/forbole/soljuno/solana/program/stake"
)

func (suite *ModuleTestSuite) TestUpdateStakeAccount() {
	testCases := []struct {
		name      string
		isLatest  bool
		account   clienttypes.AccountInfo
		shouldErr bool
	}{
		{
			name:     "skip updating returns no error",
			isLatest: true,
		},
		{
			name:     "receive empty account data and delete account properly",
			isLatest: false,
		},
		{
			name:     "fail to decode data returns error",
			isLatest: false,
			account: clienttypes.AccountInfo{
				Value: &clienttypes.AccountValue{
					Data: [2]string{"$invalid", "base64"},
				},
			},
			shouldErr: true,
		},
		{
			name:     "receive non nonce account and delete account properly",
			isLatest: false,
			account: clienttypes.AccountInfo{
				Value: &clienttypes.AccountValue{
					Data:  [2]string{"dW5rbm93bg==", "base64"},
					Owner: "unknown",
				},
			},
		},
		{
			name:     "receive stake account and update account properly",
			isLatest: false,
			account: clienttypes.AccountInfo{
				Value: &clienttypes.AccountValue{
					Data:  [2]string{"AgAAAIDVIgAAAAAAyBJA4ron1nIF2JkcS2ipPJPqFbMjSwcV5MaNxJd+RMPIEkDiuifWcgXYmRxLaKk8k+oVsyNLBxXkxo3El35EwwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHzgaQ+C+2FHX+u9FgXvIMNGqRZ3BtgpzkaZQ1xMJ4CugJ4YpAsAAAAtAQAAAAAAAP//////////AAAAAAAA0D+PcUcGAAAAAAAAAAA=", "base64"},
					Owner: stakeProgram.ProgramID,
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			db := suite.db.GetCached()
			db.WithLatest(tc.isLatest)

			client := suite.client.GetCached()
			client.WithAccount(tc.account)

			err := stake.UpdateStakeAccount("address", 1, &db, &client)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
