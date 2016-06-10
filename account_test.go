package oanda_test

import (
	"testing"

	"github.com/pkg/errors"
)

func TestListAccounts(t *testing.T) {
	s, c := testTools(200, loadStub("./stubs/account/accounts.json"))
	defer s.Close()
	accounts, err := c.ListAccounts()
	if err != nil {
		t.Error(errors.Cause(err))
		return
	}
	t.Log(accounts)

	got := accounts.AccountProperties[0].ID
	expected := "123-456-7891234-567"

	if got != expected {
		t.Errorf("ERROR: Got %s, expected %s", got, expected)
	}
}
