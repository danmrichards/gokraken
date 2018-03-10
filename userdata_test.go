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

func TestUserData_OpenPositions(t *testing.T) {
	mockResponse := []byte(`{"result":{"1234":{"ordertxid":"4321","pair":"BCHEUR","time":1520633741,"type":"buy","ordertype":"market","cost":1.23,"fee":1.23,"vol":1.23,"vol_closed":1.23,"margin":1.23,"value":1.23,"net":1.23,"misc":"foo,bar,baz","oflags":"qux,quux","viqc":1.23}}}`)

	expectedResult := OpenPositionsResponse{
		"1234": {
			OrderTxid: "4321",
			Pair:      pairs.BCHEUR.String(),
			Time:      1520633741,
			Type:      string(TradeBuy),
			OrderType: string(OrderTypeMarket),
			Cost:      1.23,
			Fee:       1.23,
			Vol:       1.23,
			VolClosed: 1.23,
			Margin:    1.23,
			Value:     1.23,
			Net:       1.23,
			Misc:      "foo,bar,baz",
			OFlags:    "qux,quux",
			Viqc:      1.23,
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

	res, err := k.UserData.OpenPositions(context.Background(), true, 1234)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_Ledgers(t *testing.T) {
	mockResponse := []byte(`{"result":{"1234":{"refid":"4321","time":1520633741,"type":"all","aclass":"currency","asset":"DASH","amount":1.23,"fee":1.23,"balance":1.23}}}`)

	expectedResult := LedgersResponse{
		"1234": {
			Refid:   "4321",
			Time:    1520633741,
			Type:    string(LedgerTypeAll),
			Aclass:  string(AssetCurrency),
			Asset:   asset.DASH.String(),
			Amount:  1.23,
			Fee:     1.23,
			Balance: 1.23,
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

	now := time.Now()
	req := LedgersRequest{
		Aclass: AssetCurrency,
		Assets: []asset.Currency{
			asset.DASH,
		},
		Type:  LedgerTypeAll,
		Start: &now,
		End:   &now,
		Ofs:   1,
	}
	res, err := k.UserData.Ledgers(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_QueryLedgers(t *testing.T) {
	mockResponse := []byte(`{"result":{"1234":{"refid":"4321","time":1520633741,"type":"all","aclass":"currency","asset":"DASH","amount":1.23,"fee":1.23,"balance":1.23}}}`)

	expectedResult := LedgersResponse{
		"1234": {
			Refid:   "4321",
			Time:    1520633741,
			Type:    string(LedgerTypeAll),
			Aclass:  string(AssetCurrency),
			Asset:   asset.DASH.String(),
			Amount:  1.23,
			Fee:     1.23,
			Balance: 1.23,
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

	res, err := k.UserData.QueryLedgers(context.Background(), 1234)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}

func TestUserData_TradeVolume(t *testing.T) {
	mockResponse := []byte(`{"error":[],"result":{"currency":"ZUSD","fees":{"BCHEUR":{"fee":"0.2600","maxfee":"0.2600","minfee":"0.1000","nextfee":"0.2400","nextvolume":"50000.0000","tiervolume":"0.0000"}},"fees_maker":{"BCHEUR":{"fee":"0.1600","maxfee":"0.1600","minfee":"0.0000","nextfee":"0.1400","nextvolume":"50000.0000","tiervolume":"0.0000"}},"volume":"0.0000"}}`)

	expectedResult := &TradeVolumeResponse{
		Currency: asset.ZUSD.String(),
		Fees: map[string]Fee{
			pairs.BCHEUR.String(): {
				Fee:        "0.2600",
				Maxfee:     "0.2600",
				Minfee:     "0.1000",
				Nextfee:    "0.2400",
				Nextvolume: "50000.0000",
				Tiervolume: "0.0000",
			},
		},
		FeesMaker: map[string]Fee{
			pairs.BCHEUR.String(): {
				Fee:        "0.1600",
				Maxfee:     "0.1600",
				Minfee:     "0.0000",
				Nextfee:    "0.1400",
				Nextvolume: "50000.0000",
				Tiervolume: "0.0000",
			},
		},
		Volume: "0.0000",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(mockResponse)
	}))

	defer ts.Close()

	k := NewWithAuth("api_key", "cHJpdmF0ZV9rZXk=")
	k.BaseURL = ts.URL

	res, err := k.UserData.TradeVolume(context.Background(), true, pairs.BCHEUR)
	if err != nil {
		t.Fatal(err)
	}

	assert(expectedResult, res, t)
}
