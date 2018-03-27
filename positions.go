package gokraken

// OpenPositionsResource is the API resource for the Kraken API open positions.
const OpenPositionsResource = "OpenPositions"

// OpenPositionsResponse represents the response from the OpenPositions endpoint
// of the Kraken API.
type OpenPositionsResponse map[string]Position

// Position represents a Kraken open position.
type Position struct {
	OrderTxid string  `json:"ordertxid"`  // Order responsible for execution of trade.
	Pair      string  `json:"pair"`       // Asset pair.
	Time      int64   `json:"time"`       // Unix timestamp of trade.
	Type      string  `json:"type"`       // Type of order used to open position (buy/sell).
	OrderType string  `json:"ordertype"`  // Order type used to open position.
	Cost      float64 `json:"cost"`       // Opening cost of position (quote currency unless viqc set in oflags).
	Fee       float64 `json:"fee"`        // Opening fee of position (quote currency).
	Vol       float64 `json:"vol"`        // Position volume (base currency unless viqc set in oflags).
	VolClosed float64 `json:"vol_closed"` // Position volume closed (base currency unless viqc set in oflags).
	Margin    float64 `json:"margin"`     // Initial margin (quote currency).
	Value     float64 `json:"value"`      // Current value of remaining position (if docalcs requested.  quote currency).
	Net       float64 `json:"net"`        // Unrealized profit/loss of remaining position (if docalcs requested.  quote currency, quote currency scale).
	Misc      string  `json:"misc"`       // Comma delimited list of miscellaneous info.
	OFlags    string  `json:"oflags"`     // Comma delimited list of order flags.
	Viqc      float64 `json:"viqc"`       // Volume in quote currency.
}
