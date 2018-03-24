package gokraken

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/danmrichards/gokraken/asset"
)

// Funding is responsible for communicating with all the private user funding
// endpoints on the Kraken API.
type Funding struct {
	Client *Kraken
}

// DepositMethods gets a list of deposit methods via the Kraken API.
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

// DepositAddresses gets a list of deposit addresses via the Kraken API.
// https://www.kraken.com/en-gb/help/api#deposit-methods
func (f *Funding) DepositAddresses(ctx context.Context, aclass AssetsClass, asset asset.Currency, method string, new bool) (res *DepositAddressesResponse, err error) {
	body := url.Values{
		"aclass": {string(aclass)},
		"asset":  {asset.String()},
		"method": {method},
		"new":    {"false"},
	}

	if new {
		body.Set("new", "true")
	}

	req, err := f.Client.DialWithAuth(ctx, http.MethodPost, DepositAddressesResource, body)
	if err != nil {
		return
	}

	krakenResp, err := f.Client.Call(req)
	if err != nil {
		return
	}

	fmt.Println(krakenResp)

	err = krakenResp.ExtractResult(&res)
	return
}
