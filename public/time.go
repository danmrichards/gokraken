package public

const TimeResource = "Time" // The API resource for the Kraken API server time.

// TimeResponse represents the response from the Time endpoint of the
// Kraken API.
type TimeResponse struct {
	UnixTime int64  `json:"unixtime"`
	Rfc1123  string `json:"rfc1123"`
}
