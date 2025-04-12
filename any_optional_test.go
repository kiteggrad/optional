package optional_test

import (
	"fmt"

	"github.com/kiteggrad/optional"
)

func ExampleA_Set() {
	type Config struct {
		Name optional.A[string]
	}
	cfg := Config{}

	cfg.Name = cfg.Name.Set("AppName")
	fmt.Println(cfg.Name.Value())

	// Output: AppName
}

func ExampleA_Unset() {
	type Config struct {
		Port optional.A[int]
	}
	cfg := Config{}

	cfg.Port = cfg.Port.Set(8080)

	cfg.Port = cfg.Port.Unset()
	fmt.Println(cfg.Port.IsSet())
	fmt.Println(cfg.Port.Value())

	// Output:
	// false
	// 0
}

func ExampleA_IsSet() {
	type Config struct {
		Port optional.A[int]
	}
	cfg := Config{}

	fmt.Println(cfg.Port.IsSet())

	cfg.Port = cfg.Port.Set(8080)
	fmt.Println(cfg.Port.IsSet())

	// Output:
	// false
	// true
}

func ExampleA_MustValue() {
	type Config struct {
		Port optional.A[int]
	}
	cfg := Config{}

	// Will panic because value is not set
	recoverAndPrint(func() { cfg.Port.MustValue() })

	// Setting a value prevents panic
	cfg.Port = cfg.Port.Set(8080)
	fmt.Println(cfg.Port.MustValue())

	// Output:
	// panic: value is not set
	// 8080
}

func ExampleA_Value() {
	type Config struct {
		Ports optional.A[[]int]
	}
	cfg := Config{}

	// Check if the value is nil
	slice := cfg.Ports.Value()
	fmt.Println(slice)
	fmt.Println(slice == nil)

	cfg.Ports = cfg.Ports.Set([]int{8080})
	fmt.Println(cfg.Ports.Value())

	// Output:
	// []
	// true
	// [8080]
}

func ExampleA_SetDefault() {
	type Config struct {
		Port optional.A[int]
	}
	cfg := Config{}

	// Will be applied because Port is not set
	cfg.Port = cfg.Port.SetDefault(8080)
	fmt.Println(cfg.Port.Value())

	// Will not be applied because is set
	cfg.Port = cfg.Port.SetDefault(9090)
	fmt.Println(cfg.Port.Value())

	// Output:
	// 8080
	// 8080
}

// Example demonstrating usage of type A with a custom type that doesn't support comparison.
func ExampleA() {
	type User struct {
		ID    int
		Names []string
	}
	type Request struct {
		User optional.A[User]
	}

	req := Request{}

	// Without setting a value
	fmt.Println(req.User.IsSet())

	// Setting the user
	user := User{ID: 1, Names: []string{"John"}}
	req.User = req.User.Set(user)

	fmt.Println(req.User.IsSet())
	fmt.Println(req.User.Value())

	// Output:
	// false
	// true
	// {1 [John]}
}

func ExampleA_SetPtr() {
	type Config struct {
		Port optional.A[int]
	}
	cfg := Config{}

	// Creating a pointer to a value
	port := 8080
	portPtr := &port

	// Setting value from the pointer
	cfg.Port = cfg.Port.SetPtr(portPtr)
	fmt.Println(cfg.Port.IsSet())
	fmt.Println(cfg.Port.Value())

	// Setting nil pointer (value will be unset)
	var nilPtr *int
	cfg.Port = cfg.Port.SetPtr(nilPtr)
	fmt.Println(cfg.Port.IsSet())

	// Output:
	// true
	// 8080
	// false
}

func ExampleNewASet() {
	// Example with a custom type
	type Point struct {
		X, Y int
	}

	pointOpt := optional.NewASet(Point{X: 10, Y: 20})
	fmt.Println(pointOpt.IsSet())
	fmt.Println(pointOpt.Value())

	// Output:
	// true
	// {10 20}
}

func ExampleNewFromPtr() {
	// Example with nil pointer
	var ptr *string
	opt := optional.NewFromPtr(ptr)
	fmt.Println(opt.IsSet())

	// Example with non-nil pointer
	str := "Hello, World!"
	ptr = &str
	opt = optional.NewFromPtr(ptr)
	fmt.Println(opt.IsSet())
	fmt.Println(opt.Value())

	// Output:
	// false
	// true
	// Hello, World!
}

func ExampleNewAFromPtr() {
	// Working with arbitrary data type
	type User struct {
		ID   int
		Name string
	}

	// Creating optional value from nil pointer
	var userPtr *User
	optUser := optional.NewAFromPtr(userPtr)
	fmt.Println(optUser.IsSet())

	// Creating optional value from non-nil pointer
	user := User{ID: 1, Name: "John"}
	userPtr = &user
	optUser = optional.NewAFromPtr(userPtr)
	fmt.Println(optUser.IsSet())
	fmt.Println(optUser.Value())

	// Output:
	// false
	// true
	// {1 John}
}

func recoverAndPrint(panicsFunc func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r) //nolint:forbidigo
		}
	}()
	panicsFunc()
}
