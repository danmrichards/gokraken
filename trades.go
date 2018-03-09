package gokraken

import (
	"time"

	"github.com/danmrichards/gokraken/pairs"
)

const (
	// TradesResource is the API resource for the Kraken API recent trades.
	TradesResource = "Trades"

	// TradesHistoryResource is the API resource for the Kraken API trade history.
	TradesHistoryResource = "TradesHistory"

	// TradesInfoResource is the API resource for the Kraken API trades info.
	TradesInfoResource = "TradesInfo"

	// TradeBuy indicates that a trade was a 'buy'.
	TradeBuy TradeBuySell = "buy"

	// TradeSell indicates that a trade was a 'sell'.
	TradeSell TradeBuySell = "sell"

	// TradeMarket indicates that a trade was 'market'.
	TradeMarket TradeMarketLimit = "market"

	// TradeLimit indicates that a trade was 'limit'.
	TradeLimit TradeMarketLimit = "limit"

	// TradeTypeAll indicates all trades.
	TradeTypeAll TradeType = "all"

	// TradeTypeAny indicates any position (opened or closed).
	TradeTypeAny TradeType = "any position"

	// TradeTypeClosed indicates positions that have closed.
	TradeTypeClosed TradeType = "closed position"

	// TradeTypeClosing indicates any trade closing all or part of a position.
	TradeTypeClosing TradeType = "closing position"

	// TradeTypeNoPosition indicates non-positional trades.
	TradeTypeNoPosition TradeType = "no position"
)

// TradesRequest represents a request to get recent trades from Kraken.
type TradesRequest struct {
	Pair  pairs.AssetPair
	Since int64
}

// TradesResponse represents the response from the Kraken recent trades endpoint.
type TradesResponse struct {
	Trades []Trade
	Last   int64
}

// Trade is a trade of asset.
type Trade struct {
	Price         float64
	Volume        float64
	Timestamp     time.Time
	BuySell       TradeBuySell
	MarketLimit   TradeMarketLimit
	Miscellaneous string
}

// TradeBuySell indicates if a trade was a 'buy' or a 'sell'.
type TradeBuySell string

// TradeMarketLimit indicates if a trade was 'market' or 'limit'.
type TradeMarketLimit string

// TradeType indicates the type of trade.
type TradeType string

// TradesHistoryRequest represents a request to get a users trade history.
type TradesHistoryRequest struct {
	Type   TradeType  // The type of trade.
	Trades bool       // Whether or not to include trades in output.
	Start  *time.Time // Starting unix timestamp or order tx id of results.
	End    *time.Time // Ending unix timestamp or order tx id of results.
	Ofs    int        // Result offset.
}

// TradesHistoryResponse represents the response from the TradesHistory endpoint
// of the Kraken API.
type TradesHistoryResponse struct {
	Trades map[string]UserTrade `json:"trades"`
	Count  int                  `json:"count"`
}

// QueryTradesResponse represents the response from the TradesInfo endpoint
// of the Kraken API.
type QueryTradesResponse map[string]UserTrade

// UserTrade represents a historical user trade.
type UserTrade struct {
	OrderTxid string       `json:"ordertxid"`
	Pair      string       `json:"pair"`
	Time      int64        `json:"time"`
	Type      TradeBuySell `json:"type"`
	OrderType OrderType    `json:"ordertype"`
	Price     float64      `json:"price"`
	Cost      float64      `json:"cost"`
	Fee       float64      `json:"fee"`
	Vol       float64      `json:"vol"`
	Margin    float64      `json:"margin"`
	Misc      string       `json:"misc"`
}
