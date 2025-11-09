package iters

import "iter"

// Contains reports whether value appears in seq.
func Contains[V comparable](seq iter.Seq[V], value V) bool {
	for item := range seq {
		if item == value {
			return true
		}
	}
	return false
}

// Contains2 reports whether the pair (key, value) appears in seq2.
func Contains2[K comparable, V comparable](seq2 iter.Seq2[K, V], key K, value V) bool {
	for k, v := range seq2 {
		if k == key && v == value {
			return true
		}
	}
	return false
}

// ContainsFunc reports whether any element in seq satisfies fn.
func ContainsFunc[V any](seq iter.Seq[V], fn Predicate[V]) bool {
	for item := range seq {
		if fn(item) {
			return true
		}
	}
	return false
}

// ContainsFunc2 reports whether any key/value pair in seq2 satisfies fn.
func ContainsFunc2[K, V any](seq2 iter.Seq2[K, V], fn Predicate2[K, V]) bool {
	for k, v := range seq2 {
		if fn(k, v) {
			return true
		}
	}
	return false
}
