package gokraken

import (
	"encoding/base64"
	"net/url"
	"testing"
)

func TestSignature_Generate(t *testing.T) {
	cases := []struct {
		name              string
		apiKey            string
		privateKey        string
		resource          string
		body              url.Values
		expectedSignature string
	}{
		{
			name:              "no body",
			apiKey:            "foo1234bar",
			privateKey:        "YmF6MjM0NXF1eA==",
			resource:          "foo",
			expectedSignature: "lZmH3F5gHvoNWoma9eQR7JhtEB/cnD9D6AvRyedJkjwE2lX6JQ/sdPs/C9Jmq4RILCUKHC+JAK5PGB1fz+4T8Q==",
		},
		{
			name:       "body",
			apiKey:     "foo1234bar",
			privateKey: "YmF6MjM0NXF1eA==",
			resource:   "foo",
			body: url.Values{
				"foo": []string{"bar"},
				"baz": []string{"qux"},
			},
			expectedSignature: "GxvTUUkwAsJZ4wTbzikFnaj+DhgxCFgjOQ2aV4HX7L1/9m5AyYlbz5jNz/WWwXJuyBIR9iZKRrlTkvlp3ggPwg==",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			k := NewWithAuth(c.apiKey, c.privateKey)

			secret, err := base64.StdEncoding.DecodeString(k.PrivateKey)
			if err != nil {
				t.Fatal(err)
			}

			signature := &Signature{
				APISecret: secret,
				Data:      c.body,
				URI:       k.ResourceURI(APIPrivateNamespace, c.resource),
			}

			assert(c.expectedSignature, signature.Generate(), t)
		})
	}
}
