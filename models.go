package oanda

// AccountContainer contains a slice of AccountProperty
type AccountContainer struct {
	AccountProperties []AccountProperty `json:"accounts"`
}

// AccountProperty contain properties related to an Account.
type AccountProperty struct {
	ID           string   `json:"id"`
	MT4AccountID int      `json:"mt4AccountID"`
	Tags         []string `json:"tags"`
}

// PricesContainer contains a slice of Price
type PricesContainer struct {
	Prices []Price
}

// Price contains a current price at a single point in time
type Price struct {
	InstrumentName string        `json:"instrument"`
	Time           string        `json:"time"`
	Status         string        `json:"status"`
	Bids           []PriceBucket `json:"bids"`
	Asks           []PriceBucket `json:"asks"`
	CloseoutBid    string        `json:"closeOutBid"`
	CloseoutAsk    string        `json:"closeOutAsk"`
}

// PriceBucket contains pricing and liquidity information for a Price
type PriceBucket struct {
	Price     string `json:"price"`
	Liquidity int    `json:"liquidity"`
}
