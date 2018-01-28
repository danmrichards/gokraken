package gokraken

import (
	"context"
	"fmt"
	"net/http"
)

const (
	// ApiBaseUrl is the base URL of the Kraken API.
	ApiBaseUrl = "api.kraken.com"

	// ApiPublicNamespace is the name of the public Kraken API namespace.
	ApiPublicNamespace = "public"

	// ApiPrivateNamespace is the name of the private Kraken API namespace.
	ApiPrivateNamespace = "private"

	// ApiProtocol is the protocol in which to access the Kraken API.
	ApiProtocol = "https"

	// ApiVersion is the current version number of the Kraken API.
	ApiVersion = 0

	// ClientVersion is the current version of this client.
	ClientVersion = "0.1.0"
)

var (
	// UserAgent is the user agent string applied to all requests.
	UserAgent = fmt.Sprintf("Go Kraken %s (github.com/danmrichards/gokraken)", ClientVersion)
)

// KrakenClient is the interface that all Kraken clients must implement.
type KrakenClient interface {
	// Performs a request against the Kraken API.
	Do(req *http.Request) (resp *Response, err error)

	// Prepares a request to send to the Kraken API.
	PrepareRequest(ctx context.Context, method, resource string, body interface{}) (req *http.Request, err error)

	// Prepares an authenticated request to send to the Kraken API.
	PrepareAuthRequest(ctx context.Context, method, resource string, body interface{}) (req *http.Request, err error)
}

// Kraken is responsible for all communication with the Kraken API.
type Kraken struct {
	ApiKey     string
	BaseUrl    string
	HTTPClient *http.Client
	PrivateKey string
}

// New returns a new Kraken object with a default HTTP client.
func New() *Kraken {
	return &Kraken{
		HTTPClient: &http.Client{},
	}
}

// NewWithAuth returns a new Kraken object with the authentication
// credentials required for private api endpoints.
func NewWithAuth(apiKey, privateKey string) *Kraken {
	return &Kraken{
		ApiKey:     apiKey,
		PrivateKey: privateKey,
		HTTPClient: &http.Client{},
	}
}

// NewWithHTTPClient returns a new Kraken object with a custom HTTP client.
func NewWithHTTPClient(httpClient *http.Client) *Kraken {
	return &Kraken{
		HTTPClient: httpClient,
	}
}

// GetBaseUrl returns the base URL of the Kraken API.
// If the BaseUrl value is not set on the Kraken struct the constant ApiBaseUrl
// will be returned instead.
func (k *Kraken) GetBaseUrl() string {
	if k.BaseUrl == "" {
		return ApiBaseUrl
	}

	return k.BaseUrl
}

// Performs a request against the Kraken API.
func (k *Kraken) Do(req *http.Request) (resp *Response, err error) {
	apiResp, err := k.HTTPClient.Do(req)
	if err != nil {
		// TODO: Error logging here.
		return
	}

	err = bindJSON(apiResp.Body, &resp)
	if err != nil {
		// TODO: Error logging here.
		return
	}

	return
}

// Prepares a new request to send to the Kraken API.
func (k *Kraken) PrepareRequest(ctx context.Context, method, resource string, body interface{}) (req *http.Request, err error) {
	req, err = http.NewRequest(method, k.ResourceUrl(ApiPublicNamespace, resource), nil)
	if err != nil {
		return
	}

	// Apply the context to the request to allow it to be cancelled.
	req = req.WithContext(ctx)

	req.Header.Add("User-Agent", UserAgent)

	return
}

// Prepares an authenticated request to send to the Kraken API.
func (k *Kraken) PrepareAuthRequest(ctx context.Context, method, resource string, body interface{}) (req *http.Request, err error) {
	req, err = k.PrepareRequest(ctx, method, resource, body)
	if err != nil {
		return
	}

	// TODO: Generate signature.

	// TODO: Apply auth headers.

	return
}

// ResourceUrl returns a fully qualified URL for the given api resource.
func (k *Kraken) ResourceUrl(namespace, resource string) string {
	return fmt.Sprintf(
		"%s://%s/%d/%s/%s",
		ApiProtocol,
		k.GetBaseUrl(),
		ApiVersion,
		namespace,
		resource,
	)
}
