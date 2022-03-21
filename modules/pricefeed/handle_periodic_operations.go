package pricefeed

import (
	"fmt"
	"math"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/forbole/soljuno/apis/coingecko"
	dbtypes "github.com/forbole/soljuno/db/types"
	"github.com/forbole/soljuno/modules/utils"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "pricefeed").Msg("setting up periodic tasks")

	// Fetch prices of tokens in 30 seconds each
	if _, err := scheduler.Every(30).Second().Do(func() {
		utils.WatchMethod(m, m.updatePrice)
	}); err != nil {
		return fmt.Errorf("error while setting up pricefeed period operations: %s", err)
	}

	return nil
}

// getTokenPrices gets the token prices in the database from coingecko
func (m *Module) getTokenPrices() ([]dbtypes.TokenPriceRow, error) {
	units, err := m.db.GetTokenUnits()
	if err != nil {
		return nil, fmt.Errorf("error while getting token units: %s", err)
	}
	// Find the id of the coins
	var ids []string
	for _, unit := range units {
		// Skip the token if the price id is empty
		if unit.PriceID == "" {
			continue
		}
		ids = append(ids, unit.PriceID)
	}

	if len(ids) == 0 {
		return nil, fmt.Errorf("no traded tokens found")
	}

	// Get the tokens prices
	prices, err := m.client.GetTokensPrices(ids)
	if err != nil {
		return nil, fmt.Errorf("error while getting tokens prices: %s", err)
	}

	return convertCoingeckoPrices(prices), err
}

// updatePrice fetch total amount of coins in the system from RPC and store it into database
func (m *Module) updatePrice() error {
	log.Debug().
		Str("module", "pricefeed").
		Str("operation", "pricefeed").
		Msg("getting token price and market cap")

	prices, err := m.getTokenPrices()
	if err != nil {
		return fmt.Errorf("error while getting token prices")
	}

	// Save the token prices
	err = m.db.SaveTokenPrices(prices)
	if err != nil {
		return fmt.Errorf("error while saving token prices: %s", err)
	}
	return nil
}

// convertCoingeckoPrices converts the MarketTicker list into TokenPriceRow list
func convertCoingeckoPrices(prices []coingecko.MarketTicker) []dbtypes.TokenPriceRow {
	tokenPrices := make([]dbtypes.TokenPriceRow, len(prices))
	for i, price := range prices {
		tokenPrices[i] = dbtypes.NewTokenPriceRow(
			price.ID,
			price.CurrentPrice,
			int64(math.Trunc(price.MarketCap)),
			price.Symbol,
			price.LastUpdated,
			price.Volume,
		)
	}
	return tokenPrices
}
