package optional_test

import (
	"fmt"

	"github.com/kiteggrad/optional"
)

func ExampleT_Set() {
	type Config struct {
		Name optional.T[string]
	}
	cfg := Config{}

	cfg.Name = cfg.Name.Set("AppName")
	fmt.Println(cfg.Name.Value())

	// Output: AppName
}

func ExampleT_SetNotEmpty() {
	type Config struct {
		Name optional.T[string]
	}
	cfg := Config{}

	cfg.Name = cfg.Name.SetNotEmpty("")
	fmt.Println(cfg.Name.IsSet())

	cfg.Name = cfg.Name.SetNotEmpty("AppName")
	fmt.Println(cfg.Name.Value())

	cfg.Name = cfg.Name.SetNotEmpty("")
	fmt.Println(cfg.Name.IsSet())

	// Output:
	// false
	// AppName
	// true
}

func ExampleT_SetAuto() {
	type Config struct {
		Port optional.T[string]
	}
	cfg := Config{}

	cfg.Port = cfg.Port.SetAuto("")
	fmt.Println(cfg.Port.IsSet())

	cfg.Port = cfg.Port.SetAuto("8080")
	fmt.Println(cfg.Port.Value())

	cfg.Port = cfg.Port.SetAuto("")
	fmt.Println(cfg.Port.IsSet())

	// Output:
	// false
	// 8080
	// false
}

func ExampleT_Unset() {
	type Config struct {
		Port optional.T[int]
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

func ExampleT_IsSet() {
	type Config struct {
		Port optional.T[int]
	}
	cfg := Config{}

	fmt.Println(cfg.Port.IsSet())

	cfg.Port = cfg.Port.Set(8080)
	fmt.Println(cfg.Port.IsSet())

	// Output:
	// false
	// true
}

func ExampleT_MustValue() {
	type Config struct {
		Port optional.T[int]
	}
	cfg := Config{}

	cfg.Port = cfg.Port.Set(8080)
	fmt.Println(cfg.Port.MustValue())

	// Output: 8080
}

func ExampleT_Value() {
	type Config struct {
		Port optional.T[int]
	}
	cfg := Config{}

	fmt.Println(cfg.Port.Value())

	cfg.Port = cfg.Port.Set(8080)
	fmt.Println(cfg.Port.Value())

	// Output:
	// 0
	// 8080
}

func ExampleT_IsEmpty() {
	type Config struct {
		Name optional.T[string]
	}
	cfg := Config{}

	// returns true if value is empty even if is not set
	fmt.Println(cfg.Name.IsEmpty())

	cfg.Name = cfg.Name.Set("AppName")
	fmt.Println(cfg.Name.IsEmpty())

	cfg.Name = cfg.Name.Set("")
	fmt.Println(cfg.Name.IsEmpty())

	// Output:
	// true
	// false
	// true
}

func ExampleT_SetDefault() {
	type Config struct {
		Port optional.T[int]
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

func ExampleT_SetDefault_in_struct() {
	type Config struct {
		Port optional.T[int]
	}
	passedCfg := Config{}

	cfgWithDefaults := Config{
		Port: passedCfg.Port.SetDefault(8080),
	}

	fmt.Println(cfgWithDefaults.Port.Value())

	// Output:
	// 8080
}

func ExampleNewSet() {
	opt := optional.NewSet("Hello, World!")
	fmt.Println(opt.IsSet())
	fmt.Println(opt.Value())

	// Output:
	// true
	// Hello, World!
}

func ExampleNewSet_empty() {
	opt := optional.NewSet("")
	fmt.Println(opt.IsSet())
	fmt.Println(opt.Value())

	// Output:
	// true
}

func ExampleNewSetNotEmpty() {
	opt := optional.NewSetNotEmpty("")
	fmt.Println(opt.IsSet())

	opt = optional.NewSetNotEmpty("Hello, World!")
	fmt.Println(opt.IsSet())
	fmt.Println(opt.Value())

	// Output:
	// false
	// true
	// Hello, World!
}
