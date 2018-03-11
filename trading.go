package gokraken

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Trading is responsible for communicating with all the private user trading
// endpoints on the Kraken API.
type Trading struct {
	Client *Kraken
}

// AddOrder adds a standard order via the Kraken API.
// https://www.kraken.com/en-gb/help/api#add-standard-order
func (u *UserData) AddOrder(ctx context.Context, order UserOrder) (res *AddOrderResponse, err error) {
	body := url.Values{
		"pair":      {order.Pair.String()},
		"type":      {string(order.Type)},
		"ordertype": {string(order.OrderType)},
		"volume":    {strconv.FormatFloat(order.Volume, 'f', 4, 64)},
	}

	if order.Price != 0 {
		body.Add("price", strconv.FormatFloat(order.Price, 'f', 4, 64))
	}

	if order.Price2 != 0 {
		body.Add("price2", strconv.FormatFloat(order.Price2, 'f', 4, 64))
	}

	if order.Leverage != "" {
		body.Add("leverage", order.Leverage)
	}

	if len(order.OFlags) > 0 {
		flagStrings := make([]string, len(order.OFlags))
		for index := range order.OFlags {
			flagStrings[index] = string(order.OFlags[index])
		}

		body.Add("oflags", strings.Join(flagStrings, ","))
	}

	if order.StartTm != "" {
		body.Add("starttm", order.StartTm)
	}

	if order.ExpireTm != "" {
		body.Add("expiretm", order.ExpireTm)
	}

	if order.Validate {
		body.Add("validate", "true")
	}

	if string(order.CloseOrderType) != "" {
		body.Add("close[ordertype]", string(order.CloseOrderType))
	}

	if order.ClosedOrderPrice != 0 {
		body.Add("close[price]", strconv.FormatFloat(order.ClosedOrderPrice, 'f', 4, 64))
	}

	if order.ClosedOrderPrice2 != 0 {
		body.Add("close[price2]", strconv.FormatFloat(order.ClosedOrderPrice2, 'f', 4, 64))
	}

	req, err := u.Client.DialWithAuth(ctx, http.MethodPost, AddOrderResource, body)
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
