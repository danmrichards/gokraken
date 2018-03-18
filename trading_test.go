package gokraken

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danmrichards/gokraken/pairs"
)

func TestUserData_AddOrder(t *testing.T) {
	mockResponse := []byte(`{"error":[],"result":{"descr":{"pair":0,"close":"4321","leverage":"","order":"1234","ordertype":"","price":"","price2":"","type":""},"txid":["2345","3456"]}}`)

	expectedResult := &AddOrderResponse{
		Description: OrderDescription{
			Order: "1234",
			Close: "4321",
		},
		TxIDs: []string{"2345", "3456"},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	order := UserOrder{
		Pair:      pairs.BCHEUR,
		Type:      TradeSell,
		OrderType: OrderTypeMarket,
		Volume:    1.23,
	}

	res, err := k.UserData.AddOrder(context.Background(), order)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_CancelOrder(t *testing.T) {
	mockResponse := []byte(`{"error":[],"result":{"count": 1, "pending": true}}`)

	expectedResult := &CancelOrderResponse{
		Count:   1,
		Pending: true,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	res, err := k.UserData.CancelOrder(context.Background(), 1234)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}
