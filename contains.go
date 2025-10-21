package iters

import "iter"

// Contains checks if a given value exists within a [sequence].
func Contains[V comparable](seq iter.Seq[V], value V) bool {
	for item := range seq {
		if item == value {
			return true
		}
	}
	return false
}

// Contains2 checks if a given pair of values exists within the [iter.Seq2].
func Contains2[K comparable, V comparable](seq2 iter.Seq2[K, V], key K, value V) bool {
	for k, v := range seq2 {
		if k == key && v == value {
			return true
		}
	}
	return false
}

// ContainsFunc checks if any element in the [sequence] satisfies the provided
// predicate function.
func ContainsFunc[V any](seq iter.Seq[V], fn func(V) bool) bool {
	for item := range seq {
		if fn(item) {
			return true
		}
	}
	return false
}

// ContainsFunc2 checks if any key-value pair in the [iter.Seq2] satisfies
// the provided predicate function.
func ContainsFunc2[K, V any](seq2 iter.Seq2[K, V], fn func(K, V) bool) bool {
	for k, v := range seq2 {
		if fn(k, v) {
			return true
		}
	}
	return false
}
