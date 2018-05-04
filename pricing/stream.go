package pricing

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Stream will get pricing information for a specified list of Instruments within an Account.
func (s *Service) Stream(instruments []string, since string) (chan *StreamResponse, error) {

	u, err := url.Parse(s.streamHost)
	if err != nil {
		return nil, fmt.Errorf("Could not parse URL: %s", err)
	}
	u.Path = fmt.Sprintf("/v3/accounts/%s/pricing/stream", s.accountID)
	q := &url.Values{}
	q.Set("instruments", strings.Join(instruments, ","))
	u.RawQuery = q.Encode()
	fmt.Println("Fetching prices from:", u.String())
	var req *http.Request
	if since != "" {
		// Have to append raw string for since manually, since url package escapes the colons
		req, err = http.NewRequest("GET", u.String()+"&since="+since, nil)
	} else {
		req, err = http.NewRequest("GET", u.String(), nil)
	}
	if err != nil {
		return nil, fmt.Errorf("Could not build request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer "+s.token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Could not execute request: %s", err)
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Non 200 response: " + resp.Status)
	}

	c := make(chan *StreamResponse)
	scn := bufio.NewScanner(resp.Body)
	go func(c chan *StreamResponse, scan *bufio.Scanner) {
		defer resp.Body.Close()
		for scan.Scan() {
			if scan.Err() == io.EOF {
				break
			}

			if scan.Err() != nil {
				break
			}

			b := scan.Bytes()
			target := &StreamResponse{}
			err = json.Unmarshal(b, &target)
			if err != nil {
				fmt.Println(err)
				continue
			}
			c <- target
		}
	}(c, scn)
	return c, nil
}
