package gokraken

// TimeResource is the API resource for the Kraken API server time.
const TimeResource = "Time"

// TimeResponse represents the response from the Time endpoint of the
// Kraken API.
type TimeResponse struct {
	UnixTime int64  `json:"unixtime"`
	Rfc1123  string `json:"rfc1123"`
}
