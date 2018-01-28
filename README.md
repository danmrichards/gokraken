# Go Kraken [![License](http://img.shields.io/badge/license-mit_bsd-blue.svg)](https://raw.githubusercontent.com/danmrichards/gokraken/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/danmrichards/gokraken)](https://goreportcard.com/report/github.com/danmrichards/gokraken)
A Go API client for the [Kraken](https://www.kraken.com) cryptocurrency exchange

> Note: This is stupidly early in development. Don't use this please...no really don't

## Documentation
Documentation is broken down by package
* [Main](https://godoc.org/github.com/danmrichards/gokraken)
* [Public](https://godoc.org/github.com/danmrichards/gokraken/public)

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
- [ ] Private API calls working
- [ ] Implement public market data endpoints
- [ ] Implement private user data endpoints
- [ ] Implement private user trading endpoints
- [ ] Implement private user funding endpoints
- [ ] Test all the things!