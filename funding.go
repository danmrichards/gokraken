package gokraken

import (
	"context"
	"net/http"
	"net/url"

	"github.com/danmrichards/gokraken/asset"
)

// Funding is responsible for communicating with all the private user funding
// endpoints on the Kraken API.
type Funding struct {
	Client *Kraken
}

// DepositMethods adds a standard order via the Kraken API.
// https://www.kraken.com/en-gb/help/api#deposit-methods
func (f *Funding) DepositMethods(ctx context.Context, aclass AssetsClass, asset asset.Currency) (res *DepositMethodsResponse, err error) {
	body := url.Values{
		"aclass": {string(aclass)},
		"asset":  {asset.String()},
	}

	req, err := f.Client.DialWithAuth(ctx, http.MethodPost, DepositMethodsResource, body)
	if err != nil {
		return
	}

	krakenResp, err := f.Client.Call(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&res)
	return
}
