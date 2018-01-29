package gokraken

const BalanceResource = "Balance" // The API resource for balance.

// BalanceResponse represents the response from the Balance endpoint of the
// Kraken API.
type BalanceResponse map[Currency]float64
