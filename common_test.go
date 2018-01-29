package gokraken

import (
	"reflect"
	"testing"
)

// Test helper for asserting values are equal.
func assert(expected, actual interface{}, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %#[1]v (%[1]T), but got %#[2]v (%[2]T)", expected, actual)
	}
}
