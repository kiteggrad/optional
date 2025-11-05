// Package optional provides a generic implementation for optional values in Go.
// It allows working with optional fields without using pointers.
package optional

import "errors"

// ErrNotSet - could be returned in some cases when value is not set.
var ErrNotSet = errors.New("value is not set")

// T - structure for optional values.
// It is used to avoid using pointers for optional values.
//
// Example:
//
//	var opt optional.T[string]
type T[V any] struct {
	value V
	isSet bool
}

// New - create new optional value.
//   - isSet - indicates whether the value is set or not.
//   - If isSet is false, value will be the zero value of type V
//     regardless of the provided value.
//
// Example:
//
//	opt := optional.New("value", true)  // opt is set to "value"
//	opt2 := optional.New("value", false) // opt2 is not set, value is ""
//	opt3 := optional.New(slice, slice != nil) // opt3 is set if slice is not nil
//	s.optField = optional.New(someValue, someCondition) // instead of unexisting Set() method
func New[V any](value V, isSet bool) T[V] {
	if !isSet {
		var empty V
		value = empty
	}

	return T[V]{value: value, isSet: isSet}
}

// NewPtr - create new optional value from pointer.
//   - If ptr is nil, the optional value will be not set.
//   - If ptr is not nil, the optional value will be set to the value pointed by ptr.
//   - Doesn't work with reference types, e.g. slices, maps, etc. You can use New(slice, slice != nil) instead.
//
// Example:
//
//	var strPtr *string
//	opt1 := optional.NewPtr(strPtr) // opt1 is not set
//
//	value := "hello"
//	strPtr = &value
//	opt2 := optional.NewPtr(strPtr) // opt2 is set to "hello"
func NewPtr[V any](ptr *V) T[V] {
	if ptr == nil {
		return T[V]{}
	}

	return New(*ptr, true)
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

// Ptr - get pointer to value or nil if value is not set.
func (o T[V]) Ptr() *V {
	if !o.isSet {
		return nil
	}

	return &o.value
}
