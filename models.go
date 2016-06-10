package oanda

// AccountsContainer contains a slice of AccountProperty
type AccountsContainer struct {
	AccountProperties []AccountProperty `json:"accounts"`
}

// AccountProperty contain properties related to an Account.
type AccountProperty struct {
	ID           string   `json:"id"`
	MT4AccountID int      `json:"mt4AccountID"`
	Tags         []string `json:"tags"`
}

// AccountContainer holds an Account
type AccountContainer struct {
	Account Account `json:"account"`
}

// Account represents a user's account
type Account struct {
	ID                          string `json:"id"`
	WithdrawalLimit             string `json:"withdrawalLimit"`
	UnrealizedPL                string `json:"unrealizedPL"`
	ResettablePL                string `json:"resettablePL"`
	PositionValue               string `json:"positionValue"`
	PL                          string `json:"pl"`
	PendingOrderCount           int    `json:"pendingOrderCount"`
	OpenTradeCount              int    `json:"openTradeCount"`
	OpenPositionCount           int    `json:"openPositionCount"`
	MarginUsed                  string `json:"marginUsed"`
	HedgingEnabled              bool   `json:"hedgingEnabled"`
	Currency                    string `json:"currency"`
	CreatedTime                 string `json:"createdTime"`
	CreatedByUserID             int    `json:"createdByUserID"`
	Balance                     string `json:"balance"`
	Alias                       string `json:"alias"`
	NAV                         string `json:"NAV"`
	MarginAvailable             string `json:"marginAvailable"`
	MarginCloseoutMarginUsed    string `json:"marginCloseoutMarginUsed"`
	MarginCloseoutNAV           string `json:"marginCloseoutNAV"`
	MarginCloseoutPercent       string `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue string `json:"marginCloseoutPositionValue"`
	MarginCloseoutUnrealizedPL  string `json:"marginCloseoutUnrealizedPL"`
	MarginRate                  string `json:"marginRate"`
	LastTransactionID           string `json:"lastTransactionID"`
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
