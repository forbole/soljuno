package coingecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Client interface {
	GetTokensPrices(ids []string) ([]MarketTicker, error)
}

var _ Client = &client{}

type client struct {
	endpoint string
	retries  int
}

func NewDefaultClient() Client {
	return &client{
		endpoint: "https://api.coingecko.com/api/v3",
		retries:  3,
	}
}

// GetTokensPrices queries the remote APIs to get the token prices of all the tokens having the given ids
func (c *client) GetTokensPrices(ids []string) ([]MarketTicker, error) {
	query := fmt.Sprintf("/coins/markets?vs_currency=usd&ids=%s", strings.Join(ids, ","))
	retries := c.retries
	var err error
	var prices []MarketTicker
	for retries > 0 {
		err = c.queryCoinGecko(query, &prices)
		if err != nil {
			retries -= 1
		} else {
			break
		}
		time.Sleep(time.Second)
	}
	return prices, err
}

// queryCoinGecko queries the CoinGecko APIs for the given endpoint
func (c *client) queryCoinGecko(path string, ptr interface{}) error {
	resp, err := http.Get(c.endpoint + path)
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
