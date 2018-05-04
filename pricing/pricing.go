package pricing

import "time"

// GetResponse is returned from the API
type GetResponse struct {
	Time   time.Time `json:"time"`
	Prices []struct {
		Type string    `json:"type"`
		Time time.Time `json:"time"`
		Bids []struct {
			Price     string `json:"price"`
			Liquidity int    `json:"liquidity"`
		} `json:"bids"`
		Asks []struct {
			Price     string `json:"price"`
			Liquidity int    `json:"liquidity"`
		} `json:"asks"`
		CloseoutBid    string `json:"closeoutBid"`
		CloseoutAsk    string `json:"closeoutAsk"`
		Status         string `json:"status"`
		Tradeable      bool   `json:"tradeable"`
		UnitsAvailable struct {
			Default struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"default"`
			OpenOnly struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"openOnly"`
			ReduceFirst struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"reduceFirst"`
			ReduceOnly struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"reduceOnly"`
		} `json:"unitsAvailable"`
		QuoteHomeConversionFactors struct {
			PositiveUnits string `json:"positiveUnits"`
			NegativeUnits string `json:"negativeUnits"`
		} `json:"quoteHomeConversionFactors"`
		Instrument string `json:"instrument"`
	} `json:"prices"`
}

// StreamResponse is the response from the streaming pricing API
type StreamResponse struct {
	Type string    `json:"type"`
	Time time.Time `json:"time"`
	Bids []struct {
		Price     string `json:"price"`
		Liquidity int    `json:"liquidity"`
	} `json:"bids"`
	Asks []struct {
		Price     string `json:"price"`
		Liquidity int    `json:"liquidity"`
	} `json:"asks"`
	CloseoutBid string `json:"closeoutBid"`
	CloseoutAsk string `json:"closeoutAsk"`
	Status      string `json:"status"`
	Tradeable   bool   `json:"tradeable"`
	Instrument  string `json:"instrument"`
}
