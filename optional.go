// Package optional provides a generic implementation for optional values in Go.
// It allows working with optional fields without using pointers, ensuring type safety
// and simplifying checks for empty or unset values.
package optional

import "errors"

// ErrNotSet - could be returned in some cases when value is not set.
var ErrNotSet = errors.New("value is not set")

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

// NewFromPtr - create new optional value from pointer.
// If pointer is nil, it returns value with isSet=false.
// Otherwise, it returns value with isSet=true.
func NewFromPtr[V comparable](ptr *V) T[V] {
	if ptr == nil {
		return T[V]{}
	}

	return T[V]{value: *ptr, isSet: true}
}

// NewSetNotEmpty - create new optional value with default value.
//   - if value is empty, returns with isSet=false.
func NewSetNotEmpty[V comparable](value V) T[V] {
	o := T[V]{}

	return o.SetNotEmpty(value)
}

// NewSet - create new optional value with setted value.
//   - Unlike NewDefault, this function does not check if value is empty.
//     It returns isSet=true always.
func NewSet[V comparable](value V) T[V] {
	return T[V]{value: value, isSet: true}
}

// Set value and change isSet to true.
func (o T[V]) Set(value V) T[V] {
	o.value = value
	o.isSet = true

	return o
}

// SetPtr - set value from pointer.
// If pointer is nil, it sets isSet=false.
func (o T[V]) SetPtr(ptr *V) T[V] {
	if ptr == nil {
		var empty V
		o.isSet = false
		o.value = empty
	} else {
		o.value = *ptr
		o.isSet = true
	}

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

// MustValue - get value or panics with ErrNotSet if value is not set.
func (o T[V]) MustValue() V {
	if !o.isSet {
		panic(ErrNotSet)
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
