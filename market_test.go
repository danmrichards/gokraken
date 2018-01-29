package gokraken

import (
	"context"
	"fmt"
	"testing"
)

func TestPublicMarket_Time(t *testing.T) {
	// TODO: Table test this with a mock api.

	k := New()

	resp, err := k.Market.Time(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("server unix time: %d\n", resp.UnixTime)
	fmt.Printf("service rfc1123 time: %s\n", resp.Rfc1123)
}
