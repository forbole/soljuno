package postgresql

func (db *Database) PruneTxsBySlot(slot uint64) error {
	_, err := db.Sqlx.Exec(`DELETE FROM transaction WHERE slot <= $1`, slot)
	return err
}

func (db *Database) PruneMsgsBySlot(slot uint64) error {
	_, err := db.Sqlx.Exec(`DELETE FROM message WHERE slot <= $1`, slot)
	if err != nil {
		return err
	}
	_, err = db.Sqlx.Exec(`DELETE FROM message_by_address WHERE slot <= $1`, slot)
	return err
}
