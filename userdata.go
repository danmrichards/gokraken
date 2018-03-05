package gokraken

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/danmrichards/gokraken/asset"
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

// TradeBalance returns an order info in open array with txid as the key.
// https://www.kraken.com/help/api#get-open-orders
func (u *UserData) OpenOrders(ctx context.Context, trades bool, userRef int64) (res *OpenOrdersResponse, err error) {
	body := url.Values{
		"trades": []string{"false"},
	}

	if trades {
		body["trades"] = []string{"true"}
	}

	if userRef != 0 {
		body["userref"] = []string{strconv.FormatInt(userRef, 10)}
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
