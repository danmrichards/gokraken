package gokraken

import (
	"context"
	"net/http"
)

// Market is responsible for communicating with all the public data markey
// endpoints on the Kraken API.
type Market struct {
	Client *Kraken
}

// Time returns the current server time of Kraken.
// https://www.kraken.com/help/api#get-server-time
func (p *Market) Time(ctx context.Context) (resp TimeResponse, err error) {
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
