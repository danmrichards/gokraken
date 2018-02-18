package gokraken

import "github.com/danmrichards/gokraken/asset"

// BalanceResource is the API resource for balance.
const BalanceResource = "Balance"

// BalanceResponse represents the response from the Balance endpoint of the
// Kraken API.
type BalanceResponse map[asset.Currency]float64
