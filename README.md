# optional

[![PkgGoDev](https://pkg.go.dev/badge/github.com/kiteggrad/optional)](https://pkg.go.dev/github.com/kiteggrad/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/kiteggrad/optional)](https://goreportcard.com/report/github.com/kiteggrad/optional)
[![codecov](https://codecov.io/gh/kiteggrad/optional/branch/main/graph/badge.svg?token=CKNT814QRN)](https://codecov.io/gh/kiteggrad/optional)

A lightweight, type-safe library for optional values in Go. Say goodbye to nil pointer panics and confusing zero values.

## Why?

Traditional Go code uses pointers for optional fields, which leads to:
- 💥 Nil pointer dereference panics
- 🤔 Ambiguity between "not set" and "zero value"
- 🧩 Excessive nil checks cluttering your code
- 🔄 Increased pressure on garbage collection

`optional` provides a clean, immutable, generic solution.

## Installation

```bash
go get github.com/kiteggrad/optional/v2
```

## Quick Start

```go
import (
	"fmt"
	"flag"
	"github.com/kiteggrad/optional/v2"
)

var portFlag = flag.Int("port", 8080, "Server port") // *int

func main() {
	flag.Parse()

	// Create optional values
	var (
		host      = optional.New("localhost", true)      // Set
		port      = optional.NewPtr(portFlag)            // From pointer
		ssl       = optional.T[bool]{}                   // Not set
		password  = optional.New("ignored value", false) // Ignore value if not set
	)

	// Check before access
	if !port.IsSet() {
		panic("Port must be set")
	}

	// Safe access to unset value
	fmt.Println(password.Value()) // prints "" because not set, even though we provided a value

	// Convert back to pointer
	fmt.Println(password.Ptr()) // prints "<nil>" because not set
	fmt.Println(port.Ptr())	    // prints pointer to int value
}
```

See more examples in [optional_test.go](./optional_test.go).

## API

**Constructors:**
- `New(value, isSet)` - Create optional with explicit set state
- `NewPtr(ptr)` - Create from pointer (nil = not set)

**Methods:**
- `IsSet() bool` - Check if value is present
- `Value() V` - Get value (zero value if not set)
- `MustValue() V` - Get value or panic
- `Ptr() *V` - Get pointer or nil if not set

## License

[MIT](./LICENSE)
