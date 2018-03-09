package gokraken

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/danmrichards/gokraken/asset"
	"github.com/danmrichards/gokraken/pairs"
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

func TestUserData_TradeBalance(t *testing.T) {
	mockResponse := []byte(`{"result": {"eb":1.23,"tb":1.23,"m":1.23,"n":1.23,"c":1.23,"v":1.23,"e":1.23,"mf":1.23,"ml":1.23}}`)

	expectedResult := &TradeBalanceResponse{
		EquivalentBalance: 1.23,
		TradeBalance:      1.23,
		MarginAmount:      1.23,
		UnrealizedNet:     1.23,
		Cost:              1.23,
		Valuation:         1.23,
		Equity:            1.23,
		FreeMargin:        1.23,
		MarginLevel:       1.23,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	res, err := k.UserData.TradeBalance(context.Background(), AssetCurrency, asset.ZUSD)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_OpenOrders(t *testing.T) {
	mockResponse := []byte(`{"result": {"open":{"1234":{"refid":"4321","userref":123456,"status":"open","opentm":1520287055,"starttm":1520287055,"expiretm":1520287055,"descr":{"pair":45,"close":"","leverage":"","order":"","ordertype":"market","price":"","price2":"","type":""},"vol":"","vol_exec":"1.23","cost":"1.23","fee":"0","price":"0","stopprice.string":0,"limitprice":"0","misc":"","oflags":"","closetm":1520287055,"reason":""}},"count":1}}`)

	expectedResult := &OpenOrdersResponse{
		Open: map[string]Order{
			"1234": {
				ReferenceID: "4321",
				UserRef:     123456,
				Status:      "open",
				OpenTime:    1520287055,
				StartTime:   1520287055,
				ExpireTime:  1520287055,
				Description: OrderDescription{
					AssetPair: pairs.BCHEUR,
					OrderType: OrderTypeMarket,
				},
				VolumeExecuted: 1.23,
				Cost:           1.23,
				CloseTime:      1520287055,
			},
		},
		Count: 1,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	res, err := k.UserData.OpenOrders(context.Background(), false, 0)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_ClosedOrders(t *testing.T) {
	mockResponse := []byte(`{"result": {"closed":{"1234":{"refid":"4321","userref":123456,"status":"open","opentm":1520287055,"starttm":1520287055,"expiretm":1520287055,"descr":{"pair":45,"close":"","leverage":"","order":"","ordertype":"market","price":"","price2":"","type":""},"vol":"","vol_exec":"1.23","cost":"1.23","fee":"0","price":"0","stopprice.string":0,"limitprice":"0","misc":"","oflags":"","closetm":1520287055,"reason":""}},"count":1}}`)

	expectedResult := &ClosedOrdersResponse{
		Closed: map[string]Order{
			"1234": {
				ReferenceID: "4321",
				UserRef:     123456,
				Status:      "open",
				OpenTime:    1520287055,
				StartTime:   1520287055,
				ExpireTime:  1520287055,
				Description: OrderDescription{
					AssetPair: pairs.BCHEUR,
					OrderType: OrderTypeMarket,
				},
				VolumeExecuted: 1.23,
				Cost:           1.23,
				CloseTime:      1520287055,
			},
		},
		Count: 1,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	now := time.Now()
	req := ClosedOrdersRequest{
		Trades:    true,
		UserRef:   123456,
		Start:     &now,
		End:       &now,
		Ofs:       1,
		CloseTime: OrderCloseTimeClose,
	}

	res, err := k.UserData.ClosedOrders(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_QueryOrders(t *testing.T) {
	mockResponse := []byte(`{"result":{"1234":{"refid":"4321","userref":123456,"status":"open","opentm":1520287055,"starttm":1520287055,"expiretm":1520287055,"descr":{"pair":45,"close":"","leverage":"","order":"","ordertype":"market","price":"","price2":"","type":""},"vol":"","vol_exec":"1.23","cost":"1.23","fee":"0","price":"0","stopprice.string":0,"limitprice":"0","misc":"","oflags":"","closetm":1520287055,"reason":""}}}`)

	expectedResult := &QueryOrdersResponse{
		"1234": {
			ReferenceID: "4321",
			UserRef:     123456,
			Status:      "open",
			OpenTime:    1520287055,
			StartTime:   1520287055,
			ExpireTime:  1520287055,
			Description: OrderDescription{
				AssetPair: pairs.BCHEUR,
				OrderType: OrderTypeMarket,
			},
			VolumeExecuted: 1.23,
			Cost:           1.23,
			CloseTime:      1520287055,
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

	res, err := k.UserData.QueryOrders(context.Background(), true, 123456, 654321)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_TradesHistory(t *testing.T) {
	mockResponse := []byte(`{"result":{"trades":{"1234":{"ordertxid":"4321","pair":"BCHEUR","time":1520633741,"type":"buy","ordertype":"market","price":1.23,"cost":1.23,"fee":1.23,"vol":1.23,"margin":1.23,"misc":"foo,bar,baz"}},"count":1}}`)

	expectedResult := &TradesHistoryResponse{
		Trades: map[string]UserTrade{
			"1234": {
				OrderTxid: "4321",
				Pair:      pairs.BCHEUR.String(),
				Time:      1520633741,
				Type:      TradeBuy,
				OrderType: OrderTypeMarket,
				Price:     1.23,
				Cost:      1.23,
				Fee:       1.23,
				Vol:       1.23,
				Margin:    1.23,
				Misc:      "foo,bar,baz",
			},
		},
		Count: 1,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	now := time.Now()
	req := TradesHistoryRequest{
		Type:   TradeTypeAll,
		Trades: true,
		Start:  &now,
		End:    &now,
		Ofs:    1,
	}
	res, err := k.UserData.TradesHistory(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_QueryTrades(t *testing.T) {
	mockResponse := []byte(`{"result":{"1234":{"ordertxid":"4321","pair":"BCHEUR","time":1520633741,"type":"buy","ordertype":"market","price":1.23,"cost":1.23,"fee":1.23,"vol":1.23,"margin":1.23,"misc":"foo,bar,baz"}}}`)

	expectedResult := &QueryTradesResponse{
		"1234": {
			OrderTxid: "4321",
			Pair:      pairs.BCHEUR.String(),
			Time:      1520633741,
			Type:      TradeBuy,
			OrderType: OrderTypeMarket,
			Price:     1.23,
			Cost:      1.23,
			Fee:       1.23,
			Vol:       1.23,
			Margin:    1.23,
			Misc:      "foo,bar,baz",
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

	res, err := k.UserData.QueryTrades(context.Background(), true, 1234)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}
