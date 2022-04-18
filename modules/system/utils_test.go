package system_test

import (
	"github.com/forbole/soljuno/modules/system"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
)

func (suite *ModuleTestSuite) TestUpdateNonceAccount() {
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
			name:     "receive nonce account and update account properly",
			isLatest: false,
			account: clienttypes.AccountInfo{
				Value: &clienttypes.AccountValue{
					Data:  [2]string{"AAAAAAEAAADKFr/7JZLeKFJKIaGunqjtXggBBBad6ejlmbYPRoSyLJ+X5Y193+WHX7lT5pRYGWf4V70JP+EScclNbE1yU9T7iBMAAAAAAAA=", "base64"},
					Owner: "unknown",
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
			client.WithNonceAccount(tc.account)

			err := system.UpdateNonceAccount("address", 1, &db, &client)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
