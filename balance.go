package gokraken

import "github.com/danmrichards/gokraken/asset"

const (
	// BalanceResource is the API resource for balance.
	BalanceResource = "Balance"

	// TradeBalanceResource is the API resource for trade balance.
	TradeBalanceResource = "TradeBalance"
)

// BalanceResponse represents the response from the Balance endpoint of the
// Kraken API.
type BalanceResponse map[asset.Currency]float64

// TradeBalanceResponse represents the response from the TradeBalance endpoint of the
// Kraken API.
type TradeBalanceResponse struct {
	EquivalentBalance float64 `json:"eb"`
	TradeBalance      float64 `json:"tb"`
	MarginAmount      float64 `json:"m"`
	UnrealizedNet     float64 `json:"n"`
	Cost              float64 `json:"c"`
	Valuation         float64 `json:"v"`
	Equity            float64 `json:"e"`
	FreeMargin        float64 `json:"mf"`
	MarginLevel       float64 `json:"ml"`
}
