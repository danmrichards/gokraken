package gokraken

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	k := New()

	if k.HTTPClient == nil {
		t.Fatalf("%s: nil http client", t.Name())
	}

	if k.Market == nil {
		t.Fatalf("%s: nil Market service", t.Name())
	}

	if k.UserData == nil {
		t.Fatalf("%s: nil UserData service", t.Name())
	}
}

func TestNewWithAuth(t *testing.T) {
	apiKey := "foo"
	privateKey := "bar"

	k := NewWithAuth("foo", "bar")

	assert(apiKey, k.APIKey, t)
	assert(privateKey, k.PrivateKey, t)

	if k.HTTPClient == nil {
		t.Fatalf("%s: nil http client", t.Name())
	}

	if k.Market == nil {
		t.Fatalf("%s: nil Market service", t.Name())
	}

	if k.UserData == nil {
		t.Fatalf("%s: nil UserData service", t.Name())
	}
}

func TestNewWithHTTPClient(t *testing.T) {
	testClient := &http.Client{
		Timeout: 1 * time.Second,
	}

	k := NewWithHTTPClient(testClient)

	assert(testClient, k.HTTPClient, t)

	if k.Market == nil {
		t.Fatalf("%s: nil Market service", t.Name())
	}

	if k.UserData == nil {
		t.Fatalf("%s: nil UserData service", t.Name())
	}
}

func TestKraken_GetBaseUrl(t *testing.T) {
	cases := []struct {
		name    string
		baseURL string
	}{
		{
			name:    "with base url",
			baseURL: "api.foo.com",
		},
		{
			name: "without base url",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := New()
			k.BaseURL = c.baseURL

			if c.baseURL != "" {
				assert(c.baseURL, k.GetBaseUrl(), t)
			} else {
				assert(APIBaseUrl, k.GetBaseUrl(), t)
			}
		})
	}
}

func TestKraken_PrepareRequest(t *testing.T) {
	exampleResp := Response{
		Result: map[string]interface{}{
			"foo": "bar",
			"baz": "qux",
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response, _ := json.Marshal(exampleResp)
		w.Write(response)
	}))

	defer ts.Close()

	k := New()
	k.BaseURL = ts.URL

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := k.Call(req)
	if err != nil {
		t.Fatal(err)
	}

	assert(&exampleResp, resp, t)
}

func TestKraken_Dial(t *testing.T) {
	cases := []struct {
		name        string
		baseURL     string
		method      string
		resource    string
		body        url.Values
		expectedURL string
		encodedBody string
	}{
		{
			name:        "with base url",
			baseURL:     "api.foo.com",
			method:      http.MethodGet,
			resource:    "bar",
			expectedURL: "https://api.foo.com/0/public/bar",
		},
		{
			name:        "without base url",
			method:      http.MethodGet,
			resource:    "bar",
			expectedURL: "https://api.kraken.com/0/public/bar",
		},
		{
			name:     "POST",
			baseURL:  "api.foo.com",
			method:   http.MethodPost,
			resource: "bar",
			body: url.Values{
				"foo": []string{"bar"},
				"baz": []string{"qux"},
			},
			expectedURL: "https://api.foo.com/0/public/bar",
			encodedBody: "baz=qux&foo=bar",
		},
		{
			name:     "PUT",
			baseURL:  "api.foo.com",
			method:   http.MethodPut,
			resource: "bar",
			body: url.Values{
				"foo": []string{"bar"},
				"baz": []string{"qux"},
			},
			expectedURL: "https://api.foo.com/0/public/bar",
			encodedBody: "baz=qux&foo=bar",
		},
		{
			name:     "PATCH",
			baseURL:  "api.foo.com",
			method:   http.MethodPut,
			resource: "bar",
			body: url.Values{
				"baz": []string{"qux"},
			},
			expectedURL: "https://api.foo.com/0/public/bar",
			encodedBody: "baz=qux",
		},
		{
			name:        "DELETE",
			baseURL:     "api.foo.com",
			method:      http.MethodDelete,
			resource:    "bar",
			expectedURL: "https://api.foo.com/0/public/bar",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := New()
			k.BaseURL = c.baseURL

			req, err := k.Dial(context.Background(), c.method, c.resource, c.body)
			if err != nil {
				t.Fatal(err)
			}

			assert(c.method, req.Method, t)
			assert(c.expectedURL, req.URL.String(), t)
			assert(UserAgent, req.Header.Get("User-Agent"), t)

			buf := new(bytes.Buffer)
			buf.ReadFrom(req.Body)
			assert(c.encodedBody, buf.String(), t)
		})
	}
}

func TestKraken_DialWithAuth(t *testing.T) {
	cases := []struct {
		name           string
		apiKey         string
		privateKey     string
		baseURL        string
		method         string
		resource       string
		body           url.Values
		expectedURL    string
		expectedError  string
		decodedPrivate string
		bodyRegex      *regexp.Regexp
	}{
		{
			name:          "no private key",
			apiKey:        "baz2345qux",
			expectedError: "missing or invalid private key",
		},
		{
			name:          "no api key",
			privateKey:    "Zm9vMTIzNGJhcg==",
			expectedError: "missing or invalid api key",
		},
		{
			name:          "no private key or api key",
			expectedError: "missing or invalid api key",
		},
		{
			name:        "with base url",
			privateKey:  "Zm9vMTIzNGJhcg==",
			apiKey:      "baz2345qux",
			baseURL:     "api.foo.com",
			method:      http.MethodGet,
			resource:    "bar",
			expectedURL: "https://api.foo.com/0/private/bar",
			bodyRegex:   regexp.MustCompile("nonce=(.*)"),
		},
		{
			name:        "without base url",
			privateKey:  "Zm9vMTIzNGJhcg==",
			apiKey:      "baz2345qux",
			method:      http.MethodGet,
			resource:    "bar",
			expectedURL: "https://api.kraken.com/0/private/bar",
			bodyRegex:   regexp.MustCompile("nonce=(.*)"),
		},
		{
			name:       "POST",
			privateKey: "Zm9vMTIzNGJhcg==",
			apiKey:     "baz2345qux",
			baseURL:    "api.foo.com",
			method:     http.MethodPost,
			resource:   "bar",
			body: url.Values{
				"foo": []string{"bar"},
				"baz": []string{"qux"},
			},
			expectedURL: "https://api.foo.com/0/private/bar",
			bodyRegex:   regexp.MustCompile("baz=qux&foo=bar&nonce=(.*)"),
		},
		{
			name:       "PUT",
			privateKey: "Zm9vMTIzNGJhcg==",
			apiKey:     "baz2345qux",
			baseURL:    "api.foo.com",
			method:     http.MethodPut,
			resource:   "bar",
			body: url.Values{
				"foo": []string{"bar"},
				"baz": []string{"qux"},
			},
			expectedURL: "https://api.foo.com/0/private/bar",
			bodyRegex:   regexp.MustCompile("baz=qux&foo=bar&nonce=(.*)"),
		},
		{
			name:       "PATCH",
			privateKey: "Zm9vMTIzNGJhcg==",
			apiKey:     "baz2345qux",
			baseURL:    "api.foo.com",
			method:     http.MethodPut,
			resource:   "bar",
			body: url.Values{
				"baz": []string{"qux"},
			},
			expectedURL: "https://api.foo.com/0/private/bar",
			bodyRegex:   regexp.MustCompile("baz=qux&nonce=(.*)"),
		},
		{
			name:        "DELETE",
			privateKey:  "Zm9vMTIzNGJhcg==",
			apiKey:      "baz2345qux",
			baseURL:     "api.foo.com",
			method:      http.MethodDelete,
			resource:    "bar",
			expectedURL: "https://api.foo.com/0/private/bar",
			bodyRegex:   regexp.MustCompile("nonce=(.*)"),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := NewWithAuth(c.apiKey, c.privateKey)
			k.BaseURL = c.baseURL

			req, err := k.DialWithAuth(context.Background(), c.method, c.resource, c.body)
			if err != nil && err.Error() != c.expectedError {
				t.Fatal(err)
			}

			if c.expectedError == "" {
				assert(c.method, req.Method, t)
				assert(c.expectedURL, req.URL.String(), t)
				assert(UserAgent, req.Header.Get("User-Agent"), t)
				assert(c.apiKey, req.Header.Get(APIKeyHeader), t)

				buf := new(bytes.Buffer)
				buf.ReadFrom(req.Body)

				if !c.bodyRegex.Match(buf.Bytes()) {
					t.Fatalf(
						"%s: could not match regex %s to body %s",
						t.Name(),
						c.bodyRegex.String(),
						buf.String(),
					)
				}
			}
		})
	}
}

func TestKraken_ResourceURI(t *testing.T) {
	cases := []struct {
		name           string
		namespace      string
		resource       string
		expectedOutput string
	}{
		{
			name:           "public",
			namespace:      APIPublicNamespace,
			resource:       "bar",
			expectedOutput: "/0/public/bar",
		},
		{
			name:           "private",
			namespace:      APIPrivateNamespace,
			resource:       "bar",
			expectedOutput: "/0/private/bar",
		},
		{
			name:           "custom",
			namespace:      "foo",
			resource:       "bar",
			expectedOutput: "/0/foo/bar",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := New()

			assert(c.expectedOutput, k.ResourceURI(c.namespace, c.resource), t)
		})
	}
}

func TestKraken_ResourceURL(t *testing.T) {
	cases := []struct {
		name           string
		baseURL        string
		namespace      string
		resource       string
		expectedOutput string
	}{
		{
			name:           "custom base url",
			baseURL:        "api.foo.com",
			namespace:      APIPublicNamespace,
			resource:       "bar",
			expectedOutput: "https://api.foo.com/0/public/bar",
		},
		{
			name:           "public",
			namespace:      APIPublicNamespace,
			resource:       "bar",
			expectedOutput: "https://api.kraken.com/0/public/bar",
		},
		{
			name:           "private",
			namespace:      APIPrivateNamespace,
			resource:       "bar",
			expectedOutput: "https://api.kraken.com/0/private/bar",
		},
		{
			name:           "custom",
			namespace:      "foo",
			resource:       "bar",
			expectedOutput: "https://api.kraken.com/0/foo/bar",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := New()
			k.BaseURL = c.baseURL

			assert(c.expectedOutput, k.ResourceURL(c.namespace, c.resource), t)
		})
	}
}
