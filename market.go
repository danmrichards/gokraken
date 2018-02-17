package gokraken

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Market is responsible for communicating with all the public data market
// endpoints on the Kraken API.
type Market struct {
	Client *Kraken
}

// Time returns the current server time of Kraken.
// https://www.kraken.com/help/api#get-server-time
func (m *Market) Time(ctx context.Context) (res *TimeResponse, err error) {
	req, err := m.Client.Dial(ctx, http.MethodGet, TimeResource, nil)
	if err != nil {
		return
	}

	krakenResp, err := m.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// Assets returns asset information from Kraken.
// https://www.kraken.com/en-gb/help/api#get-asset-info
func (m *Market) Assets(ctx context.Context, assetReq *AssetsRequest) (res *AssetsResponse, err error) {
	if assetReq == nil {
		assetReq = &AssetsRequest{}
	}

	if assetReq.Info == "" {
		assetReq.Info = AssetInfo
	}

	if assetReq.AClass == "" {
		assetReq.AClass = "currency"
	}

	body := url.Values{
		"info":   []string{string(assetReq.Info)},
		"aclass": []string{assetReq.AClass},
	}

	if len(assetReq.Asset) > 0 {
		assetStrings := make([]string, len(assetReq.Asset))
		for i, asset := range assetReq.Asset {
			assetStrings[i] = string(asset)
		}

		body.Add("asset", strings.Join(assetStrings, ","))
	}

	req, err := m.Client.Dial(ctx, http.MethodPost, AssetsResource, body)
	if err != nil {
		return
	}

	krakenResp, err := m.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// AssetPairs returns tradable asset pairs from Kraken.
// https://www.kraken.com/en-gb/help/api#get-tradable-pairs
func (m *Market) AssetPairs(ctx context.Context, assetPairReq *AssetPairsRequest) (res *AssetPairsResponse, err error) {
	if assetPairReq == nil {
		assetPairReq = &AssetPairsRequest{}
	}

	if assetPairReq.Info == "" {
		assetPairReq.Info = AssetPairsInfo
	}

	body := url.Values{
		"info": []string{string(assetPairReq.Info)},
	}

	if len(assetPairReq.Pairs) > 0 {
		pairStrings := make([]string, len(assetPairReq.Pairs))
		for i, asset := range assetPairReq.Pairs {
			pairStrings[i] = string(asset)
		}

		body.Add("pair", strings.Join(pairStrings, ","))
	}

	req, err := m.Client.Dial(ctx, http.MethodPost, AssetPairsResource, body)
	if err != nil {
		return
	}

	krakenResp, err := m.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// Ticker returns ticker information from Kraken.
// https://www.kraken.com/en-gb/help/api#get-ticker-info
func (m *Market) Ticker(ctx context.Context, pairs ...string) (res *TickerResponse, err error) {
	body := url.Values{
		"pair": []string{strings.Join(pairs, ",")},
	}

	req, err := m.Client.Dial(ctx, http.MethodPost, TickerResource, body)
	if err != nil {
		return
	}

	krakenResp, err := m.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// Ohlc returns ohlc information from Kraken.
// https://www.kraken.com/en-gb/help/api#get-ohlc-data
func (m *Market) Ohlc(ctx context.Context, ohlcReq *OhlcRequest) (res *OhlcResponse, err error) {
	if ohlcReq.Interval == 0 {
		ohlcReq.Interval = 1
	}

	body := url.Values{
		"pair":     []string{ohlcReq.Pair},
		"interval": []string{strconv.Itoa(ohlcReq.Interval)},
	}

	if ohlcReq.Since != 0 {
		body["since"] = []string{strconv.FormatInt(ohlcReq.Since, 10)}
	}

	req, err := m.Client.Dial(ctx, http.MethodPost, OhlcResource, body)
	if err != nil {
		return
	}

	krakenResp, err := m.Client.Call(req)
	if err != nil {
		return
	}

	var tmp map[string]interface{}
	err = krakenResp.ExtractResult(&tmp)
	if err != nil {
		return
	}

	lastFloat, ok := tmp["last"].(float64)
	if !ok {
		err = errors.New("could not extract last from ohlc response")
		return
	}

	res = &OhlcResponse{
		Last: int64(lastFloat),
		Data: make([]OhlcData, 0),
	}

	ohlcData, ok := tmp[ohlcReq.Pair].([]interface{})
	if !ok {
		err = fmt.Errorf("could not extract ohlc data where pair=%s", ohlcReq.Pair)
		return
	}

	for key, ohlcDatum := range ohlcData {
		ohlcDatum, ok := ohlcDatum.([]interface{})
		if !ok {
			err = fmt.Errorf("could not extract at ohlcDatum=%d", key)
			return
		}

		timestampFloat, ok := ohlcDatum[0].(float64)
		if !ok {
			err = fmt.Errorf("could not extract timestamp at ohlcDatum=%d", key)
			return
		}
		timestamp := time.Unix(int64(timestampFloat), 0)

		var open float64
		open, err = strconv.ParseFloat(ohlcDatum[1].(string), 64)
		if err != nil {
			err = fmt.Errorf("could not extract open at ohlcDatum=%d", key)
			return
		}

		var high float64
		high, err = strconv.ParseFloat(ohlcDatum[2].(string), 64)
		if err != nil {
			err = fmt.Errorf("could not extract high at ohlcDatum=%d", key)
			return
		}

		var low float64
		low, err = strconv.ParseFloat(ohlcDatum[3].(string), 64)
		if err != nil {
			err = fmt.Errorf("could not extract low at ohlcDatum=%d", key)
			return
		}

		var close float64
		close, err = strconv.ParseFloat(ohlcDatum[4].(string), 64)
		if err != nil {
			err = fmt.Errorf("could not extract close at ohlcDatum=%d", key)
			return
		}

		var vwap float64
		vwap, err = strconv.ParseFloat(ohlcDatum[5].(string), 64)
		if err != nil {
			err = fmt.Errorf("could not extract vwap at ohlcDatum=%d", key)
			return
		}

		var volume float64
		volume, err = strconv.ParseFloat(ohlcDatum[6].(string), 64)
		if err != nil {
			err = fmt.Errorf("could not extract volume at ohlcDatum=%d", key)
			return
		}

		var countFloat float64
		countFloat, ok = ohlcDatum[7].(float64)
		if !ok {
			err = fmt.Errorf("could not extract count at ohlcDatum=%d", key)
			return
		}

		res.Data = append(res.Data, OhlcData{
			Timestamp: timestamp,
			Open:      open,
			High:      high,
			Low:       low,
			Close:     close,
			Vwap:      vwap,
			Volume:    volume,
			Count:     int(countFloat),
		})
	}

	return
}

// Depth returns the order book from Kraken.
// https://www.kraken.com/en-gb/help/api#get-order-book
func (m *Market) Depth(ctx context.Context, depthReq *DepthRequest) (res *DepthResponse, err error) {
	body := url.Values{
		"pair":  []string{depthReq.Pair},
		"count": []string{strconv.Itoa(depthReq.Count)},
	}

	req, err := m.Client.Dial(ctx, http.MethodPost, DepthResource, body)
	if err != nil {
		return
	}

	krakenResp, err := m.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}

// Trades returns the recent trades from Kraken.
// https://www.kraken.com/en-gb/help/api#get-recent-trades
func (m *Market) Trades(ctx context.Context, tradeReq *TradesRequest) (res *TradesResponse, err error) {
	body := url.Values{
		"pair": []string{tradeReq.Pair},
	}

	if tradeReq.Since != 0 {
		body["since"] = []string{strconv.FormatInt(tradeReq.Since, 10)}
	}

	req, err := m.Client.Dial(ctx, http.MethodPost, TradesResource, body)
	if err != nil {
		return
	}

	krakenResp, err := m.Client.Call(req)
	if err != nil {
		return
	}

	var tmp map[string]interface{}
	err = krakenResp.ExtractResult(&tmp)
	if err != nil {
		return
	}

	last, err := strconv.ParseInt(tmp["last"].(string), 10, 64)
	if err != nil {
		return
	}

	res = &TradesResponse{
		Last:   last,
		Trades: make([]Trade, 0),
	}

	trades, ok := tmp[tradeReq.Pair].([]interface{})
	if !ok {
		err = fmt.Errorf("could not extract trades where pair=%s", tradeReq.Pair)
		return
	}

	for key, trade := range trades {
		trade, ok := trade.([]interface{})
		if !ok {
			err = fmt.Errorf("could not extract at trade=%d", key)
			return
		}

		var price float64
		price, err = strconv.ParseFloat(trade[0].(string), 64)
		if err != nil {
			err = fmt.Errorf("could not extract price at trade=%d", key)
			return
		}

		var volume float64
		volume, err = strconv.ParseFloat(trade[1].(string), 64)
		if err != nil {
			err = fmt.Errorf("could not extract volume at trade=%d", key)
			return
		}

		timestampFloat, ok := trade[2].(float64)
		if !ok {
			err = fmt.Errorf("could not extract timestamp at trade=%d", key)
			return
		}
		timestamp := time.Unix(int64(timestampFloat), 0)

		var buySell TradeBuySell
		switch trade[3].(string) {
		case "b":
			buySell = TradeBuy
		case "s":
			buySell = TradeSell
		}

		var marketLimit TradeMarketLimit
		switch trade[4].(string) {
		case "m":
			marketLimit = TradeMarket
		case "l":
			marketLimit = TradeLimit
		}

		misc, ok := trade[5].(string)
		if !ok {
			err = fmt.Errorf("could not extract miscellaneous at trade=%d", key)
			return
		}

		res.Trades = append(res.Trades, Trade{
			Price:         price,
			Volume:        volume,
			Timestamp:     timestamp,
			BuySell:       buySell,
			MarketLimit:   marketLimit,
			Miscellaneous: misc,
		})
	}

	return
}
