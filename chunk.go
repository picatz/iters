package iters

import "iter"

// Chunk splits a [sequence] into another sequence of chunks, where each chunk
// is a slice of elements with the specified size, except possibly the last chunk
// which may contain fewer elements if there are not enough remaining elements
// in the original sequence. If the specified size is less than or equal to zero,
// an empty sequence is returned.
//
// [sequence]: https://pkg.go.dev/iter#Seq
func Chunk[T any](seq iter.Seq[T], size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if size <= 0 {
			return
		}
		chunk := make([]T, 0, size)
		for item := range seq {
			chunk = append(chunk, item)
			if len(chunk) == size {
				if !yield(chunk) {
					return
				}
				chunk = make([]T, 0, size)
			}
		}
		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}

// Chunk2 splits a [sequence] of pairs of values into another sequence of chunks,
// where each chunk is a pair of slices: one slice containing keys and the other
// containing values, both with the specified size, except possibly the last chunk
// which may contain fewer elements if there are not enough remaining elements
// in the original sequence. If the specified size is less than or equal to zero,
// an empty sequence is returned.
//
// [sequence]: https://pkg.go.dev/iter#Seq2
func Chunk2[K, V any](seq2 iter.Seq2[K, V], size int) iter.Seq2[[]K, []V] {
	return func(yield func([]K, []V) bool) {
		if size <= 0 {
			return
		}
		keys := make([]K, 0, size)
		values := make([]V, 0, size)
		for k, v := range seq2 {
			keys = append(keys, k)
			values = append(values, v)
			if len(keys) == size {
				if !yield(keys, values) {
					return
				}
				keys = make([]K, 0, size)
				values = make([]V, 0, size)
			}
		}
		if len(keys) > 0 {
			yield(keys, values)
		}
	}
}
