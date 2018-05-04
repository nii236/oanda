package pricing

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/nii236/oanda/helpers"
)

// Service implements Accounter
type Service struct {
	*http.Client
	host       string
	streamHost string
	token      string
	accountID  string
}

// New returns a new Service
func New(host, streamHost, accountID, token string) *Service {
	c := &http.Client{
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

	return &Service{c, host, streamHost, token, accountID}
}

// Get will get a list of all Accounts authorized for the provided token.
func (s *Service) Get(instruments []string) (*GetResponse, error) {
	u, err := url.Parse(fmt.Sprintf("%s/v3/accounts/%s/pricing", s.host, s.accountID))
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("instruments", strings.Join(instruments, ","))
	u.RawQuery = q.Encode()

	result := &GetResponse{}

	err = helpers.Get(s.Client, u, s.token, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
