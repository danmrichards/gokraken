package gokraken

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/danmrichards/gokraken/asset"
	"github.com/danmrichards/gokraken/pairs"
)

// UserData is responsible for communicating with all the private user data
// endpoints on the Kraken API.
type UserData struct {
	Client *Kraken
}

// Balance returns an array of asset names and balance amount.
// https://www.kraken.com/en-gb/help/api#get-account-balance
func (u *UserData) Balance(ctx context.Context) (res BalanceResponse, err error) {
	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, BalanceResource, nil)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// TradeBalance returns an array of trade balance information.
// https://www.kraken.com/en-gb/help/api#get-trade-balance
func (u *UserData) TradeBalance(ctx context.Context, assetClass AssetsClass, base asset.Currency) (res *TradeBalanceResponse, err error) {
	body := url.Values{
		"aclass": []string{string(assetClass)},
		"asset":  []string{base.String()},
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, TradeBalanceResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// OpenOrders returns an order info in open array with txid as the key.
// https://www.kraken.com/help/api#get-open-orders
func (u *UserData) OpenOrders(ctx context.Context, trades bool, userRef int64) (res *OpenOrdersResponse, err error) {
	body := url.Values{
		"trades": []string{"false"},
	}

	if trades {
		body.Set("trades", "true")
	}

	if userRef != 0 {
		body.Add("userref", strconv.FormatInt(userRef, 10))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, OpenOrdersResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// ClosedOrders returns an order info array.
// https://www.kraken.com/en-gb/help/api#get-closed-orders
func (u *UserData) ClosedOrders(ctx context.Context, closedReq ClosedOrdersRequest) (res *ClosedOrdersResponse, err error) {
	body := url.Values{
		"trades":    []string{"false"},
		"closetime": []string{string(OrderCloseTimeBoth)},
	}

	if closedReq.Trades {
		body.Set("trades", "true")
	}

	if closedReq.UserRef != 0 {
		body.Add("userref", strconv.FormatInt(closedReq.UserRef, 10))
	}

	if closedReq.Start != nil {
		body.Add("start", strconv.FormatInt(closedReq.Start.Unix(), 10))
	}

	if closedReq.End != nil {
		body.Add("end", strconv.FormatInt(closedReq.End.Unix(), 10))
	}

	if closedReq.Ofs != 0 {
		body.Add("ofs", strconv.Itoa(closedReq.Ofs))
	}

	if string(closedReq.CloseTime) != "" {
		body.Set("closetime", string(closedReq.CloseTime))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, ClosedOrdersResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// QueryOrders returns an associative array of order info.
// https://www.kraken.com/en-gb/help/api#query-orders-info
func (u *UserData) QueryOrders(ctx context.Context, trades bool, userRef int64, txids ...int64) (res *QueryOrdersResponse, err error) {
	body := url.Values{
		"trades": []string{"false"},
	}

	if trades {
		body.Set("trades", "true")
	}

	if len(txids) > 0 {
		txidStrings := make([]string, len(txids))
		for i := range txids {
			txidStrings[i] = strconv.FormatInt(txids[i], 10)
		}

		body.Add("txid", strings.Join(txidStrings, ","))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, QueryOrdersResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// TradesHistory returns an array of trade info.
// https://www.kraken.com/en-gb/help/api#get-trades-history
func (u *UserData) TradesHistory(ctx context.Context, tradesReq TradesHistoryRequest) (res *TradesHistoryResponse, err error) {
	body := url.Values{
		"type":   []string{string(tradesReq.Type)},
		"trades": []string{"false"},
	}

	if string(tradesReq.Type) != "" {
		body.Set("type", string(tradesReq.Type))
	}

	if tradesReq.Trades {
		body.Set("trades", "true")
	}

	if tradesReq.Start != nil {
		body.Add("start", strconv.FormatInt(tradesReq.Start.Unix(), 10))
	}

	if tradesReq.End != nil {
		body.Add("end", strconv.FormatInt(tradesReq.End.Unix(), 10))
	}

	if tradesReq.Ofs != 0 {
		body.Add("ofs", strconv.Itoa(tradesReq.Ofs))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, TradesHistoryResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// QueryTrades returns an associative array of trade info.
// https://www.kraken.com/en-gb/help/api#query-trades-info
func (u *UserData) QueryTrades(ctx context.Context, trades bool, txids ...int64) (res *QueryTradesResponse, err error) {
	body := url.Values{
		"trades": []string{"false"},
	}

	if trades {
		body.Set("trades", "true")
	}

	if len(txids) > 0 {
		txidStrings := make([]string, len(txids))
		for i := range txids {
			txidStrings[i] = strconv.FormatInt(txids[i], 10)
		}

		body.Add("txid", strings.Join(txidStrings, ","))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, TradesInfoResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// OpenPositions returns an associative array of open positions info.
// https://www.kraken.com/en-gb/help/api#get-open-positions
func (u *UserData) OpenPositions(ctx context.Context, doCalcs bool, txids ...int64) (res OpenPositionsResponse, err error) {
	body := url.Values{
		"docalcs": []string{"false"},
	}

	if doCalcs {
		body.Set("docalcs", "true")
	}

	if len(txids) > 0 {
		txidStrings := make([]string, len(txids))
		for i := range txids {
			txidStrings[i] = strconv.FormatInt(txids[i], 10)
		}

		body.Add("txid", strings.Join(txidStrings, ","))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, OpenPositionsResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// Ledgers returns an associative array of ledgers info.
// https://www.kraken.com/en-gb/help/api#get-ledgers-info
func (u *UserData) Ledgers(ctx context.Context, ledgersReq LedgersRequest) (res LedgersResponse, err error) {
	body := url.Values{
		"aclass": []string{string(AssetCurrency)},
	}

	if string(ledgersReq.Aclass) != "" {
		body.Set("aclass", string(ledgersReq.Aclass))
	}

	if len(ledgersReq.Assets) > 0 {
		assetStrings := make([]string, len(ledgersReq.Assets))
		for index := range ledgersReq.Assets {
			assetStrings[index] = string(ledgersReq.Assets[index])
		}

		body.Add("asset", strings.Join(assetStrings, ","))
	}

	if string(ledgersReq.Type) != "" {
		body.Add("type", string(ledgersReq.Type))
	}

	if ledgersReq.Start != nil {
		body.Add("start", strconv.FormatInt(ledgersReq.Start.Unix(), 10))
	}

	if ledgersReq.End != nil {
		body.Add("end", strconv.FormatInt(ledgersReq.End.Unix(), 10))
	}

	if ledgersReq.Ofs != 0 {
		body.Add("ofs", strconv.Itoa(ledgersReq.Ofs))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, LedgersResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// QueryLedgers returns an associative array of ledgers info.
// https://www.kraken.com/en-gb/help/api#query-ledgers
func (u *UserData) QueryLedgers(ctx context.Context, ids ...int64) (res LedgersResponse, err error) {
	body := url.Values{}

	if len(ids) > 0 {
		idStrings := make([]string, len(ids))
		for index := range ids {
			idStrings[index] = string(strconv.FormatInt(ids[index], 10))
		}

		body.Add("asset", strings.Join(idStrings, ","))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, QueryLedgersResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// TradeVolume returns an associative array of trade volume.
// https://www.kraken.com/en-gb/help/api#get-trade-volume
func (u *UserData) TradeVolume(ctx context.Context, feeInfo bool, pairs ...pairs.AssetPair) (res *TradeVolumeResponse, err error) {
	body := url.Values{
		"fee-info": []string{"false"},
	}

	if feeInfo {
		body.Set("fee-info", "true")
	}

	if len(pairs) > 0 {
		pairStrings := make([]string, len(pairs))
		for index := range pairs {
			pairStrings[index] = string(pairs[index].String())
		}

		body.Add("pair", strings.Join(pairStrings, ","))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, TradeVolumeResource, body)
	if err != nil {
		return
	}

	krakenResp, err := u.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}
