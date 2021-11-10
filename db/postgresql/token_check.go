package postgresql

// CheckTokenLatest implements the db.TokenCheckerDb
// if error, it returns true since there is no address statement inside database
func (db *Database) CheckTokenLatest(address string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM token WHERE mint=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, address)
	if err != nil {
		return true
	}
	return currentSlot > latestSlot
}

// CheckTokenAccountLatest implements the db.TokenCheckerDb
// if error, it returns true since there is no address statement inside database
func (db *Database) CheckTokenAccountLatest(address string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM token_account WHERE address=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, address)
	if err != nil {
		return true
	}
	return currentSlot > latestSlot
}

// CheckMultisigLatest implements the db.TokenCheckerDb
// if error, it returns true since there is no address statement inside database
func (db *Database) CheckMultisigLatest(address string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM multisig WHERE address=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, address)
	if err != nil {
		return true
	}
	return currentSlot > latestSlot
}

// CheckDelegateLatest implements the db.TokenCheckerDb
// if error, it returns true since there is no address statement inside database
func (db *Database) CheckTokenDelegateLatest(sourceAddress string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM token_delegation WHERE source_address=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, sourceAddress)
	if err != nil {
		return true
	}
	return currentSlot > latestSlot
}

// CheckTokenSupplyLatest implements the db.TokenCheckerDb
// if error, it returns true since there is no address statement inside database
func (db *Database) CheckTokenSupplyLatest(address string, currentSlot uint64) bool {
	stmt := `SELECT slot FROM token_supply WHERE mint=$1`
	var latestSlot uint64
	err := db.Sqlx.Get(&latestSlot, stmt, address)
	if err != nil {
		return true
	}
	return currentSlot > latestSlot
}
