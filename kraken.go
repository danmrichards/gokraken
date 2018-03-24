package gokraken

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	// APIBaseURL is the base URI of the Kraken API.
	APIBaseURL = "https://api.kraken.com"

	// APIKeyHeader is the header to send the API key to Kraken in.
	APIKeyHeader = "API-Key"

	// APINonceParam is the parameter to send the nonce to Kraken in.
	APINonceParam = "nonce"

	// APIPublicNamespace is the name of the public Kraken API namespace.
	APIPublicNamespace = "public"

	// APIPrivateNamespace is the name of the private Kraken API namespace.
	APIPrivateNamespace = "private"

	// APISignHeader is the header to send the API signature to Kraken in.
	APISignHeader = "API-Sign"

	// APIVersion is the current version number of the Kraken API.
	APIVersion = 0

	// ClientVersion is the current version of this client.
	ClientVersion = "0.1.0"
)

var (
	// UserAgent is the user agent string applied to all requests.
	UserAgent = fmt.Sprintf("Go Kraken %s (github.com/danmrichards/gokraken)", ClientVersion)
)

// Kraken is responsible for all communication with the Kraken API.
type Kraken struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
	Market     *Market
	UserData   *UserData
	Trading    *Trading
	Funding    *Funding
	PrivateKey string
}

// New returns a new Kraken object with a default HTTP client.
func New() *Kraken {
	k := &Kraken{
		HTTPClient: &http.Client{},
	}
	k.initServices()

	return k
}

// NewWithAuth returns a new Kraken object with the authentication
// credentials required for private api endpoints.
func NewWithAuth(apiKey, privateKey string) *Kraken {
	k := &Kraken{
		APIKey:     apiKey,
		PrivateKey: privateKey,
		HTTPClient: &http.Client{},
	}
	k.initServices()

	return k
}

// NewWithHTTPClient returns a new Kraken object with a custom HTTP client.
func NewWithHTTPClient(httpClient *http.Client) *Kraken {
	k := &Kraken{
		HTTPClient: httpClient,
	}
	k.initServices()

	return k
}

// Initialise services for the Kraken api client.
func (k *Kraken) initServices() {
	k.Market = &Market{k}
	k.UserData = &UserData{k}
	k.Trading = &Trading{k}
	k.Funding = &Funding{k}
}

// GetBaseURL returns the base URI of the Kraken API.
// If the BaseURL value is not set on the Kraken struct the constant APIBaseURL
// will be returned instead.
func (k *Kraken) GetBaseURL() string {
	if k.BaseURL == "" {
		return APIBaseURL
	}

	return k.BaseURL
}

// Call performs a request against the Kraken API.
func (k *Kraken) Call(req *http.Request) (res *Response, err error) {
	apiResp, err := k.HTTPClient.Do(req)
	if err != nil {
		return
	}

	err = bindJSON(apiResp.Body, &res)
	if err != nil {
		return
	}

	return
}

// Dial prepares a request to send to the Kraken API.
func (k *Kraken) Dial(ctx context.Context, method, resource string, body url.Values) (req *http.Request, err error) {
	req, err = http.NewRequest(method, k.ResourceURL(APIPublicNamespace, resource), strings.NewReader(body.Encode()))
	if err != nil {
		return
	}

	// Apply the context to the request to allow it to be cancelled.
	req = req.WithContext(ctx)

	req.Header.Add("User-Agent", UserAgent)

	return
}

// DialWithAuth prepares an authenticated request to send to the Kraken API.
func (k *Kraken) DialWithAuth(ctx context.Context, method, resource string, body url.Values) (req *http.Request, err error) {
	if k.APIKey == "" {
		err = errors.New("missing or invalid api key")
		return
	}

	if k.PrivateKey == "" {
		err = errors.New("missing or invalid private key")
		return
	}

	// Create an empty map if nil passed.
	if body == nil {
		body = url.Values{}
	}

	// Decode the private key.
	secret, err := base64.StdEncoding.DecodeString(k.PrivateKey)
	if err != nil {
		err = fmt.Errorf("could not decode private key: %s", err)
		return
	}

	// Create a unique nonce value for this request.
	// https://www.kraken.com/en-gb/help/api#general-usage
	nonce := time.Now().UnixNano()
	body.Set(APINonceParam, strconv.FormatInt(nonce, 10))

	// Generate the request.
	req, err = http.NewRequest(method, k.ResourceURL(APIPrivateNamespace, resource), strings.NewReader(body.Encode()))
	if err != nil {
		return
	}

	// Apply the context to the request to allow it to be cancelled.
	req = req.WithContext(ctx)

	// Generate signature.
	signature := &Signature{
		APISecret: secret,
		Data:      body,
		URI:       k.ResourceURI(APIPrivateNamespace, resource),
	}

	// Apply headers.
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add(APIKeyHeader, k.APIKey)
	req.Header.Add(APISignHeader, signature.Generate())

	return
}

// ResourceURI returns the URI path for the given api resource.
func (k *Kraken) ResourceURI(namespace, resource string) string {
	return fmt.Sprintf(
		"/%d/%s/%s",
		APIVersion,
		namespace,
		resource,
	)
}

// ResourceURL returns a fully qualified URI for the given api resource.
func (k *Kraken) ResourceURL(namespace, resource string) string {
	return fmt.Sprintf(
		"%s%s",
		k.GetBaseURL(),
		k.ResourceURI(namespace, resource),
	)
}
