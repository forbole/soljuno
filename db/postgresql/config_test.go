package postgresql_test

func (suite *DbTestSuite) TestSaveConfigAccount() {
	type ConfigAccountRow struct {
		Address string `db:"address"`
		Slot    uint64 `db:"slot"`
		Owner   string `db:"owner"`
		Value   string `db:"value"`
	}

	testCases := []struct {
		name     string
		data     ConfigAccountRow
		expected ConfigAccountRow
	}{
		{
			name: "initialize the data",
			data: ConfigAccountRow{
				"address", 1, "owner", `{"details": "details"}`,
			},
			expected: ConfigAccountRow{
				"address", 1, "owner", `{"details": "details"}`,
			},
		},
		{
			name: "update with lower slot",
			data: ConfigAccountRow{
				"address", 0, "owner", `{"details": "pre_details"}`,
			},
			expected: ConfigAccountRow{
				"address", 1, "owner", `{"details": "details"}`,
			},
		},
		{
			name: "update with same slot",
			data: ConfigAccountRow{
				"address", 1, "owner", `{"details": "curr_details"}`,
			},
			expected: ConfigAccountRow{
				"address", 1, "owner", `{"details": "curr_details"}`,
			},
		},
		{
			name: "update with higher slot",
			data: ConfigAccountRow{
				"address", 2, "owner", `{"details": "new_details"}`,
			},
			expected: ConfigAccountRow{
				"address", 2, "owner", `{"details": "new_details"}`,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveConfigAccount(
				tc.data.Address,
				tc.data.Slot,
				tc.data.Owner,
				tc.data.Value,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []ConfigAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM config_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
