package gokraken

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// OhlcResource is the API resource for the Kraken API ohlc.
const OhlcResource = "OHLC"

// OhlcRequest represents a request to list OHLC information from Kraken.
type OhlcRequest struct {
	Pair     string
	Interval int
	Since    int64
}

// OhlcResponse represents an array of asset pairs and ohlc data.
type OhlcResponse map[string]interface{}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the value pointed to by o.
//
// Due to the nature of the Kraken API response we iterate over a map of
// interface{} keyed by string. Each key is then checked and the value type
// asserted accordingly.
//
// The element at key "last" is the id to be used as "Since" when polling for
// new, committed OHLC data.
func (o *OhlcResponse) UnmarshalJSON(data []byte) error {
	var aux map[string]interface{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	for k := range aux {
		switch k {
		case "last":
			timeUnix, ok := aux[k].(float64)
			if !ok {
				return errors.New("could not assert time as int64")
			}

			aux[k] = time.Unix(int64(timeUnix), 0)
		default:
			valBytes, err := json.Marshal(aux[k])
			if err != nil {
				return err
			}

			var ohlcData []OhlcData
			err = json.Unmarshal(valBytes, &ohlcData)
			if err != nil {
				return err
			}

			aux[k] = ohlcData
		}
	}

	*o = aux

	return nil
}

// OhlcData represents a set of OHLC data from Kraken.
type OhlcData struct {
	Time   time.Time `json:"time"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Vwap   float64   `json:"vwap"`
	Volume float64   `json:"volume"`
	Count  int       `json:"count"`
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the value pointed to by o.
//
// Due to the nature of the Kraken API response we create an auxiliary struct,
// and then a slice of pointers to the auxiliary struct fields. We unmarshal
// into this slice and finally apply any type conversions into the destination
// fields in the value pointed by o.
func (o *OhlcData) UnmarshalJSON(data []byte) error {
	aux := struct {
		time   int64
		open   string
		high   string
		low    string
		close  string
		vwap   string
		volume string
		count  int
	}{}

	tmpAux := []interface{}{
		&aux.time,
		&aux.open,
		&aux.high,
		&aux.low,
		&aux.close,
		&aux.vwap,
		&aux.volume,
		&aux.count,
	}
	err := json.Unmarshal(data, &tmpAux)
	if err != nil {
		return err
	}

	o.Time = time.Unix(aux.time, 0)

	o.Open, err = strconv.ParseFloat(aux.open, 64)
	if err != nil {
		return err
	}

	o.High, err = strconv.ParseFloat(aux.high, 64)
	if err != nil {
		return err
	}

	o.Low, err = strconv.ParseFloat(aux.low, 64)
	if err != nil {
		return err
	}

	o.Close, err = strconv.ParseFloat(aux.close, 64)
	if err != nil {
		return err
	}

	o.Vwap, err = strconv.ParseFloat(aux.vwap, 64)
	if err != nil {
		return err
	}

	o.Volume, err = strconv.ParseFloat(aux.volume, 64)
	if err != nil {
		return err
	}

	o.Count = aux.count

	return nil
}
