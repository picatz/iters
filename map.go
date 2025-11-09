package iters

import "iter"

// Map returns a sequence whose elements are fn(item) for each element of
// seq, in order.
func Map[T, R any](seq iter.Seq[T], fn Mapper[T, R]) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range seq {
			if !yield(fn(item)) {
				return
			}
		}
	}
}

// Map2 applies fn to every key/value pair from seq2 and yields the
// resulting pairs.
func Map2[K1, V1, K2, V2 any](seq2 iter.Seq2[K1, V1], fn Mapper2[K1, V1, K2, V2]) iter.Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		for k, v := range seq2 {
			if !yield(fn(k, v)) {
				return
			}
		}
	}
}
