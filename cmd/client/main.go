package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nii236/oanda"

	"github.com/nii236/oanda/account"
	"github.com/nii236/oanda/instrument"
	"github.com/nii236/oanda/pricing"
)

var token string

func init() {
	token = os.Getenv("TOKEN")
}

func main() {
	fmt.Println("Using token:", token)
	ad, err := firstAccount()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Using account ID:", ad)

	pricingService := pricing.New(oanda.APIFXPracticeURL, oanda.StreamFXPracticeURL, ad, token)
	instruments := []string{"EUR_USD", "USD_CAD"}

	fmt.Println("Subscribing to:", strings.Join(instruments, ", "))
	c, err := pricingService.Stream([]string{"EUR_USD", "USD_CAD"}, "")
	if err != nil {
		fmt.Println(err)
		return
	}

	// b := bot.New()

	for {
		select {
		case data := <-c:
			if data.Type == "HEARTBEAT" {
				fmt.Println("Ping? Pong!")
				continue
			}
			fmt.Println(data.Bids)
			fmt.Println(data.Asks)
			// b.OnTick(data.Bids, data.Asks)
		}
	}
}

// firstAccount shows an example use of the account package
func firstAccount() (string, error) {
	accountService := account.New(oanda.APIFXPracticeURL, token)
	a, err := accountService.List()
	if err != nil {
		return "", err
	}
	ad, err := accountService.Details(a.Accounts[0].ID)
	if err != nil {
		return "", err
	}
	return ad.Account.ID, nil
}

// instruments shows an example use of the instrument package
func instruments() {
	// Test with no option
	c := instrument.New(oanda.APIFXPracticeURL, token)
	container, err := c.GetCandles("EUR_GBP")

	// Display
	nbr := len(container.Candles)
	fmt.Println("No options:")
	fmt.Println(0, container.Candles[0])
	fmt.Println(nbr, container.Candles[nbr-1])

	// Test with all options (except 'CandlesTo')
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		panic(err)
	}

	container, err = c.GetCandles(
		"EUR_USD",
		instrument.CandlesAlignTimezone(*loc),
		instrument.CandlesCount(2),
		instrument.CandlesDailyAlign(2),
		instrument.CandlesFrom(time.Now().AddDate(0, -15, 0)),
		instrument.CandlesGranularity("D"),
		instrument.CandlesIncludeFirst(false),
		instrument.CandlesPrice("BA"),
		instrument.CandlesSmooth(true),
		instrument.CandlesWeeklyAlign(time.Monday),
	)
	if err != nil {
		panic(err)
	}

	// Display
	nbr = len(container.Candles)
	fmt.Println("All options:")
	fmt.Println(0, container.Candles[0])
	fmt.Println(nbr, container.Candles[nbr-1])
}
