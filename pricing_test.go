package oanda_test

import (
	"testing"
	"time"
)

func TestGetPrices(t *testing.T) {
	s, c := testTools(200, loadStub("./stubs/pricing/currentprices.json"))
	defer s.Close()

	prices, err := c.GetPrices("101-011-3478428-001", []string{"AUD_USD"}, time.Now().UTC().Format(time.RFC3339Nano))
	if err != nil {
		t.Error(err)
	}
	t.Log(prices)

	gotAsk := prices.Prices[0].Asks[0].Price
	gotBid := prices.Prices[0].Bids[0].Price
	expectedAsk := "1.10092"
	expectedBid := "1.10077"

	if gotAsk != expectedAsk {
		t.Errorf("ERROR: Got %s, expected %s", gotAsk, expectedAsk)
	}
	if gotBid != expectedBid {
		t.Errorf("ERROR: Got %s, expected %s", gotBid, expectedBid)
	}
}
