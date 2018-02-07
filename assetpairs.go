package gokraken

const (
	AssetPairsResource                     = "AssetPairs" // The API resource for the Kraken API asset info.
	AssetPairsInfo     AssetPairsInfoLevel = "info"       // Info level for asset pairs.
	AssetPairsLeverage AssetPairsInfoLevel = "leverage"   // Leverage level for asset pairs.
	AssetPairsFees     AssetPairsInfoLevel = "fees"       // Fees level for asset pairs.
	AssetPairsMargin   AssetPairsInfoLevel = "margin"     // Margin level for asset pairs.
)

// AssetPairsInfoLevel represents an info level for an asset pairs request.
type AssetPairsInfoLevel string

// AssetPairsRequest represents a request to list asset pairs from Kraken.
type AssetPairsRequest struct {
	Info  AssetPairsInfoLevel // Info to retrieve (default: info).
	Pairs []string            // List of asset pairs to get info on.
}

// AssetPairsResponse represents an array of asset pairs and their info.
type AssetPairsResponse map[string]*AssetPair

// AssetPair represents a tradeable asset pair from Kraken.
type AssetPair struct {
	Altname           string      `json:"altname"`
	AclassBase        string      `json:"aclass_base"`
	Base              string      `json:"base"`
	AclassQuote       string      `json:"aclass_quote"`
	Quote             string      `json:"quote"`
	Lot               string      `json:"lot"`
	PairDecimals      int         `json:"pair_decimals"`
	LotDecimals       int         `json:"lot_decimals"`
	LotMultiplier     int         `json:"lot_multiplier"`
	LeverageBuy       []float64   `json:"leverage_buy"`
	LeverageSell      []float64   `json:"leverage_sell"`
	Fees              [][]float64 `json:"fees"`
	FeesMaker         [][]float64 `json:"fees_maker"`
	FeeVolumeCurrency string      `json:"fee_volume_currency"`
	MarginCall        int         `json:"margin_call"`
	MarginStop        int         `json:"margin_stop"`
}
