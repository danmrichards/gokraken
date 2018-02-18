package gokraken

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danmrichards/gokraken/asset"
)

func TestUserData_Balance(t *testing.T) {
	mockResponse := Response{
		Result: map[asset.Currency]float64{
			asset.BCH:  1.23,
			asset.DASH: 2.34,
		},
	}

	expectedResult := BalanceResponse{
		asset.BCH:  1.23,
		asset.DASH: 2.34,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(mockResponse)
		w.Write(response)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	res, err := k.UserData.Balance(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}
