package gokraken

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// bindJSON takes a body from a readcloser and binds it to the target interface.
func bindJSON(rc io.ReadCloser, target interface{}) error {
	defer rc.Close()

	// Reading all of body for Unmarshal rather than
	// using Decoder for this as the response we get
	// isn't a JSON stream.
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, target)
}
