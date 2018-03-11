package gokraken

import (
	"time"

	"github.com/danmrichards/gokraken/pairs"
)

const (
	// AddOrderResource is the API resource for adding orders.
	AddOrderResource = "AddOrder"

	// ClosedOrdersResource is the API resource for closed orders.
	ClosedOrdersResource = "ClosedOrders"

	// QueryOrdersResource is the API resource for querying orders.
	QueryOrdersResource = "QueryOrders"

	// OrderCloseTimeOpen is the time an order opens.
	OrderCloseTimeOpen OrderCloseTime = "open"

	// OrderCloseTimeOpen is the time an order closes.
	OrderCloseTimeClose OrderCloseTime = "close"

	// OrderCloseTimeOpen is the time an order opens and closes.
	OrderCloseTimeBoth OrderCloseTime = "both"

	// OpenOrdersResource is the API resource for open orders.
	OpenOrdersResource = "OpenOrders"

	// OrderFlagViqc is an order flag for volume in quote currency
	// (not available for leveraged orders).
	OrderFlagViqc OrderFlag = "viqc"

	// OrderFlagFcib is an order flag for prefer fee in base currency.
	OrderFlagFcib OrderFlag = "fcib"

	// OrderFlagFciq is an order flag for prefer fee in quote currency.
	OrderFlagFciq OrderFlag = "fciq"

	// OrderFlagNompp is an order flag for no market price protection.
	OrderFlagNompp OrderFlag = "nompp"

	// OrderFlagPost is an order flag for post only order
	// (available when ordertype // limit).
	OrderFlagPost OrderFlag = "post"

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

// ClosedOrdersRequest represents a request to get closed orders from Kraken.
type ClosedOrdersRequest struct {
	Trades    bool           // Whether or not to include trades in output.
	UserRef   int64          // Restrict results to given user reference id.
	Start     *time.Time     // Starting unix timestamp or order tx id of results.
	End       *time.Time     // Ending unix timestamp or order tx id of results (optional.  inclusive)
	Ofs       int            // Result offset.
	CloseTime OrderCloseTime // Which time to use.
}

// ClosedOrdersResponse represents the response from the ClosedOrders endpoint
// of the Kraken API.
type ClosedOrdersResponse struct {
	Closed map[string]Order `json:"closed"`
	Count  int              `json:"count"`
}

// OrderCloseTime is the time to close an order.
type OrderCloseTime string

// QueryOrdersResponse represents the response from the QueryOrders endpoint
// of the Kraken API.
type QueryOrdersResponse map[string]Order

// OrderFlag is a flag to use when creating a Kraken order.
type OrderFlag string

// OrderType is a type of Kraken order.
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

// OpenOrdersResponse represents the response from the OpenOrders endpoint
// of the Kraken API.
type OpenOrdersResponse struct {
	Open  map[string]Order `json:"open"`
	Count int              `json:"count"`
}

// UserOrder represents a user request to add an order.
type UserOrder struct {
	Pair              pairs.AssetPair // Asset pair.
	Type              TradeBuySell    // Type of order (buy/sell).
	OrderType         OrderType       // Order type.
	Price             float64         // Price.
	Price2            float64         // Secondary price.
	Volume            float64         // Order volume in lots.
	Leverage          string          // Amount of leverage desired.
	OFlags            []OrderFlag     // List of order flags.
	StartTm           string          // Scheduled start time: 0 (now - default), +<n> (schedule start time <n> seconds from now) , <n> (unix timestamp of start time).
	ExpireTm          string          // expiration time 0 (no expiration - default), +<n> (expire <n> seconds from now) , <n> (unix timestamp of expiration time).
	UserRef           int             // User reference id.
	Validate          bool            // Validate inputs only.
	CloseOrderType    OrderType       // Type of closing order to add to system when order gets filled.
	ClosedOrderPrice  float64         // Price of closing order to add to system when order gets filled.
	ClosedOrderPrice2 float64         // Secondary price of closing order to add to system when order gets filled.
}

// AddOrderResponse represents the response from the AddOrder endpoint
// of the Kraken API.
type AddOrderResponse struct {
	Description OrderDescription `json:"descr"`
	TxIDs       []string         `json:"txid"`
}
