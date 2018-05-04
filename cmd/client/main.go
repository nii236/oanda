package main

import (
	"fmt"

	"github.com/nii236/oanda"

	"github.com/nii236/oanda/account"
	"github.com/nii236/oanda/pricing"
)

func main() {
	token := "EXAMPLE"
	accountService := account.New(oanda.APIFXPracticeURL, token)
	a, _ := accountService.List()
	fmt.Println(a)
	ad, _ := accountService.Details(a.Accounts[0].ID)
	fmt.Println(ad)

	pricingService := pricing.New(oanda.APIFXPracticeURL, oanda.StreamFXPracticeURL, ad.Account.ID, token)
	prices, _ := pricingService.Get([]string{"EUR_USD", "USD_CAD"})

	fmt.Println(prices)

	c, err := pricingService.Stream([]string{"EUR_USD", "USD_CAD"}, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case data := <-c:
			fmt.Println(data)
		}
	}
}
