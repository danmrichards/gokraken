package gokraken

const (
	// WithdrawInfoResource is the API resource for withdrawal information.
	WithdrawInfoResource = "WithdrawInfo"

	// WithdrawResource is the API resource for withdrawing funds.
	WithdrawResource = "Withdraw"

	// WithdrawStatusResource is the API resource for getting the status of a
	// withdrawal.
	WithdrawStatusResource = "WithdrawStatus"

	// WithdrawCancelResource is the API resource for canceling a withdrawal.
	WithdrawCancelResource = "WithdrawCancel"
)

// WithdrawInfoResponse represents withdrawal information.
type WithdrawInfoResponse struct {
	Method string  `json:"method"`
	Limit  float64 `json:"limit"`
	Fee    float64 `json:"fee"`
}

// WithdrawResponse represents the response to a withdraw funds request.
type WithdrawResponse struct {
	RefID string `json:"refid"`
}

// WithdrawStatusResponse represents a list of withdrawal statuses.
type WithdrawStatusResponse []WithdrawStatus

// WithdrawStatus represents the status of a withdrawal.
type WithdrawStatus struct {
	Method     string             `json:"method"`
	Aclass     string             `json:"aclass"`
	Asset      string             `json:"asset"`
	RefID      string             `json:"refid"`
	TxID       string             `json:"txid"`
	Info       string             `json:"info"`
	Amount     float64            `json:"float64"`
	Fee        float64            `json:"fee"`
	Time       int64              `json:"time"`
	Status     string             `json:"status"`
	StatusProp WithdrawStatusProp `json:"status-prop"`
}

// WithdrawStatusProp contains additional status properties for a withdrawal.
type WithdrawStatusProp struct {
	CancelPending bool `json:"cancel-pending"`
	Canceled      bool `json:"canceled"`
	CancelDenied  bool `json:"cancel-denied"`
	Return        bool `json:"return"`
	OnHold        bool `json:"onhold"`
}
