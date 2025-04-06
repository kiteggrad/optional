# optional

[![PkgGoDev](https://pkg.go.dev/badge/github.com/kiteggrad/optional)](https://pkg.go.dev/github.com/kiteggrad/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/kiteggrad/optional)](https://goreportcard.com/report/github.com/kiteggrad/optional)
[![codecov](https://codecov.io/gh/kiteggrad/optional/branch/main/graph/badge.svg?token=CKNT814QRN)](https://codecov.io/gh/kiteggrad/optional)

`optional` is a library for working with optional values in Go, designed to avoid using pointers for optional fields.

## Installation

```bash
go get github.com/kiteggrad/optional
```

## Usage Example

See more examples in [example_test.go](./example_test.go)

```go
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/kiteggrad/optional"
)

type Config struct {
	Host       optional.T[string]
	Port       optional.T[string]
	RetryCount optional.T[int]
}

func main() {
	// Initialize some config with som
	cfg := Config{
		Host:       optional.NewSetNotEmpty(getSomeHost()),
		Port:       optional.NewSetNotEmpty("8080"),
		RetryCount: optional.NewSet(0),
	}
	fmt.Println(cfg.Host.IsSet())       // Output: false / localhost
	fmt.Println(cfg.Port.IsSet())       // Output: true
	fmt.Println(cfg.RetryCount.IsSet()) // Output: true

	// Set a value
	cfg.Port = cfg.Host.Set("localhost")
	fmt.Println(cfg.Host.Value()) // Output: localhost
	fmt.Println(cfg.Host.IsSet()) // Output: true

	// Unset the value
	cfg.Port = cfg.Host.Unset()
	fmt.Println(cfg.Host.IsSet()) // Output: false
}

func getSomeHost() string {
	if rand.Int()%2 == 0 {
		return ""
	}

	return "localhost"
}
```

## Features

- Set a value (`Set`, `SetNotEmpty`, `SetAuto`)
- Unset a value (`Unset`)
- Check if a value is set (`IsSet`)
- Retrieve a value with checks (`Value`, `MustValue`)
- Set a default value (`SetDefault`)
- Check if a value is empty (`IsEmpty`)
- **Constructors for optional values**:
  - `NewSet` — creates an optional value with the provided value and marks it as set.
  - `NewSetNotEmpty` — creates an optional value only if the provided value is not empty; otherwise, it remains unset.

