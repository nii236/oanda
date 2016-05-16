package oanda

type Account struct {
	AccountProperties []AccountProperty `json:"accounts"`
}

type AccountProperty struct {
	ID           string   `json:"id"`
	MT4AccountID int      `json:"mt4AccountID"`
	Tags         []string `json:"tags"`
}

type AccountID struct {
	Type    string `json:"Type"`
	Format  string `json:"Format"`
	Example string `json:"Example"`
}

type PricesContainer struct {
	Prices []Price
}

type Price struct {
	InstrumentName string        `json:"instrument"`
	Time           string        `json:"time"`
	Status         string        `json:"status"`
	Bids           []PriceBucket `json:"bids"`
	Asks           []PriceBucket `json:"asks"`
	CloseoutBid    string        `json:"closeOutBid"`
	CloseoutAsk    string        `json:"closeOutAsk"`
}

type PriceBucket struct {
	Price     string `json:"price"`
	Liquidity int    `json:"liquidity"`
}
