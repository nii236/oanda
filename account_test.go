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
	got := accounts.AccountProperties[0].ID
	expected := "101-001-100000-001"

	if got != expected {
		t.Errorf("ERROR: Got %s, expected %s", got, expected)
	}
}

func TestGetAccount(t *testing.T) {
	s, c := testTools(200, loadStub("./stubs/account/account.json"))
	defer s.Close()
	account, err := c.AccountDetails("101-001-100000-001")
	if err != nil {
		t.Fatal(errors.Cause(err))
		return
	}
	t.Log(account)
	got := account.ID
	expected := "101-001-100000-001"

	if got != expected {
		t.Errorf("ERROR: Got %s, expected %s", got, expected)
	}
}
