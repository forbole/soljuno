package postgresql_test

import dbtypes "github.com/forbole/soljuno/db/types"

func (suite *DbTestSuite) TestSaveBufferAccount() {
	testCases := []struct {
		name     string
		data     dbtypes.BufferAccountRow
		expected dbtypes.BufferAccountRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewBufferAccountRow(
				"address", 1, "owner",
			),
			expected: dbtypes.NewBufferAccountRow(
				"address", 1, "owner",
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewBufferAccountRow(
				"address", 0, "pre_owner",
			),
			expected: dbtypes.NewBufferAccountRow(
				"address", 1, "owner",
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewBufferAccountRow(
				"address", 1, "cur_owner",
			),
			expected: dbtypes.NewBufferAccountRow(
				"address", 1, "cur_owner",
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewBufferAccountRow(
				"address", 2, "new_owner",
			),
			expected: dbtypes.NewBufferAccountRow(
				"address", 2, "new_owner",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveBufferAccount(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.BufferAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM buffer_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestDeleteBufferAccount() {
	err := suite.database.SaveBufferAccount(
		dbtypes.NewBufferAccountRow(
			"address",
			0,
			"owner",
		),
	)
	suite.Require().NoError(err)

	rows := []struct {
		Address   string `db:"address"`
		Slot      uint64 `db:"slot"`
		Authority string `db:"authority"`
		State     string `db:"state"`
	}{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM buffer_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	rows = nil

	err = suite.database.DeleteBufferAccount("address")
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM buffer_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}

func (suite *DbTestSuite) TestSaveProgramAccount() {

	testCases := []struct {
		name     string
		data     dbtypes.ProgramAccountRow
		expected dbtypes.ProgramAccountRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewProgramAccountRow(
				"address", 1, "program_data",
			),
			expected: dbtypes.NewProgramAccountRow(
				"address", 1, "program_data",
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewProgramAccountRow(
				"address", 0, "pre_program_data",
			),
			expected: dbtypes.NewProgramAccountRow(
				"address", 1, "program_data",
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewProgramAccountRow(
				"address", 1, "cur_program_data",
			),
			expected: dbtypes.NewProgramAccountRow(
				"address", 1, "cur_program_data",
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewProgramAccountRow(
				"address", 2, "new_program_data",
			),
			expected: dbtypes.NewProgramAccountRow(
				"address", 2, "new_program_data",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveProgramAccount(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.ProgramAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM program_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestDeleteProgramAccount() {
	err := suite.database.SaveProgramAccount(
		dbtypes.NewProgramAccountRow(
			"address",
			0,
			"data",
		),
	)
	suite.Require().NoError(err)

	rows := []struct {
		Address            string `db:"address"`
		Slot               uint64 `db:"slot"`
		ProgramDataAccount string `db:"program_data_account"`
	}{}
	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM program_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	rows = nil

	err = suite.database.DeleteProgramAccount("address")
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM program_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}

func (suite *DbTestSuite) TestSaveProgramDataAccount() {
	testCases := []struct {
		name     string
		data     dbtypes.ProgramDataAccountRow
		expected dbtypes.ProgramDataAccountRow
	}{
		{
			name: "initialize the data",
			data: dbtypes.NewProgramDataAccountRow(
				"address", 1, 0, "owner",
			),
			expected: dbtypes.NewProgramDataAccountRow(
				"address", 1, 0, "owner",
			),
		},
		{
			name: "update with lower slot",
			data: dbtypes.NewProgramDataAccountRow(
				"address", 0, 0, "pre_owner",
			),
			expected: dbtypes.NewProgramDataAccountRow(
				"address", 1, 0, "owner",
			),
		},
		{
			name: "update with same slot",
			data: dbtypes.NewProgramDataAccountRow(
				"address", 1, 0, "cur_owner",
			),
			expected: dbtypes.NewProgramDataAccountRow(
				"address", 1, 0, "cur_owner",
			),
		},
		{
			name: "update with higher slot",
			data: dbtypes.NewProgramDataAccountRow(
				"address", 2, 0, "new_owner",
			),
			expected: dbtypes.NewProgramDataAccountRow(
				"address", 2, 0, "new_owner",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			err := suite.database.SaveProgramDataAccount(tc.data)
			suite.Require().NoError(err)

			// Verify the data
			rows := []dbtypes.ProgramDataAccountRow{}
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM program_data_account")
			suite.Require().NoError(err)
			suite.Require().Len(rows, 1)
			suite.Require().Equal(tc.expected, rows[0])
		})
	}
}

func (suite *DbTestSuite) TestDeleteProgramDataAccount() {
	err := suite.database.SaveProgramDataAccount(
		dbtypes.NewProgramDataAccountRow(
			"address",
			0,
			0,
			"owner",
		),
	)
	suite.Require().NoError(err)
	rows := []struct {
		Address          string `db:"address"`
		Slot             uint64 `db:"slot"`
		LastModifiedSlot uint64 `db:"last_modified_slot"`
		UpdateAuthority  string `db:"update_authority"`
	}{}

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM program_data_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	rows = nil

	err = suite.database.DeleteProgramDataAccount("address")
	suite.Require().NoError(err)

	err = suite.database.Sqlx.Select(&rows, "SELECT * FROM program_data_account")
	suite.Require().NoError(err)
	suite.Require().Len(rows, 0)
}
