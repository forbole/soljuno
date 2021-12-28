package coingecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	dbtypes "github.com/forbole/soljuno/db/types"
)

// GetCoinsList allows to fetch from the remote APIs the list of all the supported tokens
func GetCoinsList() (coins Tokens, err error) {
	err = queryCoinGecko("/coins/list", &coins)
	return coins, err
}

// GetTokensPrices queries the remote APIs to get the token prices of all the tokens having the given ids
func GetTokensPrices(ids []string) ([]dbtypes.TokenPriceRow, error) {
	var prices []MarketTicker
	query := fmt.Sprintf("/coins/markets?vs_currency=usd&ids=%s", strings.Join(ids, ","))
	err := queryCoinGecko(query, &prices)
	if err != nil {
		return nil, err
	}

	return ConvertCoingeckoPrices(prices), nil
}

func ConvertCoingeckoPrices(prices []MarketTicker) []dbtypes.TokenPriceRow {
	tokenPrices := make([]dbtypes.TokenPriceRow, len(prices))
	for i, price := range prices {
		tokenPrices[i] = dbtypes.NewTokenPriceRow(
			price.ID,
			price.CurrentPrice,
			int64(math.Trunc(price.MarketCap)),
			price.Symbol,
			price.LastUpdated,
		)
	}
	return tokenPrices
}

// queryCoinGecko queries the CoinGecko APIs for the given endpoint
func queryCoinGecko(endpoint string, ptr interface{}) error {
	resp, err := http.Get("https://api.coingecko.com/api/v3" + endpoint)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error while reading response body: %s", err)
	}

	err = json.Unmarshal(bz, &ptr)
	if err != nil {
		return fmt.Errorf("error while unmarshaling response body: %s", err)
	}

	return nil
}
