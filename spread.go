package gokraken

import "time"

// SpreadResource is the API resource for the Kraken API spread data.
const SpreadResource = "Spread"

// SpreadRequest represents a request to get spread data from Kraken.
type SpreadRequest struct {
	Pair  string
	Since int64
}

// SpreadResponse represents the response from the Kraken spread data endpoint.
type SpreadResponse struct {
	Data []SpreadData
	Last int64
}

// SpreadData is the spread of data for trades.
type SpreadData struct {
	Timestamp time.Time
	Bid       float64
	Ask       float64
}
