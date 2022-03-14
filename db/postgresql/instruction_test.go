package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestSaveInstructions() {
	err := suite.database.CreateInstructionPartition(0)
	suite.Require().NoError(err)

	testCases := []struct {
		name        string
		data        dbtypes.InstructionRow
		expectedLen int
		shouldErr   bool
	}{
		{
			name:        "initialize the data",
			data:        dbtypes.NewInstructionRow("hash", 1, 1, 0, "program", []string{"address"}, "raw", "type", "{}"),
			expectedLen: 1,
		},
		{
			name: "insert the wrong tx",
			data: dbtypes.InstructionRow{
				TxHash:           "txHash",
				Slot:             1,
				Index:            1,
				InnerIndex:       0,
				Program:          "program",
				InvolvedAccounts: []string{"address"},
				RawData:          "raw",
				Type:             "type",
				Value:            "{}",
				PartitionId:      100,
			},
			shouldErr: true,
		},
		{
			name:        "insert the existed data",
			data:        dbtypes.NewInstructionRow("hash", 1, 1, 0, "program", []string{"address"}, "raw", "type", "{}"),
			expectedLen: 1,
		},
		{
			name:        "insert the new data",
			data:        dbtypes.NewInstructionRow("hash", 1, 1, 1, "program", []string{"address"}, "raw", "type", "{}"),
			expectedLen: 2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveInstructions([]dbtypes.InstructionRow{tc.data})
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				// Verify the data
				rows := []dbtypes.InstructionRow{}
				err = suite.database.Sqlx.Select(&rows, "SELECT * FROM instruction")
				suite.Require().NoError(err)
				suite.Require().Len(rows, tc.expectedLen)
			}
		})
	}
}

func (suite *DbTestSuite) TestPruneInstructionsBeforeSlot() {
	err := suite.database.CreateInstructionPartition(0)
	suite.Require().NoError(err)

	err = suite.database.SaveInstructions([]dbtypes.InstructionRow{
		dbtypes.NewInstructionRow("hash", 1, 1, 0, "program", []string{"address"}, "raw", "type", "{}"),
	})
	suite.Require().NoError(err)

	err = suite.database.PruneInstructionsBeforeSlot(10000)
	suite.Require().NoError(err)

	rows := []dbtypes.InstructionRow{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM instruction")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}
