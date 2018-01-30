package gokraken

import "testing"

func TestResponse_ExtractResult(t *testing.T) {
	exampleResp := Response{
		Result: map[string]interface{}{
			"foo": "bar",
			"baz": "qux",
		},
	}

	var dst map[string]interface{}
	exampleResp.ExtractResult(&dst)

	assert(exampleResp.Result, dst, t)
}
