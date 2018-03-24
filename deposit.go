package gokraken

const (
	// DepositMethodsResource is the API resource for deposit methods.
	DepositMethodsResource = "DepositMethods"

	// DepositAddressesResource is the API resource for deposit addresses.
	DepositAddressesResource = "DepositAddresses"

	// DepositStatusResource is the API resource for deposit status.
	DepositStatusResource = "DepositStatus"
)

// DepositMethodsResponse represents the response from the DepositMethods
// endpoint of the Kraken API.
type DepositMethodsResponse map[string]DepositMethod

// DepositMethod represents a Kraken deposit method.
type DepositMethod struct {
	Method          string  `json:"method"`
	Limit           float64 `json:"limit"`
	Fee             float64 `json:"fee"`
	AddressSetupFee bool    `json:"address-setup-fee"`
}

// DepositAddressesResponse represents the response from the DepositAddresses
// endpoint of the Kraken API.
type DepositAddressesResponse map[string]DepositAddress

// DepositAddress represents a Kraken deposit address.
type DepositAddress struct {
	Address    string `json:"address"`
	ExpireTime int64  `json:"expiretm"`
	New        bool   `json:"new"`
}

// DepositStatusResponse represents the response from the DepositStatus
// endpoint of the Kraken API.
type DepositStatusResponse struct {
	Method     string            `json:"method"`
	AClass     string            `json:"aclass"`
	Asset      string            `json:"asset"`
	RefID      string            `json:"refid"`
	TxID       string            `json:"txid"`
	Info       string            `json:"info"`
	Amount     float64           `json:"amount"`
	Fee        float64           `json:"fee"`
	Time       int64             `json:"time"`
	Status     string            `json:"status"`
	StatusProp DepositStatusProp `json:"status-prop"`
}

// DepositStatusProp represents additional status properties on a Kraken deposit
// status.
type DepositStatusProp struct {
	OnHold string `json:"onhold"`
}
