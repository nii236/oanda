package oanda


// Accounter contains methods that handle account information and details for a user
type Accounter interface {
	ListAccounts() (*AccountContainer, error)
	AccountDetails()
	AccountSummary()
	AccountInstruments()
	ConfigureAccount()
	PollAccountUpdates()
}

// Transactioner contains methods that handle transaction information for a user
type Transactioner interface {
	ListTransactions()
	TransactionDetails()
	TransactionIDRange()
	TransactionsSinceID()
}

// Trader contains methods that handle trade information for a user
type Trader interface {
	ListTrades()
	ListOpenTrades()
	TradeDetails()
	CloseTrade()
	SetTradeClientExtensions()
	SetDependentOrders()
}

// Pricer contains methods that handle pricing information for a user
type Pricer interface {
	GetPrices(accountID string, instruments []string, since string)
}

// Positioner contains methods that handle position information for a user
type Positioner interface {
	OpenPositions()
	ListPositions()
	InstrumentPosition()
	ClosePosition()
}

// Loginer contains methods that handle login information for a user
type Loginer interface {
	Login()
	Logout()
}

// Orderer contains methods that handle order information for a user
type Orderer interface {
	ListOrders()
	CreateOrder()
	FetchOrder()
	ReplaceOrder()
	CancelOrder()
	SetOrderClientExtensions()
	PendingOrders()
}

// Userer contains methods that handle user information for a user
type Userer interface {
	UserAccountList()
}
// Streamer contains methods that handle streaming information for a user
type Streamer interface{}
