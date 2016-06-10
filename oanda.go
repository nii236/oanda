// Package oanda is an implementation of the new V20 REST API from Oanda in Go
package oanda

import (
	"net"
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/pkg/errors"
)

// Client is an instance of the Oanda client
type Client struct {
	Token  string
	Client *http.Client
	Host   string
}

func newClient(env string, token string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout: 10 * time.Second,
				MaxIdleConnsPerHost: 2,
			},
		}
	}
	var host string
	if env == "fxpractice" {
		host = "https://api-fxpractice.oanda.com"
	} else if env == "fxtrade" {
		host = "https://api-fxtrade.oanda.com"
	} else {
		return nil, errors.New("No trading environment specified. Please state fxtrade or fxpractice.")
	}
	return &Client{
		Token:  token,
		Client: httpClient,
		Host:   host,
	}, nil
}

func init() {
	if os.Getenv("DEBUG") != "" {
		log.SetLevel(log.DebugLevel)
	}
}

// NewFxPracticeClient will return a new instance of an Oanda client, ready for practice trading
func NewFxPracticeClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("No token provided")
	}
	return newClient("fxpractice", token, nil)
}

// NewFxClient will return a new instance of an Oanda client, ready for live trading
func NewFxClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("No token provided")
	}
	return newClient("fxtrade", token, nil)
}
