// Package optional provides a generic implementation for optional values in Go.
// This file contains implementation for any type (not just comparable).
package optional

// A - structure for optional values that can hold any type.
// Unlike T, this structure works with any Go type, not just comparable ones.
// It cannot check if a value is empty, but still allows checking if a value is set.
//
// A is most useful for non-comparable structs (structs containing maps, slices, etc.).
// For reference types (slices, maps, channels), which can already be checked for nil,
// consider using them directly rather than wrapping in optional.A, as direct nil checks
// are more idiomatic in Go. Optional.A adds the most value when working with custom
// non-comparable structs where nil checks aren't applicable.
//
// Example:
//
//	var opt optional.A[CustomStruct]
type A[V any] struct {
	value V
	isSet bool
}

// NewAFromPtr - create new optional value from pointer.
// If pointer is nil, it returns value with isSet=false.
// Otherwise, it returns value with isSet=true.
func NewAFromPtr[V any](ptr *V) A[V] {
	if ptr == nil {
		return A[V]{}
	}

	return A[V]{value: *ptr, isSet: true}
}

// NewASet - create new optional value with setted value.
// It returns isSet=true always.
func NewASet[V any](value V) A[V] {
	return A[V]{value: value, isSet: true}
}

// Set value and change isSet to true.
func (o A[V]) Set(value V) A[V] {
	o.value = value
	o.isSet = true

	return o
}

// SetPtr - set value from pointer.
// If pointer is nil, it sets isSet=false.
func (o A[V]) SetPtr(ptr *V) A[V] {
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

// Unset - set value to zero value and change isSet to false.
func (o A[V]) Unset() A[V] {
	var empty V

	o.value = empty
	o.isSet = false

	return o
}

// IsSet - check if value is set.
func (o A[V]) IsSet() bool {
	return o.isSet
}

// MustValue - get value or panics with ErrNotSet if value is not set.
func (o A[V]) MustValue() V {
	if !o.isSet {
		panic(ErrNotSet)
	}

	return o.value
}

// Value - get value and return zero value if value is not set.
func (o A[V]) Value() V {
	return o.value
}

// SetDefault - set value and isSet=true if isSet is false.
func (o A[V]) SetDefault(value V) A[V] {
	if o.isSet {
		return o
	}

	o.value = value
	o.isSet = true

	return o
}
