package iters

import "iter"

// Chunk groups seq into contiguous slices of length size and returns a new
// sequence that yields those slices in order. The final chunk may be shorter
// if the input length is not divisible by size. When size <= 0 no chunks are
// produced.
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

// ChunkFunc groups seq into slices, starting a new chunk each time pred
// returns true for an element. The matching element begins the next chunk.
// Any partial chunk is yielded when seq is exhausted.
func ChunkFunc[T any](seq iter.Seq[T], pred Predicate[T]) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		var chunk []T
		for item := range seq {
			if pred(item) && len(chunk) > 0 {
				if !yield(chunk) {
					return
				}
				chunk = nil
			}
			chunk = append(chunk, item)
		}
		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}

// Chunk2 is the keyed companion to Chunk; it groups seq2 into slices of keys
// and values with the requested size. The final chunk may be shorter. When
// size <= 0 no chunks are produced.
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

// ChunkFunc2 is the keyed companion to ChunkFunc; it starts a fresh chunk when
// pred reports true for a key/value pair. The matching pair becomes the first
// element of the next chunk.
func ChunkFunc2[K, V any](seq2 iter.Seq2[K, V], pred Predicate2[K, V]) iter.Seq2[[]K, []V] {
	return func(yield func([]K, []V) bool) {
		var keys []K
		var values []V
		for k, v := range seq2 {
			if pred(k, v) && len(keys) > 0 {
				if !yield(keys, values) {
					return
				}
				keys = nil
				values = nil
			}
			keys = append(keys, k)
			values = append(values, v)
		}
		if len(keys) > 0 {
			yield(keys, values)
		}
	}
}
