package gokraken

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPublicMarket_Time(t *testing.T) {
	mockResponse := Response{
		Result: map[string]interface{}{
			"unixtime": time.Now().Unix(),
			"rfc1123":  time.Now().Format(time.RFC1123),
		},
	}

	expectedResult := TimeResponse{
		UnixTime: time.Now().Unix(),
		Rfc1123:  time.Now().Format(time.RFC1123),
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(mockResponse)
		w.Write(response)
	}))

	defer ts.Close()

	k := New()
	k.BaseURL = ts.URL

	resp, err := k.Market.Time(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, resp, t)
}
