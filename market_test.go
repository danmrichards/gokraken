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

	expectedResult := &TimeResponse{
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

func TestMarket_Assets(t *testing.T) {
	cases := []struct {
		name             string
		assetRequest     *AssetsRequest
		expectedResponse *AssetsResponse
		mockResponse     Response
	}{
		{
			name: "no request",
			mockResponse: Response{
				Result: map[Currency]*Asset{
					BCH: {
						AltName:         string(BCH),
						AClass:          "currency",
						Decimals:        10,
						DisplayDecimals: 5,
					},
					DASH: {
						AltName:         string(DASH),
						AClass:          "currency",
						Decimals:        10,
						DisplayDecimals: 5,
					},
					EOS: {
						AltName:         string(EOS),
						AClass:          "currency",
						Decimals:        10,
						DisplayDecimals: 5,
					},
				},
			},
			expectedResponse: &AssetsResponse{
				BCH: &Asset{
					AltName:         string(BCH),
					AClass:          "currency",
					Decimals:        10,
					DisplayDecimals: 5,
				},
				DASH: &Asset{
					AltName:         string(DASH),
					AClass:          "currency",
					Decimals:        10,
					DisplayDecimals: 5,
				},
				EOS: &Asset{
					AltName:         string(EOS),
					AClass:          "currency",
					Decimals:        10,
					DisplayDecimals: 5,
				},
			},
		},
		{
			name: "filtered request",
			mockResponse: Response{
				Result: map[Currency]*Asset{
					BCH: {
						AltName:         string(BCH),
						AClass:          "currency",
						Decimals:        10,
						DisplayDecimals: 5,
					},
					DASH: {
						AltName:         string(DASH),
						AClass:          "currency",
						Decimals:        10,
						DisplayDecimals: 5,
					},
				},
			},
			assetRequest: &AssetsRequest{
				Asset: []Currency{
					BCH,
					DASH,
				},
			},
			expectedResponse: &AssetsResponse{
				BCH: &Asset{
					AltName:         string(BCH),
					AClass:          "currency",
					Decimals:        10,
					DisplayDecimals: 5,
				},
				DASH: &Asset{
					AltName:         string(DASH),
					AClass:          "currency",
					Decimals:        10,
					DisplayDecimals: 5,
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				response, _ := json.Marshal(c.mockResponse)
				w.Write(response)
			}))

			defer ts.Close()

			k := New()
			k.BaseURL = ts.URL

			resp, err := k.Market.Assets(context.Background(), c.assetRequest)
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, resp, t)
		})
	}
}
