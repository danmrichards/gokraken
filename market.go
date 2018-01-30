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
func (p *Market) Assets(ctx context.Context, assetReq *AssetsRequest) (resp *AssetsResponse, err error) {
	if assetReq == nil {
		assetReq = &AssetsRequest{}
	}

	if assetReq.Info == "" {
		assetReq.Info = "info"
	}

	if assetReq.AClass == "" {
		assetReq.AClass = "currency"
	}

	body := url.Values{
		"info":   []string{assetReq.Info},
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
