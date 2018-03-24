package gokraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danmrichards/gokraken/asset"
	"golang.org/x/text/currency"
)

func TestFunding_DepositMethods(t *testing.T) {
	mockResponse := []byte(`{"error":[],"result":{"1234": {"method": "BACS", "limit": 1.23, "fee": 0.12, "address-setup-fee": true}}}`)

	expectedResult := DepositMethodsResponse{
		"1234": {
			Method:          "BACS",
			Limit:           1.23,
			Fee:             0.12,
			AddressSetupFee: true,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	res, err := k.Funding.DepositMethods(context.Background(), AssetCurrency, asset.BCH)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestFunding_DepositAddresses(t *testing.T) {
	mockResponse := []byte(`{"error":[],"result":{"1234":{"address":"1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX","expiretm":0,"new":false}}}`)

	expectedResult := DepositAddressesResponse{
		"1234": {
			Address:    "1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX",
			ExpireTime: 0,
			New:        false,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	res, err := k.Funding.DepositAddresses(context.Background(), AssetCurrency, asset.BCH, "test", false)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestFunding_DepositStatus(t *testing.T) {
	mockResponse := []byte(`{"error":[],"result":{"method":"test","aclass":"currency","asset":"GBP","refid":"1234","txid":"4321","info":"","amount":1.23,"fee":0.12,"time":1521890577,"status":"ok","status-prop":{"onhold":"false"}}}`)

	expectedResult := &DepositStatusResponse{
		Method: "test",
		AClass: string(AssetCurrency),
		Asset:  currency.GBP.String(),
		RefID:  "1234",
		TxID:   "4321",
		Amount: 1.23,
		Fee:    0.12,
		Time:   1521890577,
		Status: "ok",
		StatusProp: DepositStatusProp{
			OnHold: "false",
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	res, err := k.Funding.DepositStatus(context.Background(), AssetCurrency, asset.BCH, "test")
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}
