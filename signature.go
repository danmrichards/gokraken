package gokraken

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"net/url"
)

// Signature represents a cryptographic signature for signing private Kraken
// API requests.
type Signature struct {
	APISecret []byte
	Data      url.Values
	URI       string
}

// Generate returns a message signature for use on private Kraken API endpoints.
// https://www.kraken.com/en-gb/help/api#general-usage
func (s *Signature) Generate() string {
	// SHA256 of nonce and request data.
	sha := sha256.New()
	sha.Write([]byte(s.Data.Get(APINonceParam) + s.Data.Encode()))
	shaSum := sha.Sum(nil)

	// HMAC of request URI and the SHA256 sum.
	mac := hmac.New(sha512.New, s.APISecret)
	mac.Write(append([]byte(s.URI), shaSum...))
	macSum := mac.Sum(nil)

	return base64.StdEncoding.EncodeToString(macSum)
}
