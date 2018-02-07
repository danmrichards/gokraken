package gokraken

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

// Market is responsible for communicating with all the public data market
// endpoints on the Kraken API.
type Market struct {
	Client *Kraken
}

// Time returns the current server time of Kraken.
// https://www.kraken.com/help/api#get-server-time
func (p *Market) Time(ctx context.Context) (resp *TimeResponse, err error) {
	req, err := p.Client.Dial(ctx, http.MethodGet, TimeResource, nil)
	if err != nil {
		return
	}

	krakenResp, err := p.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&resp)
	return
}

// Assets returns asset information from Kraken.
// https://www.kraken.com/en-gb/help/api#get-asset-info
func (p *Market) Assets(ctx context.Context, assetReq *AssetsRequest) (resp AssetsResponse, err error) {
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

	req, err := p.Client.Dial(ctx, http.MethodPost, AssetsResource, body)
	if err != nil {
		return
	}

	krakenResp, err := p.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&resp)
	return
}

// AssetPairs returns tradable asset pairs from Kraken.
// https://www.kraken.com/en-gb/help/api#get-tradable-pairs
func (p *Market) AssetPairs(ctx context.Context, assetPairReq *AssetPairsRequest) (resp AssetPairsResponse, err error) {
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

	req, err := p.Client.Dial(ctx, http.MethodPost, AssetPairsResource, body)
	if err != nil {
		return
	}

	krakenResp, err := p.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&resp)
	return
}

// Ticker returns ticker information from Kraken.
// https://www.kraken.com/en-gb/help/api#get-ticker-info
func (p *Market) Ticker(ctx context.Context, pairs ...string) (resp TickerResponse, err error) {
	body := url.Values{
		"pair": []string{strings.Join(pairs, ",")},
	}

	req, err := p.Client.Dial(ctx, http.MethodPost, TickerResource, body)
	if err != nil {
		return
	}

	krakenResp, err := p.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&resp)
	return
}
