package gokraken

import "github.com/danmrichards/gokraken/pairs"

// TickerResource is the API resource for the Kraken API ticker.
const TickerResource = "Ticker"

// TickerResponse represents the response from the Ticker endpoint of the
// Kraken API.
type TickerResponse map[pairs.AssetPair]TickerInfo

// TickerInfo represents the TickerInfo for an asset pair.
type TickerInfo struct {
	A []string `json:"a"` // ask array(<price>, <whole lot volume>, <lot volume>).
	B []string `json:"b"` // bid array(<price>, <whole lot volume>, <lot volume>).
	C []string `json:"c"` // last trade closed array(<price>, <lot volume>).
	V []string `json:"v"` // volume array(<today>, <last 24 hours>).
	P []string `json:"p"` // volume weighted average price array(<today>, <last 24 hours>).
	T []int    `json:"t"` // number of trades array(<today>, <last 24 hours>).
	L []string `json:"l"` // low array(<today>, <last 24 hours>).
	H []string `json:"h"` // high array(<today>, <last 24 hours>).
	O string   `json:"o"` // today's opening price.
}
