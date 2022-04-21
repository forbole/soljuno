package postgresql

// CheckBufferAccountLatest implements the db.StakeCheckerDb
// if error, it returns false since there is no address statement inside database
func (db *Database) CheckBufferAccountLatest(address string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM buffer_account WHERE address=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, address)
	if err != nil {
		return false
	}
	return currentSlot > latestSlot
}

// CheckProgramAccountLatest implements the db.StakeCheckerDb
// if error, it returns false since there is no address statement inside database
func (db *Database) CheckProgramAccountLatest(address string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM program_account WHERE address=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, address)
	if err != nil {
		return false
	}
	return currentSlot > latestSlot
}

// CheckProgramDataAccountLatest implements the db.StakeCheckerDb
// if error, it returns false since there is no address statement inside database
func (db *Database) CheckProgramDataAccountLatest(address string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM program_data_account WHERE address=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, address)
	if err != nil {
		return false
	}
	return currentSlot > latestSlot
}
