package public

import (
	"context"
	"fmt"
	"testing"

	"github.com/danmrichards/gokraken"
)

func TestPublicMarket_Time(t *testing.T) {
	// TODO: Table test this with a mock api.

	k := gokraken.New()
	p := &Market{Client: k}

	resp, err := p.Time(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("server unix time: %d\n", resp.UnixTime)
	fmt.Printf("service rfc1123 time: %s\n", resp.Rfc1123)
}
