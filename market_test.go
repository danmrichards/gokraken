package gokraken

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMarket_Time(t *testing.T) {
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
		expectedResponse AssetsResponse
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
			expectedResponse: AssetsResponse{
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
			expectedResponse: AssetsResponse{
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

func TestMarket_AssetPairs(t *testing.T) {
	cases := []struct {
		name             string
		infoLevel        AssetPairsInfoLevel
		pairs            []string
		expectedResponse AssetPairsResponse
		mockResponse     Response
	}{
		{
			name: "no request",
			expectedResponse: AssetPairsResponse{
				"BCHEUR": &AssetPair{
					Altname:    "BCHEUR",
					AclassBase: "currency",
					Base:       "BCH",
				},
				"BCHUSD": &AssetPair{
					Altname:    "BCHUSD",
					AclassBase: "currency",
					Base:       "BCH",
				},
			},
			mockResponse: Response{
				Result: map[string]*AssetPair{
					"BCHEUR": {
						Altname:    "BCHEUR",
						AclassBase: "currency",
						Base:       "BCH",
					},
					"BCHUSD": {
						Altname:    "BCHUSD",
						AclassBase: "currency",
						Base:       "BCH",
					},
				},
			},
		},
		{
			name:      "leverage request",
			infoLevel: AssetPairsLeverage,
			expectedResponse: AssetPairsResponse{
				"BCHEUR": &AssetPair{
					LeverageBuy:  []float64{1.23},
					LeverageSell: []float64{2.34},
				},
			},
			mockResponse: Response{
				Result: map[string]*AssetPair{
					"BCHEUR": {
						LeverageBuy:  []float64{1.23},
						LeverageSell: []float64{2.34},
					},
				},
			},
		},
		{
			name:      "fees request",
			infoLevel: AssetPairsFees,
			expectedResponse: AssetPairsResponse{
				"BCHEUR": &AssetPair{
					Fees: [][]float64{
						{1.23, 2.34},
					},
				},
			},
			mockResponse: Response{
				Result: map[string]*AssetPair{
					"BCHEUR": {
						Fees: [][]float64{
							{1.23, 2.34},
						},
					},
				},
			},
		},
		{
			name:      "margin request",
			infoLevel: AssetPairsMargin,
			expectedResponse: AssetPairsResponse{
				"BCHEUR": &AssetPair{
					MarginCall: 10,
					MarginStop: 20,
				},
			},
			mockResponse: Response{
				Result: map[string]*AssetPair{
					"BCHEUR": {
						MarginCall: 10,
						MarginStop: 20,
					},
				},
			},
		},
		{
			name:  "pair filtered request",
			pairs: []string{"BCHUSD"},
			expectedResponse: AssetPairsResponse{
				"BCHUSD": &AssetPair{
					MarginCall: 10,
					MarginStop: 20,
				},
			},
			mockResponse: Response{
				Result: map[string]*AssetPair{
					"BCHUSD": {
						MarginCall: 10,
						MarginStop: 20,
					},
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

			resp, err := k.Market.AssetPairs(context.Background(), &AssetPairsRequest{
				Info:  c.infoLevel,
				Pairs: c.pairs,
			})
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, resp, t)
		})
	}
}

func TestMarket_Ticker(t *testing.T) {
	cases := []struct {
		name             string
		pairs            []string
		mockResponse     Response
		expectedResponse TickerResponse
	}{
		{
			name:  "valid request",
			pairs: []string{"BCHEUR", "BCHUSD"},
			mockResponse: Response{
				Result: map[string]TickerInfo{
					"BCHEUR": {
						A: []string{
							"804.900000",
							"1",
							"1.000",
						},
						B: []string{
							"802.100000",
							"1",
							"1.000",
						},
						C: []string{
							"805.000000",
							"0.09409409",
						},
						V: []string{
							"6285.91000112",
							"6402.41926847",
						},
						P: []string{
							"790.741428",
							"790.497060",
						},
						T: []int{
							12672,
							12902,
						},
						L: []string{
							"718.200000",
							"718.200000",
						},
						H: []string{
							"850.600000",
							"850.600000",
						},
						O: "774.800000",
					},
					"BCHUSD": {
						A: []string{
							"804.900000",
							"1",
							"1.000",
						},
						B: []string{
							"802.100000",
							"1",
							"1.000",
						},
						C: []string{
							"805.000000",
							"0.09409409",
						},
						V: []string{
							"6285.91000112",
							"6402.41926847",
						},
						P: []string{
							"790.741428",
							"790.497060",
						},
						T: []int{
							12672,
							12902,
						},
						L: []string{
							"718.200000",
							"718.200000",
						},
						H: []string{
							"850.600000",
							"850.600000",
						},
						O: "774.800000",
					},
				},
			},
			expectedResponse: TickerResponse{
				"BCHEUR": {
					A: []string{
						"804.900000",
						"1",
						"1.000",
					},
					B: []string{
						"802.100000",
						"1",
						"1.000",
					},
					C: []string{
						"805.000000",
						"0.09409409",
					},
					V: []string{
						"6285.91000112",
						"6402.41926847",
					},
					P: []string{
						"790.741428",
						"790.497060",
					},
					T: []int{
						12672,
						12902,
					},
					L: []string{
						"718.200000",
						"718.200000",
					},
					H: []string{
						"850.600000",
						"850.600000",
					},
					O: "774.800000",
				},
				"BCHUSD": {
					A: []string{
						"804.900000",
						"1",
						"1.000",
					},
					B: []string{
						"802.100000",
						"1",
						"1.000",
					},
					C: []string{
						"805.000000",
						"0.09409409",
					},
					V: []string{
						"6285.91000112",
						"6402.41926847",
					},
					P: []string{
						"790.741428",
						"790.497060",
					},
					T: []int{
						12672,
						12902,
					},
					L: []string{
						"718.200000",
						"718.200000",
					},
					H: []string{
						"850.600000",
						"850.600000",
					},
					O: "774.800000",
				},
			},
		},
		{
			name: "empty pairs",
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

			resp, err := k.Market.Ticker(context.Background(), c.pairs...)
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, resp, t)
		})
	}
}
