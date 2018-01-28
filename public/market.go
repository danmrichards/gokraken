package public

import (
	"context"
	"net/http"

	"github.com/danmrichards/gokraken"
)

// Market is responsible for communicating with all the public
// public data endpoints on the Kraken API.
type Market struct {
	Client gokraken.KrakenClient
}

// Time returns the current server time of Kraken.
// https://www.kraken.com/help/api#get-server-time
func (p *Market) Time(ctx context.Context) (resp *TimeResponse, err error) {
	req, err := p.Client.PrepareRequest(ctx, http.MethodGet, TimeResource, nil)
	if err != nil {
		return
	}

	krakenResp, err := p.Client.Do(req)
	if err != nil {
		return
	}

	err = krakenResp.ExtractResult(&resp)
	return
}
