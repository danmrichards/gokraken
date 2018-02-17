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

	res, err := k.Market.Time(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
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

			res, err := k.Market.Assets(context.Background(), c.assetRequest)
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, res, t)
		})
	}
}

func TestMarket_AssetPairs(t *testing.T) {
	cases := []struct {
		name             string
		infoLevel        AssetPairsInfoLevel
		pairs            []string
		expectedResponse *AssetPairsResponse
		mockResponse     Response
	}{
		{
			name: "no request",
			expectedResponse: &AssetPairsResponse{
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
			expectedResponse: &AssetPairsResponse{
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
			expectedResponse: &AssetPairsResponse{
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
			expectedResponse: &AssetPairsResponse{
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
			expectedResponse: &AssetPairsResponse{
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

			res, err := k.Market.AssetPairs(context.Background(), &AssetPairsRequest{
				Info:  c.infoLevel,
				Pairs: c.pairs,
			})
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, res, t)
		})
	}
}

func TestMarket_Ticker(t *testing.T) {
	cases := []struct {
		name             string
		pairs            []string
		mockResponse     Response
		expectedResponse *TickerResponse
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
			expectedResponse: &TickerResponse{
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

			res, err := k.Market.Ticker(context.Background(), c.pairs...)
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, res, t)
		})
	}
}

func TestMarket_Ohlc(t *testing.T) {
	cases := []struct {
		name             string
		request          *OhlcRequest
		mockResponse     []byte
		expectedResponse *OhlcResponse
	}{
		{
			name: "valid request",
			request: &OhlcRequest{
				Pair: "BCHEUR",
			},
			mockResponse: []byte(`{"error":[],"result":{"BCHEUR":[[1518774960,"1196.0","1196.0","1196.0","1196.0","0.0","0.00000000",0]],"last":1518818040}}`),
			expectedResponse: &OhlcResponse{
				Data: []OhlcData{
					{
						Timestamp: time.Unix(1518774960, 0),
						Open:      1196.0,
						High:      1196.0,
						Low:       1196.0,
						Close:     1196.0,
						Vwap:      0.0,
						Volume:    0.00000000,
						Count:     0,
					},
				},
				Last: 1518818040,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				w.Write(c.mockResponse)
			}))

			defer ts.Close()

			k := New()
			k.BaseURL = ts.URL

			res, err := k.Market.Ohlc(context.Background(), c.request)
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, res, t)
		})
	}
}

func TestMarket_Depth(t *testing.T) {
	cases := []struct {
		name             string
		request          *DepthRequest
		mockResponse     []byte
		expectedResponse *DepthResponse
	}{
		{
			name:         "valid request",
			request:      &DepthRequest{Pair: "BCHEUR"},
			mockResponse: []byte(`{"error":[],"result":{"BCHEUR":{"asks":[["1225.000000","3.729",1518899703]],"bids":[["1222.600000","0.664",1518899718]]}}}`),
			expectedResponse: &DepthResponse{
				"BCHEUR": Depth{
					Asks: []DepthItem{
						{
							Price:     1225.000000,
							Volume:    3.729,
							Timestamp: time.Unix(1518899703, 0),
						},
					},
					Bids: []DepthItem{
						{
							Price:     1222.600000,
							Volume:    0.664,
							Timestamp: time.Unix(1518899718, 0),
						},
					},
				},
			},
		},
		{
			name:         "count request",
			request:      &DepthRequest{Pair: "BCHEUR", Count: 1},
			mockResponse: []byte(`{"error":[],"result":{"BCHEUR":{"asks":[["1230.100000","14.673",1518900219],["1231.300000","0.112",1518900211]],"bids":[["1230.000000","0.486",1518900183],["1229.800000","0.108",1518900204]]}}}`),
			expectedResponse: &DepthResponse{
				"BCHEUR": Depth{
					Asks: []DepthItem{
						{
							Price:     1230.100000,
							Volume:    14.673,
							Timestamp: time.Unix(1518900219, 0),
						},
						{
							Price:     1231.300000,
							Volume:    0.112,
							Timestamp: time.Unix(1518900211, 0),
						},
					},
					Bids: []DepthItem{
						{
							Price:     1230.000000,
							Volume:    0.486,
							Timestamp: time.Unix(1518900183, 0),
						},
						{
							Price:     1229.800000,
							Volume:    0.108,
							Timestamp: time.Unix(1518900204, 0),
						},
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

				w.Write(c.mockResponse)
			}))

			defer ts.Close()

			k := New()
			k.BaseURL = ts.URL

			res, err := k.Market.Depth(context.Background(), c.request)
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, res, t)
		})
	}
}

func TestMarket_Trades(t *testing.T) {
	cases := []struct {
		name             string
		request          *TradesRequest
		mockResponse     []byte
		expectedResponse *TradesResponse
	}{
		{
			name: "valid request",
			request: &TradesRequest{
				Pair: "BCHEUR",
			},
			mockResponse: []byte(`{"error":[],"result":{"BCHEUR":[["700000.000000","0.00050000",1501603433.7669,"s","l",""]],"last":"1501605300157840478"}}`),
			expectedResponse: &TradesResponse{
				Trades: []Trade{
					{
						Price:         700000,
						Volume:        0.00050000,
						Timestamp:     time.Unix(1501603433, 0),
						BuySell:       TradeSell,
						MarketLimit:   TradeLimit,
						Miscellaneous: "",
					},
				},
				Last: 1501605300157840478,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				w.Write(c.mockResponse)
			}))

			defer ts.Close()

			k := New()
			k.BaseURL = ts.URL

			res, err := k.Market.Trades(context.Background(), c.request)
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, res, t)
		})
	}
}

func TestMarket_Spread(t *testing.T) {
	cases := []struct {
		name             string
		request          *SpreadRequest
		mockResponse     []byte
		expectedResponse *SpreadResponse
	}{
		{
			name:         "valid request",
			request:      &SpreadRequest{Pair: "BCHEUR"},
			mockResponse: []byte(`{"error":[],"result":{"BCHEUR":[[1518904771,"1225.600000","1229.200000"]],"last":1518905570}}`),
			expectedResponse: &SpreadResponse{
				Data: []SpreadData{
					{
						Timestamp: time.Unix(1518904771, 0),
						Bid:       1225.6,
						Ask:       1229.2,
					},
				},
				Last: 1518905570,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				w.Write(c.mockResponse)
			}))

			defer ts.Close()

			k := New()
			k.BaseURL = ts.URL

			res, err := k.Market.Spread(context.Background(), &SpreadRequest{Pair: "BCHEUR"})
			if err != nil {
				t.Fatal(err)
			}

			assert(c.expectedResponse, res, t)
		})
	}
}
