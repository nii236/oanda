package helpers

import (
	"encoding/json"
	"net/http/httputil"

	"github.com/nii236/oanda/errors"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get will get and unmarshal the response
func Get(c *http.Client, u *url.URL, token string, target interface{}) error {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return fmt.Errorf("could not build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("could not execute request: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// b, _ := ioutil.ReadAll(resp.Body)
		b, _ := httputil.DumpResponse(resp, true)
		return errors.New("non 200 status returned", resp.StatusCode, string(b))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, target)
}
