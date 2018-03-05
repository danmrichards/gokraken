package gokraken

import "github.com/danmrichards/gokraken/pairs"

const (
	// OpenOrdersResource is the API resource for open orders.
	OpenOrdersResource = "OpenOrders"

	// OrderTypeMarket is the market order type.
	OrderTypeMarket OrderType = "market"

	// OrderTypeLimit (price = limit price).
	OrderTypeLimit OrderType = "limit"

	// OrderTypeStopLoss (price = stop loss price).
	OrderTypeStopLoss OrderType = "stop-loss"

	// OrderTypeTakeProfit (price = take profit price).
	OrderTypeTakeProfit OrderType = "take-profit"

	// OrderTypeStopLossProfit (price = stop loss price, price2 = take profit price).
	OrderTypeStopLossProfit OrderType = "stop-loss-profit"

	// OrderTypeStopLossProfitLimit (price = stop loss price, price2 = take profit price).
	OrderTypeStopLossProfitLimit OrderType = "stop-loss-profit-limit"

	// OrderTypeStopLossLimit (price = stop loss trigger price, price2 = triggered limit price).
	OrderTypeStopLossLimit OrderType = "stop-loss-limit"

	// OrderTypeTakeProfitLimit (price = take profit trigger price, price2 = triggered limit price).
	OrderTypeTakeProfitLimit OrderType = "take-profit-limit"

	// OrderTypeTrailingStop (price = trailing stop offset).
	OrderTypeTrailingStop OrderType = "trailing-stop"

	// OrderTypeTrailingStopLimit (price = trailing stop offset, price2 = triggered limit offset).
	OrderTypeTrailingStopLimit OrderType = "trailing-stop-limit"

	// OrderTypeStopLossAndLimit (price = stop loss price, price2 = limit price).
	OrderTypeStopLossAndLimit OrderType = "stop-loss-and-limit"

	// OrderTypeSettlePosition.
	OrderTypeSettlePosition OrderType = "settle-position"
)

type OrderType string

// Order represents a single order
type Order struct {
	TransactionID  string           `json:"-"`
	ReferenceID    string           `json:"refid"`
	UserRef        int64            `json:"userref"`
	Status         string           `json:"status"`
	OpenTime       float64          `json:"opentm"`
	StartTime      float64          `json:"starttm"`
	ExpireTime     float64          `json:"expiretm"`
	Description    OrderDescription `json:"descr"`
	Volume         string           `json:"vol"`
	VolumeExecuted float64          `json:"vol_exec,string"`
	Cost           float64          `json:"cost,string"`
	Fee            float64          `json:"fee,string"`
	Price          float64          `json:"price,string"`
	StopPrice      float64          `json:"stopprice.string"`
	LimitPrice     float64          `json:"limitprice,string"`
	Misc           string           `json:"misc"`
	OrderFlags     string           `json:"oflags"`
	CloseTime      float64          `json:"closetm"`
	Reason         string           `json:"reason"`
}

// OrderDescription describes the detail of an order including the assets pair
// involved, prices and the order type.
type OrderDescription struct {
	AssetPair      pairs.AssetPair `json:"pair"`
	Close          string          `json:"close"`
	Leverage       string          `json:"leverage"`
	Order          string          `json:"order"`
	OrderType      OrderType       `json:"ordertype"`
	PrimaryPrice   string          `json:"price"`
	SecondaryPrice string          `json:"price2"`
	Type           string          `json:"type"`
}

// OpenOrdersResponse represents an open order from Kraken.
type OpenOrdersResponse struct {
	Open  map[string]Order `json:"open"`
	Count int              `json:"count"`
}
