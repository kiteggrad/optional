package optional

// type R[T comparable] struct {
// 	value T
// 	isSet bool
// }

// // Set value and change isSet to true.
// func (o R[V]) Set(value V) R[V] {
// 	o.value = value
// 	o.isSet = true

// 	return o
// }

// // SetNotEmpty - sets the value if it is not empty and marks it as set.
// func (o R[V]) SetNotEmpty(value V) R[V] {
// 	var empty V
// 	if value == empty {
// 		return o
// 	}

// 	o.value = value
// 	o.isSet = true

// 	return o
// }

// // SetAuto - set value and change isSet to false if value is empty.
// func (o R[V]) SetAuto(value V) R[V] {
// 	var empty V

// 	if value == empty {
// 		o.isSet = false
// 	}

// 	o.value = value

// 	return o
// }

// // Unset - set value to empty and change isSet to false.
// func (o R[V]) Unset() R[V] {
// 	var empty V

// 	o.value = empty
// 	o.isSet = false

// 	return o
// }

// // IsSet - check if value is set.
// func (o R[V]) IsSet() bool {
// 	return o.isSet
// }

// // MustValue - get value and panic if value is not set.
// func (o R[V]) MustValue() V {
// 	if !o.isSet {
// 		panic(ErrNotSet)
// 	}

// 	return o.value
// }

// // Value - get value and return empty value if value is not set.
// func (o R[V]) Value() (V, error) {
// 	if !o.isSet {
// 		var empty V
// 		return empty, ErrNotSet
// 	}

// 	return o.value, nil
// }

// // IsEmpty - check if value is empty.
// func (o R[V]) IsEmpty() bool {
// 	var empty V

// 	return o.value == empty
// }

// // SetDefault - set value and isSet=true if isSet is false.
// func (o R[V]) SetDefault(value V) R[V] {
// 	if o.isSet {
// 		return o
// 	}

// 	o.value = value
// 	o.isSet = true

// 	return o
// }
