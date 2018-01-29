package gokraken

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestBindJSON(t *testing.T) {
	type testData struct {
		Foo string `json:"foo"`
		Baz string `json:"baz"`
	}

	input := `{"foo": "bar", "baz": "qux"}`
	expectedOutput := testData{
		Foo: "bar",
		Baz: "qux",
	}

	var output testData
	err := bindJSON(ioutil.NopCloser(bytes.NewBuffer([]byte(input))), &output)
	if err != nil {
		t.Fatalf("could not bind json: %s", err)
	}

	assert(expectedOutput, output, t)
}
