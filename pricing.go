package oanda

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// GetPrices will get pricing information for a specified list of Instruments within an Account.
func (c *Client) GetPrices(accountID string, instruments []string, since string) (*PricesContainer, error) {
	var result PricesContainer
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse URL")
	}
	u.Path = "/v3/accounts/" + accountID + "/pricing"
	q := &url.Values{}
	q.Set("instruments", strings.Join(instruments, ","))
	u.RawQuery = q.Encode()
	var req *http.Request
	if since != "" {
		// Have to append raw string for since manually, since url package escapes the colons
		req, err = http.NewRequest("GET", u.String()+"&since="+since, nil)
	} else {
		req, err = http.NewRequest("GET", u.String(), nil)
	}
	if err != nil {
		return nil, errors.Wrap(err, "Could not build request")
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Client.Do(req)

	if err != nil {
		return nil, errors.Wrap(err, "Could not execute request")
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Could not read from resp body")
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.Wrap(err, "Could not unmarshal JSON")
	}

	return &result, nil
}
