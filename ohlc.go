package gokraken

import (
	"encoding/json"
	"errors"
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
	Open   string    `json:"open"`
	High   string    `json:"high"`
	Low    string    `json:"low"`
	Close  string    `json:"close"`
	Vwap   string    `json:"vwap"`
	Volume string    `json:"volume"`
	Count  int       `json:"count"`
}

// UnmarshalJSON parses the JSON-encoded data and stores the result
// in the value pointed to by o.
//
// Due to the nature of the Kraken API response we iterate over a map of
// interface{} keyed by string. Each key is then checked and the value type
// asserted accordingly.
func (o *OhlcData) UnmarshalJSON(data []byte) error {
	var aux []interface{}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	for k := range aux {
		switch k {
		case 0:
			timeUnix, ok := aux[k].(float64)
			if !ok {
				return errors.New("could not assert time as int64")
			}
			o.Time = time.Unix(int64(timeUnix), 0)
		case 1:
			open, ok := aux[k].(string)
			if !ok {
				return errors.New("could not assert open as string")
			}
			o.Open = open
		case 2:
			high, ok := aux[k].(string)
			if !ok {
				return errors.New("could not assert high as string")
			}
			o.High = high
		case 3:
			low, ok := aux[k].(string)
			if !ok {
				return errors.New("could not assert low as string")
			}
			o.Low = low
		case 4:
			closeVal, ok := aux[k].(string)
			if !ok {
				return errors.New("could not assert close as string")
			}
			o.Close = closeVal
		case 5:
			vwap, ok := aux[k].(string)
			if !ok {
				return errors.New("could not assert vwap as string")
			}
			o.Vwap = vwap
		case 6:
			volume, ok := aux[k].(string)
			if !ok {
				return errors.New("could not assert volume as string")
			}
			o.Volume = volume
		case 7:
			count, ok := aux[k].(float64)
			if !ok {
				return errors.New("could not assert count as int64")
			}
			o.Count = int(count)
		}
	}

	return nil
}
