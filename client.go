package oanda

import (
	"github.com/nii236/oanda/account"
	"github.com/nii236/oanda/instrument"
	"github.com/nii236/oanda/pricing"
)

// Client has all the methods for Oanda
type Client interface {
	account.Provider
	instrument.Provider
	pricing.Provider
}
