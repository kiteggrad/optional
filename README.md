# optional

[![PkgGoDev](https://pkg.go.dev/badge/github.com/kiteggrad/optional)](https://pkg.go.dev/github.com/kiteggrad/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/kiteggrad/optional)](https://goreportcard.com/report/github.com/kiteggrad/optional)

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

	"github.com/kiteggrad/optional"
)

func main() {
	var opt optional.T[string]

	// Check if the value is set
	fmt.Println(opt.IsSet()) // Output: false

	// Set a value
	opt = opt.Set("Hello, World!")
	fmt.Println(opt.Value()) // Output: Hello, World!

	// Check if the value is set
	fmt.Println(opt.IsSet()) // Output: true

	// Unset the value
	opt = opt.Unset()
	fmt.Println(opt.IsSet()) // Output: false
}
```

## Features

- Set a value (`Set`, `SetNotEmpty`, `SetAuto`)
- Unset a value (`Unset`)
- Check if a value is set (`IsSet`)
- Retrieve a value with checks (`Value`, `MustValue`)
- Set a default value (`SetDefault`)
- Check if a value is empty (`IsEmpty`)
