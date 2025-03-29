package iters

import "iter"

// Filter filters elements from a [sequence] based on a provided predicate
// function. It returns a new sequence that only includes elements for which
// the predicate function returns true. This is similar to the [filter]
// function found in many programming languages.
//
// [sequence]: https://pkg.go.dev/iter#Seq
// [filter]: https://en.wikipedia.org/wiki/Filter_(higher-order_function)
func Filter[V any](seq iter.Seq[V], fn func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range seq {
			if fn(item) && !yield(item) {
				return
			}
		}
	}
}

// Filter2 filters pairs of values from a [sequence] based on a provided
// predicate function that takes both the key and the value as arguments.
// It returns a new sequence that only includes value pairs for which
// the predicate function returns true. This is similar to the [filter]
// function found in many programming languages, but for value pairs.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
// [filter]: https://en.wikipedia.org/wiki/Filter_(higher-order_function)
func Filter2[K, V any](seq2 iter.Seq2[K, V], fn func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq2 {
			if fn(k, v) && !yield(k, v) {
				return
			}
		}
	}
}
