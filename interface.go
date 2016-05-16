package oanda

type Accounter interface {
	ListAccounts() (*Account, error)
	AccountDetails()
	AccountSummary()
	AccountInstruments()
	ConfigureAccount()
	PollAccountUpdates()
}

type Transactioner interface {
	ListTransactions()
	TransactionDetails()
	TransactionIDRange()
	TransactionsSinceID()
}

type Trader interface {
	ListTrades()
	ListOpenTrades()
	TradeDetails()
	CloseTrade()
	SetTradeClientExtensions()
	SetDependentOrders()
}

type Pricer interface {
	GetPrices(accountID string, instruments []string, since string)
}

type Positioner interface {
	OpenPositions()
	ListPositions()
	InstrumentPosition()
	ClosePosition()
}
type Loginer interface {
	Login()
	Logout()
}
type Orderer interface {
	ListOrders()
	CreateOrder()
	FetchOrder()
	ReplaceOrder()
	CancelOrder()
	SetOrderClientExtensions()
	PendingOrders()
}
type Userer interface {
	UserAccountList()
}
type Streamer interface{}
