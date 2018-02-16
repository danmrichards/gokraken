package gokraken

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Market is responsible for communicating with all the public data market
// endpoints on the Kraken API.
type Market struct {
	Client *Kraken
}

// Time returns the current server time of Kraken.
// https://www.kraken.com/help/api#get-server-time
func (m *Market) Time(ctx context.Context) (resp *TimeResponse, err error) {
	req, err := m.Client.Dial(ctx, http.MethodGet, TimeResource, nil)
	if err != nil {
		return
	}

	krakenResp, err := m.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&resp)
	return
}

// Assets returns asset information from Kraken.
// https://www.kraken.com/en-gb/help/api#get-asset-info
func (m *Market) Assets(ctx context.Context, assetReq *AssetsRequest) (resp AssetsResponse, err error) {
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

	err = krakenResp.ExtractResult(&resp)
	return
}

// AssetPairs returns tradable asset pairs from Kraken.
// https://www.kraken.com/en-gb/help/api#get-tradable-pairs
func (m *Market) AssetPairs(ctx context.Context, assetPairReq *AssetPairsRequest) (resp AssetPairsResponse, err error) {
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

	err = krakenResp.ExtractResult(&resp)
	return
}

// Ticker returns ticker information from Kraken.
// https://www.kraken.com/en-gb/help/api#get-ticker-info
func (m *Market) Ticker(ctx context.Context, pairs ...string) (resp TickerResponse, err error) {
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

	err = krakenResp.ExtractResult(&resp)
	return
}

// Ohlc returns holc information from Kraken.
// https://www.kraken.com/en-gb/help/api#get-ohlc-data
func (m *Market) Ohlc(ctx context.Context, ohlcReq *OhlcRequest) (resp OhlcResponse, err error) {
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

	err = krakenResp.ExtractResult(&resp)
	return
}
