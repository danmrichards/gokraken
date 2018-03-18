package gokraken

const (
	// DepositMethodsResource is the API resource for deposit methods.
	DepositMethodsResource = "DepositMethods"
)

// DepositMethodsResponse represents the response from the DepositMethods
// endpoint of the Kraken API.
type DepositMethodsResponse struct {
	Method          string  `json:"method"`
	Limit           float64 `json:"limit"`
	Fee             float64 `json:"fee"`
	AddressSetupFee bool    `json:"address-setup-fee"`
}
