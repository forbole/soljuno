package postgresql_test

func (suite *DbTestSuite) TestSaveBufferAccount() {
	type BufferAccountRow struct {
		Mint      string `db:"address"`
		Slot      uint64 `db:"slot"`
		Authority string `db:"authority"`
		State     string `db:"state"`
	}

	testCases := []struct {
		name     string
		data     BufferAccountRow
		expected BufferAccountRow
	}{
		{
			name: "initialize the data",
			data: BufferAccountRow{
				"address", 1, "owner", "initialized",
			},
			expected: BufferAccountRow{
				"address", 1, "owner", "initialized",
			},
		},
		{
			name: "update with lower slot",
			data: BufferAccountRow{
				"address", 0, "pre_owner", "initialized",
			},
			expected: BufferAccountRow{
				"address", 1, "owner", "initialized",
			},
		},
		{
			name: "update with same slot",
			data: BufferAccountRow{
				"address", 1, "cur_owner", "initialized",
			},
			expected: BufferAccountRow{
				"address", 1, "cur_owner", "initialized",
			},
		},
		{
			name: "update with higher slot",
			data: BufferAccountRow{
				"address", 2, "new_owner", "initialized",
			},
			expected: BufferAccountRow{
				"address", 2, "new_owner", "initialized",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveBufferAccount(
				tc.data.Mint,
				tc.data.Slot,
				tc.data.Authority,
				tc.data.State,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []BufferAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM buffer_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveProgramAccount() {
	type ProgramAccountAccountRow struct {
		Mint               string `db:"address"`
		Slot               uint64 `db:"slot"`
		ProgramDataAccount string `db:"program_data_account"`
		State              string `db:"state"`
	}

	testCases := []struct {
		name     string
		data     ProgramAccountAccountRow
		expected ProgramAccountAccountRow
	}{
		{
			name: "initialize the data",
			data: ProgramAccountAccountRow{
				"address", 1, "program_data", "initialized",
			},
			expected: ProgramAccountAccountRow{
				"address", 1, "program_data", "initialized",
			},
		},
		{
			name: "update with lower slot",
			data: ProgramAccountAccountRow{
				"address", 0, "pre_program_data", "initialized",
			},
			expected: ProgramAccountAccountRow{
				"address", 1, "program_data", "initialized",
			},
		},
		{
			name: "update with same slot",
			data: ProgramAccountAccountRow{
				"address", 1, "cur_program_data", "initialized",
			},
			expected: ProgramAccountAccountRow{
				"address", 1, "cur_program_data", "initialized",
			},
		},
		{
			name: "update with higher slot",
			data: ProgramAccountAccountRow{
				"address", 2, "new_program_data", "initialized",
			},
			expected: ProgramAccountAccountRow{
				"address", 2, "new_program_data", "initialized",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveProgramAccount(
				tc.data.Mint,
				tc.data.Slot,
				tc.data.ProgramDataAccount,
				tc.data.State,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []ProgramAccountAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM program_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestSaveProgramDataAccount() {
	type ProgramDataAccountRow struct {
		Mint             string `db:"address"`
		Slot             uint64 `db:"slot"`
		LastModifiedSlot uint64 `db:"last_modified_slot"`
		UpdateAuthority  string `db:"update_authority"`
		State            string `db:"state"`
	}

	testCases := []struct {
		name     string
		data     ProgramDataAccountRow
		expected ProgramDataAccountRow
	}{
		{
			name: "initialize the data",
			data: ProgramDataAccountRow{
				"address", 1, 0, "owner", "initialized",
			},
			expected: ProgramDataAccountRow{
				"address", 1, 0, "owner", "initialized",
			},
		},
		{
			name: "update with lower slot",
			data: ProgramDataAccountRow{
				"address", 0, 0, "pre_owner", "initialized",
			},
			expected: ProgramDataAccountRow{
				"address", 1, 0, "owner", "initialized",
			},
		},
		{
			name: "update with same slot",
			data: ProgramDataAccountRow{
				"address", 1, 0, "cur_owner", "initialized",
			},
			expected: ProgramDataAccountRow{
				"address", 1, 0, "cur_owner", "initialized",
			},
		},
		{
			name: "update with higher slot",
			data: ProgramDataAccountRow{
				"address", 2, 0, "new_owner", "initialized",
			},
			expected: ProgramDataAccountRow{
				"address", 2, 0, "new_owner", "initialized",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveProgramDataAccount(
				tc.data.Mint,
				tc.data.Slot,
				tc.data.LastModifiedSlot,
				tc.data.UpdateAuthority,
				tc.data.State,
			)
			suite.Require().NoError(err)

			// Verify the data
			rows := []ProgramDataAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM program_data_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}
