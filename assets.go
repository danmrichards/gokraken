package gokraken

const (
	// AssetsResource is the API resource for the Kraken API asset info.
	AssetsResource = "Assets"

	// AssetInfo is the info level for assets.
	AssetInfo AssetsInfoLevel = "info"
)

// AssetsInfoLevel represents an info level for an assets request.
type AssetsInfoLevel string

// AssetsRequest represents a request to list asset information from Kraken.
type AssetsRequest struct {
	Info   AssetsInfoLevel // Info to retrieve (default: info).
	AClass string          // Asset class (default: currency).
	Asset  []Currency      // List of assets.
}

// AssetsResponse represents an array of asset names and their info.
type AssetsResponse map[Currency]*Asset

// Asset represents a Kraken asset.
type Asset struct {
	AltName         string `json:"altname"`          // Alternate name.
	AClass          string `json:"aclass"`           // Asset class.
	Decimals        int    `json:"decimals"`         // Scaling decimal places for record keeping.
	DisplayDecimals int    `json:"display_decimals"` // Scaling decimal places for output display.
}
