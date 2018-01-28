# Go Kraken
A Go API client for the Kraken cryptocurrency exchange

> Note: This is stupidly early in development. Don't use this please...no really don't

## Basic Usage
```go
package main

import (
	"context"
	"fmt"
	"log"
	
	"github.com/danmrichards/gokraken"
	"github.com/danmrichards/gokraken/public"
)

func main() {
	kraken := gokraken.New()
	market := &public.Market{
		Client: kraken,
	} 
	
	resp, err := market.Time(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("server unix time: %d\n", resp.UnixTime)
	fmt.Printf("service rfc1123 time: %s\n", resp.Rfc1123)
}
```

## Advanced Usage
Coming soon

## Roadmap
- [x] Base repo structure
- [x] Public API calls working
- [ ] Implement public market data endpoints
- [ ] Implement private user data endpoints
- [ ] Implement private user trading endpoints
- [ ] Implement private user funding endpoints
- [ ] Test all the things!