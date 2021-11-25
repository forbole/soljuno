package types

type TokenUnitRow struct {
	TokenName string `db:"token_name"`
	Address   string `db:"address"`
	PriceID   string `db:"price_id"`
}
