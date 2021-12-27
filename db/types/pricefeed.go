package types

import "time"

type TokenUnitRow struct {
	Address     string `db:"address"`
	PriceID     string `db:"price_id"`
	Name        string `db:"unit_name"`
	LogoURI     string `db:"logo_uri"`
	Description string `db:"description"`
	Website     string `db:"website"`
}

func NewTokenUnitRow(
	address string,
	priceID string,
	name string,
	logoURI string,
	description string,
	website string,
) TokenUnitRow {
	return TokenUnitRow{
		Address:     address,
		PriceID:     priceID,
		Name:        name,
		LogoURI:     logoURI,
		Description: description,
		Website:     website,
	}
}

type TokenPriceRow struct {
	ID        string    `db:"id"`
	Price     float64   `db:"price"`
	MarketCap int64     `db:"market_cap"`
	Symbol    string    `db:"symbol"`
	Timestamp time.Time `db:"timestamp"`
}

func NewTokenPriceRow(id string, price float64, marketCap int64, symbol string, timestamp time.Time) TokenPriceRow {
	return TokenPriceRow{
		ID:        id,
		Price:     price,
		MarketCap: marketCap,
		Symbol:    symbol,
		Timestamp: timestamp,
	}
}
