package account

import (
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/nii236/oanda/helpers"
	"github.com/pkg/errors"
)

const listURL = "/v3/accounts"
const detailsURL = "/v3/accounts/"

// Service implements Accounter
type Service struct {
	*http.Client
	host  string
	token string
}

// New returns a new Service
func New(host, token string) *Service {
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

	return &Service{c, host, token}
}

// List will get a list of all Accounts authorized for the provided token.
func (s *Service) List() (*ListResponse, error) {
	u, err := url.Parse(s.host + listURL)
	if err != nil {
		return nil, err
	}

	result := &ListResponse{}

	err = helpers.Get(s.Client, u, s.token, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Details will get the full details for a single Details that a client has access to. Full pending Order, open Trade and open Position representations are provided.
func (s *Service) Details(accountID string) (*DetailsResponse, error) {
	u, err := url.Parse(s.host + detailsURL + accountID)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse URL")
	}
	result := &DetailsResponse{}
	err = helpers.Get(s.Client, u, s.token, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
