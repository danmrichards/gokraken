package gokraken

import (
	"time"

	"github.com/danmrichards/gokraken/pairs"
)

// OhlcResource is the API resource for the Kraken API ohlc.
const OhlcResource = "OHLC"

// OhlcRequest represents a request to list OHLC information from Kraken.
type OhlcRequest struct {
	Pair     pairs.AssetPair
	Interval int
	Since    int64
}

// OhlcResponse represents the response from the Kraken ohlc endpoint.
type OhlcResponse struct {
	Data []OhlcData
	Last int64
}

// OhlcData represents a set of OHLC data from Kraken.
type OhlcData struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Vwap      float64
	Volume    float64
	Count     int
}
