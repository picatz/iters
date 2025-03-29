package iters

import "iter"

// Map applies a function to each element in the input sequence and
// returns a new sequence with the results. This is similar to the
// [map] function found in many programming languages.
//
// [map]: https://en.wikipedia.org/wiki/Map_(higher-order_function)
func Map[T, R any](seq iter.Seq[T], fn func(T) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range seq {
			if !yield(fn(item)) {
				return
			}
		}
	}
}

// Map2 applies pairs of values from a [sequence] based on a provided
// and returns a new sequence with the results. This is similar to the
// [map] function but for key-value pairs. It allows you to transform
// both the key and the value of each entry in a map-like structure.
// This is useful for transforming data structures that have both
// keys and values, such as [maps].
//
// [map]: https://en.wikipedia.org/wiki/Map_(higher-order_function)
// [maps]: https://golang.org/pkg/maps/
func Map2[K1, V1, K2, V2 any](seq2 iter.Seq2[K1, V1], fn func(K1, V1) (K2, V2)) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		for k, v := range seq2 {
			if !yield(fn(k, v)) {
				return
			}
		}
	}
}
