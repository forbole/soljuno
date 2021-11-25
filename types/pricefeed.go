package types

import "time"

// TokenUnit represents a unit of a token
type TokenUnit struct {
	ID      string
	Address string
	Name    string
}

func NewTokenUnit(id string, address string, name string) TokenUnit {
	return TokenUnit{
		ID:      id,
		Address: address,
		Name:    name,
	}
}

// TokenPrice represents the price at a given moment in time of a token unit
type TokenPrice struct {
	UnitName  string
	Price     float64
	MarketCap int64
	Timestamp time.Time
}

// NewTokenPrice returns a new TokenPrice instance containing the given data
func NewTokenPrice(unitName string, price float64, marketCap int64, timestamp time.Time) TokenPrice {
	return TokenPrice{
		UnitName:  unitName,
		Price:     price,
		MarketCap: marketCap,
		Timestamp: timestamp,
	}
}