package optional_test

import (
	"fmt"

	"github.com/kiteggrad/optional"
)

func ExampleNew() {
	// Creating optional with value set
	opt1 := optional.New("AppName", true)
	fmt.Println(opt1.IsSet())
	fmt.Println(opt1.Value())

	// Creating optional without value set
	opt2 := optional.New("AppName", false)
	fmt.Println(opt2.IsSet())
	fmt.Println(opt2.Value())

	// Output:
	// true
	// AppName
	// false
	//
}

func ExampleNewPtr() {
	// With non-nil pointer
	value := "Hello"
	opt1 := optional.NewPtr(&value)
	fmt.Println(opt1.IsSet())
	fmt.Println(opt1.Value())

	// With nil pointer
	var nilPtr *string
	opt2 := optional.NewPtr(nilPtr)
	fmt.Println(opt2.IsSet())

	// With different types
	num := 42
	opt3 := optional.NewPtr(&num)
	fmt.Println(opt3.IsSet())
	fmt.Println(opt3.Value())

	// Output:
	// true
	// Hello
	// false
	// true
	// 42
}

func ExampleT_IsSet() {
	var opt optional.T[int]
	fmt.Println(opt.IsSet())

	opt = optional.New(8080, true)
	fmt.Println(opt.IsSet())

	// Output:
	// false
	// true
}

func ExampleT_Value() {
	var opt optional.T[int]
	fmt.Println(opt.Value())

	opt = optional.New(8080, true)
	fmt.Println(opt.Value())

	// Output:
	// 0
	// 8080
}

func ExampleT_MustValue() {
	opt := optional.New(8080, true)
	fmt.Println(opt.MustValue())

	// Output: 8080
}

func ExampleT_Ptr() {
	var opt optional.T[int]
	ptr := opt.Ptr()
	fmt.Println(ptr == nil)

	opt = optional.New(8080, true)
	ptr = opt.Ptr()
	fmt.Println(ptr != nil)
	fmt.Println(*ptr)

	// Output:
	// true
	// true
	// 8080
}

func ExampleT_MustValue_panic() {
	var opt optional.T[string]

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	opt.MustValue() // This will panic

	// Output: Recovered from panic: value is not set
}

func ExampleNew_customTypes() {
	type User struct {
		ID   int
		Name string
	}

	// Custom struct
	user := User{ID: 1, Name: "John"}
	opt1 := optional.New(user, true)
	fmt.Println(opt1.IsSet())
	fmt.Println(opt1.Value())

	// Slice
	slice := []string{"hello", "world"}
	opt2 := optional.New(slice, len(slice) > 0)
	fmt.Println(opt2.IsSet())
	fmt.Println(opt2.Value())

	// Output:
	// true
	// {1 John}
	// true
	// [hello world]
}

func ExampleNewPtr_configuration() {
	type DatabaseConfig struct {
		Host     string
		Port     optional.T[int]
		Username optional.T[string]
		SSL      optional.T[bool]
	}

	// Function to create config from pointers
	createConfig := func(host string, port *int, username *string, ssl *bool) DatabaseConfig {
		return DatabaseConfig{
			Host:     host,
			Port:     optional.NewPtr(port),
			Username: optional.NewPtr(username),
			SSL:      optional.NewPtr(ssl),
		}
	}

	// Usage
	port := 5432
	username := "admin"
	config := createConfig("localhost", &port, &username, nil)

	fmt.Println("Host:", config.Host)
	fmt.Println("Port set:", config.Port.IsSet())
	fmt.Println("Username set:", config.Username.IsSet())
	fmt.Println("SSL set:", config.SSL.IsSet())

	// Output:
	// Host: localhost
	// Port set: true
	// Username set: true
	// SSL set: false
}

func ExampleT_Ptr_apiIntegration() {
	type APIPayload struct {
		Name string `json:"name"`
		Age  *int   `json:"age,omitempty"`
	}

	type User struct {
		Name string
		Age  optional.T[int]
	}

	user := User{
		Name: "John",
		Age:  optional.New(30, true),
	}

	// Convert to API payload using Ptr()
	payload := APIPayload{
		Name: user.Name,
		Age:  user.Age.Ptr(), // Returns *int or nil
	}

	fmt.Println("Name:", payload.Name)
	fmt.Println("Age not nil:", payload.Age != nil)

	// Output:
	// Name: John
	// Age not nil: true
}
