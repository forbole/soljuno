package postgresql

// CheckNonceAccountLatest implements the db.SystemCheckerDb
// if error, it returns false since there is no address statement inside database
func (db *Database) CheckNonceAccountLatest(address string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM nonce_account WHERE address=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, address)
	if err != nil {
		return false
	}
	return latestSlot >= currentSlot
}
