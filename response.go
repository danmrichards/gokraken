package gokraken

import "encoding/json"

// Response represents a response from the Kraken API.
type Response struct {
	Error  []string    `json:"error"`
	Result interface{} `json:"result"`
}

// ExtractResult extracts the result from a Kraken API response into the
// destination parameter.
func (r *Response) ExtractResult(dst interface{}) error {
	resultJSON, err := json.Marshal(r.Result)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resultJSON, &dst)
	if err != nil {
		return err
	}

	return nil
}
