package postgresql_test

import "github.com/forbole/soljuno/db/postgresql"

func (suite *DbTestSuite) TestInsertBatch() {
	_, err := suite.database.Exec("CREATE TABLE test (id INT, num INT);")
	suite.Require().NoError(err)

	type testRow struct {
		ID  int `db:"id"`
		Num int `db:"num"`
	}

	testCases := []struct {
		name        string
		dataLen     int
		expectedLen int
	}{
		{
			name:        "insert empty data",
			dataLen:     0,
			expectedLen: 0,
		},
		{
			name:        "insert 1 data",
			dataLen:     1,
			expectedLen: 1,
		},
		{
			name:        "insert max length data",
			dataLen:     postgresql.MAX_PARAMS_LENGTH,
			expectedLen: postgresql.MAX_PARAMS_LENGTH,
		},
		{
			name:        "insert over max length data",
			dataLen:     postgresql.MAX_PARAMS_LENGTH + 1,
			expectedLen: postgresql.MAX_PARAMS_LENGTH + 1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			_, err = suite.database.Exec("DELETE FROM test;")
			suite.Require().NoError(err)

			insertData := make([]interface{}, 2*tc.dataLen)
			for i := 0; i < tc.dataLen; i++ {
				bi := 2 * i
				insertData[bi] = 0
				insertData[bi+1] = 0
			}

			err := suite.database.InsertBatch("INSERT INTO test VALUES", "", insertData, 2)
			suite.Require().NoError(err)
			var rows []testRow
			err = suite.database.Sqlx.Select(&rows, "SELECT * FROM test")
			suite.Require().NoError(err)
			suite.Require().Len(rows, tc.expectedLen)
		})
	}
}
