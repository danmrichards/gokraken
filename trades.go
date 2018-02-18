package gokraken

import (
	"time"

	"github.com/danmrichards/gokraken/pairs"
)

const (
	// TradesResource is the API resource for the Kraken API recent trades.
	TradesResource = "Trades"

	// TradeBuy indicates that a trade was a 'buy'.
	TradeBuy TradeBuySell = "buy"

	// TradeSell indicates that a trade was a 'sell'.
	TradeSell TradeBuySell = "sell"

	// TradeMarket indicates that a trade was 'market'.
	TradeMarket TradeMarketLimit = "market"

	// TradeLimit indicates that a trade was 'limit'.
	TradeLimit TradeMarketLimit = "limit"
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
