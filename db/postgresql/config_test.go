package postgresql_test

import "encoding/json"

func (suite *DbTestSuite) TestSaveConfigAccount() {
	type ConfigAccountRow struct {
		Address         string `db:"address"`
		Slot            uint64 `db:"slot"`
		Owner           string `db:"owner"`
		Name            string `db:"name"`
		KeybaseUsername string `db:"keybase_username"`
		Website         string `db:"website"`
		Details         string `db:"details"`
	}

	type parsedInfo struct {
		Name            string `json:"name"`
		KeybaseUsername string `json:"keybaseUsername"`
		Website         string `json:"website"`
		Details         string `json:"details"`
	}

	testCases := []struct {
		name     string
		data     ConfigAccountRow
		expected ConfigAccountRow
	}{
		{
			name: "initialize the data",
			data: ConfigAccountRow{
				"address", 1, "owner", "name", "", "website", "details",
			},
			expected: ConfigAccountRow{
				"address", 1, "owner", "name", "", "website", "details",
			},
		},
		{
			name: "update with lower slot",
			data: ConfigAccountRow{
				"address", 0, "owner", "name", "", "website", "pre_details",
			},
			expected: ConfigAccountRow{
				"address", 1, "owner", "name", "", "website", "details",
			},
		},
		{
			name: "update with same slot",
			data: ConfigAccountRow{
				"address", 1, "owner", "name", "", "website", "curr_details",
			},
			expected: ConfigAccountRow{
				"address", 1, "owner", "name", "", "website", "curr_details",
			},
		},
		{
			name: "update with higher slot",
			data: ConfigAccountRow{
				"address", 2, "owner", "name", "", "website", "new_details",
			},
			expected: ConfigAccountRow{
				"address", 2, "owner", "name", "", "website", "new_details",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			value, err := json.Marshal(
				parsedInfo{
					tc.data.Name,
					tc.data.KeybaseUsername,
					tc.data.Website,
					tc.data.Details,
				},
			)
			err = suite.database.SaveConfigAccount(
				tc.data.Address,
				tc.data.Slot,
				tc.data.Owner,
				string(value),
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []ConfigAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM validator_config")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
