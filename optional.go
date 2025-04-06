// Package optional provides a generic implementation for optional values in Go.
// It allows working with optional fields without using pointers, ensuring type safety
// and simplifying checks for empty or unset values.
package optional

// T - structure for optional values.
// It is used to avoid using pointers for optional values.
// We use comparable type for having possibility to easily check if value is empty.
//
// Do not try to use it with pointer types, the methods will not work as expected.
// Pointer types are already optional by default - just use nil check.
//
// Example:
//
//	var opt optional.T[string]
type T[V comparable] struct {
	value V
	isSet bool
}

// Set value and change isSet to true.
func (o T[V]) Set(value V) T[V] {
	o.value = value
	o.isSet = true

	return o
}

// SetNotEmpty - sets the value if it is not empty and marks it as set.
func (o T[V]) SetNotEmpty(value V) T[V] {
	var empty V
	if value == empty {
		return o
	}

	o.value = value
	o.isSet = true

	return o
}

// SetAuto - set value and change isSet to false if value is empty.
func (o T[V]) SetAuto(value V) T[V] {
	var empty V

	if value == empty {
		o.isSet = false
	}

	o.value = value

	return o
}

// Unset - set value to empty and change isSet to false.
func (o T[V]) Unset() T[V] {
	var empty V

	o.value = empty
	o.isSet = false

	return o
}

// IsSet - check if value is set.
func (o T[V]) IsSet() bool {
	return o.isSet
}

// MustValue - get value and panic if value is not set.
func (o T[V]) MustValue() V {
	if !o.isSet {
		panic("value is not set")
	}

	return o.value
}

// Value - get value and return empty value if value is not set.
func (o T[V]) Value() V {
	return o.value
}

// IsEmpty - check if value is empty.
func (o T[V]) IsEmpty() bool {
	var empty V

	return o.value == empty
}

// SetDefault - set value and isSet=true if isSet is false.
func (o T[V]) SetDefault(value V) T[V] {
	if o.isSet {
		return o
	}

	o.value = value
	o.isSet = true

	return o
}
