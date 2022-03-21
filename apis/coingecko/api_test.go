package coingecko_test

import (
	"encoding/json"
	"testing"

	"github.com/forbole/soljuno/apis/coingecko"
	"github.com/stretchr/testify/require"
)

func TestConvertCoingeckoPrices(t *testing.T) {
	result := `
[
	{
		"id": "cosmos",
		"symbol": "atom",
		"name": "Cosmos Hub",
		"image": "https://assets.coingecko.com/coins/images/1481/large/cosmos_hub.png?1555657960",
		"current_price": 28.18,
		"market_cap": 8191073374,
		"market_cap_rank": 20,
		"fully_diluted_valuation": null,
		"total_volume": 398670823,
		"high_24h": 28.67,
		"low_24h": 27.59,
		"price_change_24h": -0.373685617141,
		"price_change_percentage_24h": -1.30869,
		"market_cap_change_24h": -115912658.19612217,
		"market_cap_change_percentage_24h": -1.39536,
		"circulating_supply": 290821201.019102,
		"total_supply": null,
		"max_supply": null,
		"ath": 44.45,
		"ath_change_percentage": -36.69548,
		"ath_date": "2022-01-17T00:34:41.497Z",
		"atl": 1.16,
		"atl_change_percentage": 2325.53309,
		"atl_date": "2020-03-13T02:27:44.591Z",
		"roi": {
			"times": 280.8051949323397,
			"currency": "usd",
			"percentage": 28080.51949323397
		},
		"last_updated": "2022-03-21T06:51:54.031Z"
	},
	{
		"id": "bitcanna",
		"symbol": "bcna",
		"name": "BitCanna",
		"image": "https://assets.coingecko.com/coins/images/4716/large/bcna.png?1547040016",
		"current_price": 0.086703,
		"market_cap": 0,
		"market_cap_rank": null,
		"fully_diluted_valuation": null,
		"total_volume": 248346,
		"high_24h": 0.08792,
		"low_24h": 0.085098,
		"price_change_24h": -0.000899249852,
		"price_change_percentage_24h": -1.02652,
		"market_cap_change_24h": 0,
		"market_cap_change_percentage_24h": 0,
		"circulating_supply": 0,
		"total_supply": 397622881.521035,
		"max_supply": null,
		"ath": 0.922537,
		"ath_change_percentage": -90.6066,
		"ath_date": "2020-02-06T22:44:34.222Z",
		"atl": 0.00216099,
		"atl_change_percentage": 3910.08759,
		"atl_date": "2020-01-01T11:24:08.916Z",
		"roi": {
			"times": -0.27747859436794514,
			"currency": "usd",
			"percentage": -27.747859436794517
		},
		"last_updated": "2022-03-21T06:49:17.249Z"
	},
	{
		"id": "bitcoin",
		"symbol": "btc",
		"name": "Bitcoin",
		"image": "https://assets.coingecko.com/coins/images/1/large/bitcoin.png?1547033579",
		"current_price": 40941,
		"market_cap": 777407839306,
		"market_cap_rank": 1,
		"fully_diluted_valuation": 859714534378,
		"total_volume": 17973772098,
		"high_24h": 41966,
		"low_24h": 40615,
		"price_change_24h": -1007.718483085111,
		"price_change_percentage_24h": -2.40228,
		"market_cap_change_24h": -19547093722.273804,
		"market_cap_change_percentage_24h": -2.45272,
		"circulating_supply": 18989518,
		"total_supply": 21000000,
		"max_supply": 21000000,
		"ath": 69045,
		"ath_change_percentage": -40.72825,
		"ath_date": "2021-11-10T14:24:11.849Z",
		"atl": 67.81,
		"atl_change_percentage": 60251.9368,
		"atl_date": "2013-07-06T00:00:00.000Z",
		"roi": null,
		"last_updated": "2022-03-21T06:50:55.068Z"
	}
]
`

	var apisPrices []coingecko.MarketTicker
	err := json.Unmarshal([]byte(result), &apisPrices)
	require.NoError(t, err)

	require.Equal(t, float64(8191073374), apisPrices[0].MarketCap)
	require.Equal(t, float64(0), apisPrices[1].MarketCap)
	require.Equal(t, float64(777407839306), apisPrices[2].MarketCap)
}
