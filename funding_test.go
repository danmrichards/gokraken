package gokraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danmrichards/gokraken/asset"
)

func TestFunding_DepositMethods(t *testing.T) {
	mockResponse := []byte(`{"error":[],"result":{"method": "BACS", "limit": 1.23, "fee": 0.12, "address-setup-fee": true}}`)

	expectedResult := &DepositMethodsResponse{
		Method:          "BACS",
		Limit:           1.23,
		Fee:             0.12,
		AddressSetupFee: true,
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
