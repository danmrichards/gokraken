package gokraken

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/danmrichards/gokraken/asset"
)

// Funding is responsible for communicating with all the private user funding
// endpoints on the Kraken API.
type Funding struct {
	Client *Kraken
}

// DepositMethods gets a list of deposit methods via the Kraken API.
// https://www.kraken.com/en-gb/help/api#deposit-methods
func (f *Funding) DepositMethods(ctx context.Context, aclass AssetsClass, asset asset.Currency) (res DepositMethodsResponse, err error) {
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
func (f *Funding) DepositAddresses(ctx context.Context, aclass AssetsClass, asset asset.Currency, method string, new bool) (res DepositAddressesResponse, err error) {
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

	err = krakenResp.ExtractResult(&res)
	return
}

// DepositStatus gets the status of recent deposits via the Kraken api.
// https://www.kraken.com/en-gb/help/api#deposit-methods
func (f *Funding) DepositStatus(ctx context.Context, aclass AssetsClass, asset asset.Currency, method string) (res *DepositStatusResponse, err error) {
	body := url.Values{
		"aclass": {string(aclass)},
		"asset":  {asset.String()},
		"method": {method},
	}

	req, err := f.Client.DialWithAuth(ctx, http.MethodPost, DepositStatusResource, body)
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

// WithdrawInfo gets withdrawal information via the Kraken api.
// https://www.kraken.com/en-gb/help/api#get-withdrawal-info
func (f *Funding) WithdrawInfo(ctx context.Context, aclass AssetsClass, asset asset.Currency, key string, amount float64) (res *WithdrawInfoResponse, err error) {
	body := url.Values{
		"aclass": {string(aclass)},
		"asset":  {asset.String()},
		"key":    {key},
		"amount": {strconv.FormatFloat(amount, 'f', 4, 64)},
	}

	req, err := f.Client.DialWithAuth(ctx, http.MethodPost, WithdrawInfoResource, body)
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

// Withdraw withdraws funds via the Kraken api.
// https://www.kraken.com/en-gb/help/api#withdraw-funds
func (f *Funding) Withdraw(ctx context.Context, aclass AssetsClass, asset asset.Currency, key string, amount float64) (res *WithdrawResponse, err error) {
	body := url.Values{
		"aclass": {string(aclass)},
		"asset":  {asset.String()},
		"key":    {key},
		"amount": {strconv.FormatFloat(amount, 'f', 4, 64)},
	}

	req, err := f.Client.DialWithAuth(ctx, http.MethodPost, WithdrawResource, body)
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

// WithdrawStatus gets status of recent withdrawals via the Kraken api.
// https://www.kraken.com/en-gb/help/api#withdraw-status
func (f *Funding) WithdrawStatus(ctx context.Context, aclass AssetsClass, asset asset.Currency, method string) (res WithdrawStatusResponse, err error) {
	body := url.Values{
		"aclass": {string(aclass)},
		"asset":  {asset.String()},
		"method": {method},
	}

	req, err := f.Client.DialWithAuth(ctx, http.MethodPost, WithdrawStatusResource, body)
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

// WithdrawCancel requests withdrawal cancellation via the Kraken api.
// https://www.kraken.com/en-gb/help/api#withdraw-cancel
func (f *Funding) WithdrawCancel(ctx context.Context, aclass AssetsClass, asset asset.Currency, refID string) (res bool, err error) {
	body := url.Values{
		"aclass": {string(aclass)},
		"asset":  {asset.String()},
		"refid":  {refID},
	}

	req, err := f.Client.DialWithAuth(ctx, http.MethodPost, WithdrawCancelResource, body)
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
