package main

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/nii236/oanda"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type pairslice []string

var pairs pairslice
var log *logrus.Logger

func init() {
	viper.SetDefault("token", "")
	viper.SetDefault("account", "")
	viper.SetDefault("pairs", []string{"EUR_USD"})

	pflag.StringP("token", "t", "", "Access token from Oanda")
	pflag.StringP("account", "a", "", "Account ID from Oanda")
	pflag.StringSliceP("pairs", "p", []string{"EUR_USD", "AUD_USD"}, "Slice of pairs to subscribe")
	pflag.BoolP("help", "h", false, "Prints the help text")
	viper.BindPFlag("token", pflag.Lookup("token"))
	viper.BindPFlag("account", pflag.Lookup("account"))
	viper.BindPFlag("pairs", pflag.Lookup("pairs"))
	viper.BindPFlag("help", pflag.Lookup("help"))

	pflag.Parse()

	viper.SetEnvPrefix("OANDA")
	viper.BindEnv("token")
	viper.BindEnv("account")
	viper.BindEnv("pairs")

	log = logrus.New()
}

func main() {
	if viper.GetBool("help") {
		pflag.PrintDefaults()
		return
	}
	token := viper.GetString("token")
	account := viper.GetString("account")
	pairs := viper.GetStringSlice("pairs")

	log.Infoln("token:", token)
	log.Infoln("account:", account)
	log.Infoln("pairs:", pairs)

	if token == "" || account == "" {
		log.Error("Need token and account.")
		return
	}

	c, err := oanda.NewFxPracticeClient(token)
	if err != nil {
		log.Errorln(err)
		return
	}

	t := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-t.C:
			log.Infoln("Fetching latest prices...")
			p, err := c.GetPrices(account, pairs, "")
			if err != nil {
				log.Errorln(err)
				t.Stop()
				return
			}
			for _, price := range p.Prices {

				log.Infoln("Instrument:", price.InstrumentName)
				log.Infoln("Status:", price.Status)
				t, err := time.Parse(time.RFC3339Nano, price.Time)
				if err != nil {
					log.Infoln("Could not parse time")
				}
				log.Infoln("Time:", t.Local())
				log.Infoln("Asks:")
				for _, ask := range price.Asks {
					log.Infoln("	Price:", ask.Price, "Liquidity:", ask.Liquidity)
				}
				log.Infoln("Bids:")
				for _, bid := range price.Bids {
					log.Infoln("	Price:", bid.Price, "Liquidity:", bid.Liquidity)
				}
			}

		}

	}
}
