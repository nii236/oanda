package oanda

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func (c *Client) ListAccounts() (*Account, error) {
	result := &Account{}
	u, err := url.Parse(c.Host)

	if err != nil {
		return nil, errors.Wrap(err, "Could not parse URL")
	}
	u.Path = "/v3/accounts"
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Could not build request")
	}
	fmt.Println("Request URL:", req.URL)

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)

	if err != nil {
		return nil, errors.Wrap(err, "Could not execute request")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Could not read response")
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.Wrap(err, "Could not unmarshal JSON response")
	}

	return result, nil
}
func (c *Client) AccountDetails()     {}
func (c *Client) AccountSummary()     {}
func (c *Client) AccountInstruments() {}
func (c *Client) ConfigureAccount()   {}
func (c *Client) PollAccountUpdates() {}
