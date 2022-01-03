package postgresql_test

import (
	dbtypes "github.com/forbole/soljuno/db/types"
)

func (suite *DbTestSuite) TestSaveSupplyInfo() {
	testCases := []struct {
		name     string
		data     dbtypes.SupplyInfoRow
		expected dbtypes.SupplyInfoRow
	}{
		{
			name:     "initialize the data",
			data:     dbtypes.NewSupplyInfoRow(2, 1, 1, 1),
			expected: dbtypes.NewSupplyInfoRow(2, 1, 1, 1),
		},
		{
			name:     "update with lower epoch",
			data:     dbtypes.NewSupplyInfoRow(1, 2, 1, 1),
			expected: dbtypes.NewSupplyInfoRow(2, 1, 1, 1),
		},
		{
			name:     "update with same epoch",
			data:     dbtypes.NewSupplyInfoRow(2, 2, 1, 1),
			expected: dbtypes.NewSupplyInfoRow(2, 2, 1, 1),
		},
		{
			name:     "update with higher epoch",
			data:     dbtypes.NewSupplyInfoRow(3, 3, 1, 1),
			expected: dbtypes.NewSupplyInfoRow(3, 3, 1, 1),
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveSupplyInfo(tc.data)
			suite.Require().NoError(err)
			var rows []dbtypes.SupplyInfoRow
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM supply_info")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
