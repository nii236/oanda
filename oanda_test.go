package oanda_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/nii236/oanda"
)

func testTools(code int, body string) (*httptest.Server, *oanda.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	httpClient := &http.Client{Transport: transport}
	client := &oanda.Client{
		Client: httpClient,
		Host:   server.URL,
	}

	return server, client
}

func loadStub(filename string) string {
	var err error
	stub := []byte("")
	if len(filename) > 0 {
		stub, err = ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
	}
	return string(stub)
}
