package gokraken

import (
	"time"

	"github.com/danmrichards/gokraken/asset"
)

const (
	// LedgersResource is the API resource for the Kraken API ledgers endpoint.
	LedgersResource = "Ledgers"

	// LedgerTypeAll is used to list all ledgers.
	LedgerTypeAll LedgerType = "all"

	// LedgerTypeDeposit is used to list deposit ledgers.
	LedgerTypeDeposit LedgerType = "deposit"

	// LedgerTypeWithdrawal is used to list withdrawal ledgers.
	LedgerTypeWithdrawal LedgerType = "withdrawal"

	// LedgerTypeTrade is used to list trade ledgers.
	LedgerTypeTrade LedgerType = "trade"

	// LedgerTypeMargin is used to list margin ledgers.
	LedgerTypeMargin LedgerType = "margin"

	// QueryLedgersResource is the API resource for the Kraken API query ledgers endpoint.
	QueryLedgersResource = "QueryLedgers"
)

// LedgerType is a type of kraken ledger.
type LedgerType string

// LedgersRequest represents a request to list ledger information from Kraken.
type LedgersRequest struct {
	Aclass AssetsClass      // Asset class.
	Assets []asset.Currency // List of assets to restrict output to.
	Type   LedgerType       // Type of ledger to retrieve.
	Start  *time.Time       // Starting timestamp.
	End    *time.Time       // Ending timestamp.
	Ofs    int              // Offset.
}

// LedgersResponse represents the response from the OpenPositions endpoint
// of the Kraken API.
type LedgersResponse map[string]Ledger

// Ledger represent a Kraken ledger entry.
type Ledger struct {
	Refid   string  `json:"refid"`
	Time    int64   `json:"time"`
	Type    string  `json:"type"`
	Aclass  string  `json:"aclass"`
	Asset   string  `json:"asset"`
	Amount  float64 `json:"amount"`
	Fee     float64 `json:"fee"`
	Balance float64 `json:"balance"`
}
