package postgresql

func (db *Database) PruneTxsBySlot(slot uint64) error {
	_, err := db.Sqlx.Exec(`DELETE FROM transaction WHERE slot <= $1`, slot)
	return err
}
