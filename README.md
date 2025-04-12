# optional

[![PkgGoDev](https://pkg.go.dev/badge/github.com/kiteggrad/optional)](https://pkg.go.dev/github.com/kiteggrad/optional)
[![Go Report Card](https://goreportcard.com/badge/github.com/kiteggrad/optional)](https://goreportcard.com/report/github.com/kiteggrad/optional)
[![codecov](https://codecov.io/gh/kiteggrad/optional/branch/main/graph/badge.svg?token=CKNT814QRN)](https://codecov.io/gh/kiteggrad/optional)

`optional` is a library for working with optional values in Go, designed to avoid using pointers for optional fields.

- [optional](#optional)
	- [Why Use Optional?](#why-use-optional)
	- [Installation](#installation)
	- [Types Overview](#types-overview)
		- [Type T - For Comparable Types](#type-t-for-comparable-types)
		- [Type A - For Any Types](#type-a-for-any-types)
	- [Usage Example](#usage-example)
	- [Examples with Non-Comparable Types](#examples-with-non-comparable-types)
	- [Features and Methods](#features-and-methods)
		- [Constructor Functions](#constructor-functions)
		- [Core Methods](#core-methods)
	- [Best Practices](#best-practices)
	- [Contributing](#contributing)
	- [License](#license)

## Why Use Optional?

When working with optional values in Go, developers often use pointers. However, this approach has drawbacks:

- Pointers can be nil, leading to potential nil pointer dereference panics
- It's hard to distinguish between "value is not set" and "value is set to empty"
- The code becomes littered with nil checks

This library solves these problems by providing a clean, type-safe way to work with optional values.

## Installation

```bash
go get github.com/kiteggrad/optional
```

## Types Overview

### Type T - For Comparable Types

`T[V comparable]` is designed for basic Go types that support comparison:
- Primitives (int, string, bool, etc.)
- Types that implement the comparable constraint

### Type A - For Any Types

`A[V any]` can be used with any Go type, including:
- Custom structs that don't implement comparable
- Slices, maps, and other non-comparable types
- Complex types that need optional treatment

While `A[V]` can work with reference types (slices, maps, channels), it's often better to use these types directly and check for nil, which is more idiomatic in Go. The real value of `A[V]` is when working with custom non-comparable structs that contain maps, slices, or other non-comparable fields, where simple nil checks aren't applicable.

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
	// Initialize some config with some values
	cfg := Config{
		Host:       optional.NewSetNotEmpty(getSomeHost()),
		Port:       optional.NewSetNotEmpty("8080"),
		RetryCount: optional.NewSet(0),
	}
	fmt.Println(cfg.Host.IsSet())       // Output: true / false
	fmt.Println(cfg.Port.IsSet())       // Output: true
	fmt.Println(cfg.RetryCount.IsSet()) // Output: true

	// Set a value
	cfg.Host = cfg.Host.Set("localhost")
	fmt.Println(cfg.Host.Value()) // Output: localhost
	fmt.Println(cfg.Host.IsSet()) // Output: true

	// Unset the value
	cfg.Host = cfg.Host.Unset()
	fmt.Println(cfg.Host.IsSet()) // Output: false
}

func getSomeHost() string {
	if rand.Int()%2 == 0 {
		return ""
	}

	return "localhost"
}
```

## Examples with Non-Comparable Types

```go
package main

import (
	"fmt"

	"github.com/kiteggrad/optional"
)

// Custom non-comparable struct (contains map field)
type UserPreferences struct {
	Theme       string
	Permissions map[string]bool // makes the struct non-comparable
	Language    optional.T[string]
}

// Config demonstrating proper use of optional.A
type Config struct {
	// For non-comparable structs, use optional.A
	Preferences optional.A[UserPreferences]

	// For reference types, better to use them directly
	// and check for nil instead of using optional
	AllowedHosts []string // can check if nil directly
}

var defaultPrefs = UserPreferences{
	Theme: "light",
	Permissions: map[string]bool{
		"admin": false,
		"edit":  true,
	},
}

func main() {
	// Initialize config with non-comparable struct
	cfg := Config{
		Preferences: optional.NewASet(UserPreferences{
			Theme: "dark",
			Permissions: map[string]bool{
				"admin": true,
				"edit":  true,
			},
		}),
		AllowedHosts: []string{"localhost"},
	}

	// SetDefault() to provide default values
	// without overriding existing ones.
	prefs := cfg.Preferences.SetDefault(defaultPrefs).Value()
	fmt.Println("Theme:", prefs.Theme)                // Theme: dark
	fmt.Println("Admin:", prefs.Permissions["admin"]) // Admin: false

	// SetDefault() will overwrite the value if it's not set
	language := cfg.Preferences.Language.SetDefault("en").Value()
	fmt.Println("Language:", language) // Language: en

	// With reference types, just check for nil directly
	if cfg.AllowedHosts != nil {
		fmt.Println("Hosts:", cfg.AllowedHosts)
	}
}
```

## Features and Methods

### Constructor Functions

- **NewSet(value)** - Creates an optional value with the provided value and marks it as set
- **NewSetNotEmpty(value)** - Creates an optional value only if the provided value is not empty
- **NewASet(value)** - Creates an optional value for any type (including non-comparable types)
- **NewFromPtr(ptr)** - Creates an optional value from a pointer (nil ptr = unset value)
- **NewAFromPtr(ptr)** - Creates an optional value of any type from a pointer

### Core Methods

- **Set(value)** - Sets a value and marks it as set
- **SetNotEmpty(value)** - Sets a value only if it's not empty
- **SetAuto(value)** - Sets a value and automatically marks it as unset if empty
- **Unset()** - Unsets a value
- **IsSet()** - Checks if a value is set
- **Value()** - Gets the current value (returns zero value if unset)
- **MustValue()** - Gets the value or panics if it's not set
- **IsEmpty()** - Checks if a value is empty
- **SetDefault(value)** - Sets a default value if the current value is not set

## Best Practices

- Use `T[V]` for basic types like string, int, etc.
- Use `A[V]` for complex types, slices, maps, or custom structs
- Check `IsSet()` before accessing values with `MustValue()`
- Use `SetDefault()` to provide default values without overriding existing ones
- For field-level defaults, use the pattern shown in `ExampleT_SetDefault_in_struct()`

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

See [LICENSE](./LICENSE) file for details.
