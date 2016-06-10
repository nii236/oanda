package oanda

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
)

// ListAccounts will get a list of all Accounts authorized for the provided token.
func (c *Client) ListAccounts() (*AccountsContainer, error) {
	result := &AccountsContainer{}
	u, err := url.Parse(c.Host)

	if err != nil {
		return nil, errors.Wrap(err, "Could not parse URL")
	}
	u.Path = "/v3/accounts"
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Could not build request")
	}
	log.Debugln("Request URL:", req.URL)

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

// AccountDetails will get the full details for a single Account that a client has access to. Full pending Order, open Trade and open Position representations are provided.
func (c *Client) AccountDetails(id string) (*Account, error) {
	result := &AccountContainer{}
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse URL")
	}
	u.Path = "/v3/accounts/" + id
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
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
	return &result.Account, nil
}

// AccountSummary will get a summary for a single Account that a client has access to.
func (c *Client) AccountSummary() error {
	return errors.New("Not implemented yet")
}

// AccountInstruments will get the list of tradeable instruments for the given Account.
func (c *Client) AccountInstruments() error {
	return errors.New("Not implemented yet")
}

// ConfigureAccount will set the client-configurable portions of an Account.
func (c *Client) ConfigureAccount() error {
	return errors.New("Not implemented yet")
}

// PollAccountUpdates will poll an Account for its current state and changes since a specified TransactionID.
func (c *Client) PollAccountUpdates() error {
	return errors.New("Not implemented yet")
}
