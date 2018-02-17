package gokraken

import (
	"encoding/json"
	"strconv"
	"time"
)

// DepthResource is the API resource for the Kraken API order book.
const DepthResource = "Depth"

// DepthRequest represents a request to list the order book from Kraken.
type DepthRequest struct {
	Pair  string
	Count int
}

// TradesResponse represents the response from the Kraken order book endpoint.
type DepthResponse map[string]Depth

// Depth is an order book response for a given asset pair.
type Depth struct {
	Asks []DepthItem `json:"asks"`
	Bids []DepthItem `json:"bids"`
}

// DepthItem is either the asks or bids for an assert pair order book entry.
type DepthItem struct {
	Price     float64
	Volume    float64
	Timestamp time.Time
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the value pointed to by d.
//
// Due to the nature of the Kraken API response we create an auxiliary struct,
// and then a slice of pointers to the auxiliary struct fields. We unmarshal
// into this slice and finally apply any type conversions into the destination
// fields in the value pointed by d.
func (d *DepthItem) UnmarshalJSON(data []byte) error {
	aux := struct {
		price     string
		volume    string
		timestamp int64
	}{}

	tmpAux := []interface{}{
		&aux.price,
		&aux.volume,
		&aux.timestamp,
	}

	err := json.Unmarshal(data, &tmpAux)
	if err != nil {
		return err
	}

	d.Price, err = strconv.ParseFloat(aux.price, 64)
	if err != nil {
		return err
	}

	d.Volume, err = strconv.ParseFloat(aux.volume, 64)
	if err != nil {
		return err
	}

	d.Timestamp = time.Unix(aux.timestamp, 0)

	return nil
}
