package instrument

// CandlesContainer contains a slice of Candle
type CandlesContainer struct {
	Candles []Candle
}

// Candle contains informations concerning a Candle
type Candle struct {
	Ask      OHLC   `json:"ask"`
	Bid      OHLC   `json:"bid"`
	Mid      OHLC   `json:"mid"`
	Complete bool   `json:"complete"`
	Time     string `json:"time"`
	Volume   int    `json:"volume"`
}

// OHLC contains open, high, low and close prices
type OHLC struct {
	Open  string `json:"o"`
	High  string `json:"h"`
	Low   string `json:"l"`
	Close string `json:"c"`
}
