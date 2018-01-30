package gokraken

import (
	"reflect"
	"testing"
)

// Test helper for asserting values are equal.
func assert(expected, actual interface{}, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("%s: expected: %#[2]v (%[2]T), but got %#[3]v (%[3]T)", t.Name(), expected, actual)
	}
}
