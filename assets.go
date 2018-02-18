package gokraken

import "github.com/danmrichards/gokraken/asset"

const (
	// AssetsResource is the API resource for the Kraken API asset info.
	AssetsResource = "Assets"

	// AssetInfo is the info level for asset.
	AssetInfo AssetsInfoLevel = "info"

	// AssetCurrency is the asset class for asset.
	AssetCurrency AssetsClass = "currency"
)

// AssetsInfoLevel represents an info level for an asset request.
type AssetsInfoLevel string

// AssetsClass represents an asset class for an asset request.
type AssetsClass string

// AssetsResponse represents an array of asset names and their info.
type AssetsResponse map[asset.Currency]Asset

// Asset represents a Kraken asset.
type Asset struct {
	AltName         string      `json:"altname"`          // Alternate name.
	AClass          AssetsClass `json:"aclass"`           // Asset class.
	Decimals        int         `json:"decimals"`         // Scaling decimal places for record keeping.
	DisplayDecimals int         `json:"display_decimals"` // Scaling decimal places for output display.
}
