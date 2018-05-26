package account

import "time"

// ListResponse is returned from the API
type ListResponse struct {
	Accounts []struct {
		ID   string        `json:"id"`
		Tags []interface{} `json:"tags"`
	} `json:"accounts"`
}

// DetailsResponse is the account details returned from the API
type DetailsResponse struct {
	Account struct {
		GuaranteedStopLossOrderMode string        `json:"guaranteedStopLossOrderMode"`
		ID                          string        `json:"id"`
		CreatedTime                 time.Time     `json:"createdTime"`
		Currency                    string        `json:"currency"`
		CreatedByUserID             int           `json:"createdByUserID"`
		Alias                       string        `json:"alias"`
		MarginRate                  string        `json:"marginRate"`
		HedgingEnabled              bool          `json:"hedgingEnabled"`
		LastTransactionID           string        `json:"lastTransactionID"`
		Balance                     string        `json:"balance"`
		OpenTradeCount              int           `json:"openTradeCount"`
		OpenPositionCount           int           `json:"openPositionCount"`
		PendingOrderCount           int           `json:"pendingOrderCount"`
		Pl                          string        `json:"pl"`
		ResettablePL                string        `json:"resettablePL"`
		ResettablePLTime            time.Time     `json:"resettablePLTime"`
		Financing                   string        `json:"financing"`
		Commission                  string        `json:"commission"`
		GuaranteedExecutionFees     string        `json:"guaranteedExecutionFees"`
		Orders                      []interface{} `json:"orders"`
		Positions                   []interface{} `json:"positions"`
		Trades                      []interface{} `json:"trades"`
		UnrealizedPL                string        `json:"unrealizedPL"`
		NAV                         string        `json:"NAV"`
		MarginUsed                  string        `json:"marginUsed"`
		MarginAvailable             string        `json:"marginAvailable"`
		PositionValue               string        `json:"positionValue"`
		MarginCloseoutUnrealizedPL  string        `json:"marginCloseoutUnrealizedPL"`
		MarginCloseoutNAV           string        `json:"marginCloseoutNAV"`
		MarginCloseoutMarginUsed    string        `json:"marginCloseoutMarginUsed"`
		MarginCloseoutPositionValue string        `json:"marginCloseoutPositionValue"`
		MarginCloseoutPercent       string        `json:"marginCloseoutPercent"`
		WithdrawalLimit             string        `json:"withdrawalLimit"`
		MarginCallMarginUsed        string        `json:"marginCallMarginUsed"`
		MarginCallPercent           string        `json:"marginCallPercent"`
	} `json:"account"`
	LastTransactionID string `json:"lastTransactionID"`
}

type Provider interface {
	List() (*ListResponse, error)
	Details(accountID string) (*DetailsResponse, error)
}
