package main

import (
	"fmt"
	"time"

	"github.com/nii236/oanda"
)

var token = "XXX-YYY"

func main() {
	// Get client
	c, err := oanda.NewFxPracticeClient(token)
	if err != nil {
		panic(err)
	}

	// Test with no option
	container, err := c.GetCandles("EUR_GBP")

	// Display
	nbr := len(container.Candles)
	fmt.Println("No options:")
	fmt.Println(0, container.Candles[0])
	fmt.Println(nbr, container.Candles[nbr-1])

	// Test with all options
	loc, _ := time.LoadLocation("Europe/Paris")
	container, err = c.GetCandles(
		"EUR_USD",
		oanda.CandlesAlignTimezone(*loc),
		oanda.CandlesCount(2),
		oanda.CandlesDailyAlign(2),
		oanda.CandlesFrom(time.Now().AddDate(0, -15, 0)),
		oanda.CandlesGranularity("D"),
		oanda.CandlesIncludeFirst(false),
		oanda.CandlesPrice("BA"),
		oanda.CandlesSmooth(true),
		//oanda.CandlesTo(time.Now().AddDate(0, -1, 0)), -> You can't set From and To at the same time
		oanda.CandlesWeeklyAlign(time.Monday),
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
