package oanda_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/nii236/oanda"
)

var listAccountResponse = `{
  "accounts":
    [
      {
        "id":"123-456-7891234-567","tags":
          [
          ]
      }
    ]
}`

var currentPricesResponse = `{
  "prices": [
    {
      "unitsAvailable": {
        "reduceOnly": {
          "short": "0",
          "long": "0"
        },
        "reduceFirst": {
          "short": "3272221",
          "long": "3270843"
        },
        "openOnly": {
          "short": "3272221",
          "long": "3270843"
        },
        "default": {
          "short": "3272221",
          "long": "3270843"
        }
      },
      "asks": [
        {
          "price": "1.10092",
          "liquidity": 2000000
        },
        {
          "price": "1.10093",
          "liquidity": 5000000
        },
        {
          "price": "1.10094",
          "liquidity": 10000000
        }
      ],
      "bids": [
        {
          "price": "1.10077",
          "liquidity": 2000000
        },
        {
          "price": "1.10076",
          "liquidity": 5000000
        },
        {
          "price": "1.10075",
          "liquidity": 10000000
        }
      ],
      "closeoutAsk": "1.10094",
      "closeoutBid": "1.10075",
      "instrument": "EUR_USD",
      "quoteHomeConversionFactors": {
        "positiveUnits": "1.38813698",
        "negativeUnits": "1.38852247"
        },
        "status": "tradeable",
        "time": "2016-02-24T21:20:09.022253858Z"
    }
  ]
}`

func testTools(code int, body string) (*httptest.Server, *oanda.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	client := &oanda.Client{
		Client: httpClient,
		Host:   server.URL,
	}

	return server, client
}

func TestListAccounts(t *testing.T) {
	s, c := testTools(200, listAccountResponse)
	defer s.Close()
	accounts, err := c.ListAccounts()
	if err != nil {
		t.Error(errors.Cause(err))
		return
	}
	t.Log(accounts)

	got := accounts.AccountProperties[0].ID
	expected := "123-456-7891234-567"

	if got != expected {
		t.Errorf("ERROR: Got %s, expected %s", got, expected)
	}
}

func TestGetPrices(t *testing.T) {
	s, c := testTools(200, currentPricesResponse)
	defer s.Close()

	prices, err := c.GetPrices("101-011-3478428-001", []string{"AUD_USD"}, time.Now().UTC().Format(time.RFC3339Nano))
	if err != nil {
		t.Error(err)
	}
	t.Log(prices)

	gotAsk := prices.Prices[0].Asks[0].Price
	gotBid := prices.Prices[0].Bids[0].Price
	expectedAsk:= "1.10092"
	expectedBid:= "1.10077"

	if gotAsk != expectedAsk{
		t.Errorf("ERROR: Got %s, expected %s", gotAsk, expectedAsk)
	}
	if gotBid != expectedBid{
		t.Errorf("ERROR: Got %s, expected %s", gotBid, expectedBid)
	}
}
