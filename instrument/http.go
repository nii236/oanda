package instrument

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	candlesTimeFormat = "2006-01-02T15:04:05Z"
)

func (c *Client) GetCandles(instrument string, args ...CandlesArg) (*CandlesContainer, error) {
	var result CandlesContainer

	// Set url base
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, errors.Wrap(err, "Could not parse URL")
	}
	u.Path = "/v3/instruments/" + instrument + "/candles"

	// Set values
	q := &url.Values{}
	for _, a := range args {
		a.applyCandlesArg(q)
	}
	u.RawQuery = q.Encode()

	// Form a new http request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Could not build request")
	}

	// Set header
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Could not execute request")
	}
	defer resp.Body.Close()

	// Check error code
	if resp.StatusCode != 200 {
		return nil, errors.New("Non 200 response: " + resp.Status)
	}

	// Read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Could not read from resp body")
	}

	// Parse response
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.Wrap(err, "Could not unmarshal JSON")
	}

	return &result, nil
}

// Arguments

type CandlesArg interface {
	applyCandlesArg(*url.Values)
}

type (
	CandlesPrice         string
	CandlesGranularity   string
	CandlesCount         int
	CandlesFrom          time.Time
	CandlesTo            time.Time
	CandlesSmooth        bool
	CandlesIncludeFirst  bool
	CandlesDailyAlign    int
	CandlesAlignTimezone time.Location
	CandlesWeeklyAlign   time.Weekday
)

// Private

func (cp CandlesPrice) applyCandlesArg(v *url.Values) {
	v.Set("price", string(cp))
}

func (cg CandlesGranularity) applyCandlesArg(v *url.Values) {
	v.Set("granularity", string(cg))
}

func (cc CandlesCount) applyCandlesArg(v *url.Values) {
	v.Set("count", strconv.Itoa(int(cc)))
}

func (cf CandlesFrom) applyCandlesArg(v *url.Values) {
	tf := time.Time(cf).Format(candlesTimeFormat)
	v.Set("from", tf)
}

func (ct CandlesTo) applyCandlesArg(v *url.Values) {
	tf := time.Time(ct).Format(candlesTimeFormat)
	v.Set("to", tf)
}

func (cs CandlesSmooth) applyCandlesArg(v *url.Values) {
	v.Set("smooth", strconv.FormatBool(bool(cs)))
}

func (cif CandlesIncludeFirst) applyCandlesArg(v *url.Values) {
	v.Set("includeFirst", strconv.FormatBool(bool(cif)))
}

func (cda CandlesDailyAlign) applyCandlesArg(v *url.Values) {
	v.Set("dailyAlignment", strconv.Itoa(int(cda)))
}

func (cat CandlesAlignTimezone) applyCandlesArg(v *url.Values) {
	loc := time.Location(cat)
	v.Set("alignmentTimezone", loc.String())
}

func (cwa CandlesWeeklyAlign) applyCandlesArg(v *url.Values) {
	v.Set("weeklyAlignment", time.Weekday(cwa).String())
}
