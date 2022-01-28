package postgresql_test

import "github.com/forbole/soljuno/db/postgresql"

func (suite *DbTestSuite) TestInsertBatch() {
	_, err := suite.database.Exec("CREATE TABLE test (id INT);")
	suite.Require().NoError(err)

	type testRow struct {
		ID int `db:"id"`
	}

	testCases := []struct {
		name        string
		data        []int
		expectedLen int
	}{
		{
			name:        "insert empty data",
			data:        make([]int, 0),
			expectedLen: 0,
		},
		{
			name:        "insert 1 data",
			data:        make([]int, 1),
			expectedLen: 1,
		},
		{
			name:        "insert max length data",
			data:        make([]int, postgresql.MAX_PARAMS_LENGTH),
			expectedLen: postgresql.MAX_PARAMS_LENGTH,
		},
		{
			name:        "insert over max length data",
			data:        make([]int, postgresql.MAX_PARAMS_LENGTH+1),
			expectedLen: postgresql.MAX_PARAMS_LENGTH + 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			_, err = suite.database.Exec("DELETE FROM test;")
			suite.Require().NoError(err)

			insertData := make([]interface{}, len(tc.data))
			for i, v := range tc.data {
				insertData[i] = v
			}

			err := suite.database.InsertBatch("INSERT INTO test VALUES", "", insertData, 1)
			suite.Require().NoError(err)
			var rows []testRow
			err = suite.database.Sqlx.Select(&rows, "SELECT id FROM test")
			suite.Require().Len(rows, tc.expectedLen)
		})
	}
}
