package gokraken

import (
	"context"
	"net/http"
)

// UserData is responsible for communicating with all the private user data
// endpoints on the Kraken API.
type UserData struct {
	Client *Kraken
}

// Balance returns an array of asset names and balance amount.
// https://www.kraken.com/en-gb/help/api#get-account-balance
func (u *UserData) Balance(ctx context.Context) (res *BalanceResponse, err error) {
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
