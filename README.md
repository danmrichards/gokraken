# Go Kraken [![GoDoc](https://godoc.org/github.com/danmrichards/gokraken?status.svg)](https://godoc.org/github.com/danmrichards/gokraken) [![License](http://img.shields.io/badge/license-mit_bsd-blue.svg)](https://raw.githubusercontent.com/danmrichards/gokraken/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/danmrichards/gokraken)](https://goreportcard.com/report/github.com/danmrichards/gokraken) [![Build Status](https://travis-ci.org/danmrichards/gokraken.svg?branch=master)](https://travis-ci.org/danmrichards/gokraken)
A Go API client for the [Kraken](https://www.kraken.com) cryptocurrency exchange

> Note: This is stupidly early in development. Don't use it please...no really don't

## Usage
### Public API
```go
package main

import (
	"context"
	"fmt"
	"log"
	
	"github.com/danmrichards/gokraken"
)

func main() {
	kraken := gokraken.New()
	
	resp, err := kraken.Market.Time(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("server unix time: %d\n", resp.UnixTime)
	fmt.Printf("service rfc1123 time: %s\n", resp.Rfc1123)
}
```

### Private API
```go
package main

import (
	"context"
	"fmt"
	"log"
	
	"github.com/danmrichards/gokraken"
)

func main() {
	kraken := gokraken.NewWithAuth("API_KEY", "PRIVATE_KEY")
	
	resp, err := kraken.UserData.Balance(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	
	for currency, balance := range resp {
	    fmt.Printf("%s: %f'n", currency, balance)
	}
}
```

## Advanced Usage
Coming soon

## Roadmap
- [x] Base repo structure
- [x] Public API calls working
- [x] Private API calls working
- [x] Travis CI
- [ ] Implement public market data endpoints
- [ ] Implement private user data endpoints
- [ ] Implement private user trading endpoints
- [ ] Implement private user funding endpoints
- [ ] Test all the things!