package postgresql_test

import (
	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestSaveConfigAccount() {
	testCases := []struct {
		name     string
		data     dbtypes.ValidatorConfigRow
		expected dbtypes.ValidatorConfigRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewValidatorConfigRow(
				"address",
				1,
				"owner",
				dbtypes.NewParsedValidatorConfig(
					"name", "", "website", "details",
				),
				"",
			),
			expected: dbtypes.NewValidatorConfigRow(
				"address",
				1,
				"owner",
				dbtypes.NewParsedValidatorConfig(
					"name", "", "website", "details",
				),
				"",
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewValidatorConfigRow(
				"address",
				0,
				"owner",
				dbtypes.NewParsedValidatorConfig(
					"name", "", "website", "pre_details",
				),
				"",
			),
			expected: dbtypes.NewValidatorConfigRow(
				"address",
				1,
				"owner",
				dbtypes.NewParsedValidatorConfig(
					"name", "", "website", "details",
				),
				"",
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewValidatorConfigRow(
				"address",
				1,
				"owner",
				dbtypes.NewParsedValidatorConfig(
					"name", "", "website", "curr_details",
				),
				"",
			),
			expected: dbtypes.NewValidatorConfigRow(
				"address",
				1,
				"owner",
				dbtypes.NewParsedValidatorConfig(
					"name", "", "website", "curr_details",
				),
				"",
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewValidatorConfigRow(
				"address",
				2,
				"owner",
				dbtypes.NewParsedValidatorConfig(
					"name", "", "website", "new_details",
				),
				"",
			),
			expected: dbtypes.NewValidatorConfigRow(
				"address",
				2,
				"owner",
				dbtypes.NewParsedValidatorConfig(
					"name", "", "website", "new_details",
				),
				"",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveValidatorConfig(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.ValidatorConfigRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM validator_config")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
