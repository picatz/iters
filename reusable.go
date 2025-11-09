package iters

import "iter"

// Reusable caches every element pulled from seq so that the returned
// sequence can be iterated multiple times. The cache grows until seq is
// exhausted, so using Reusable with an unbounded input can consume
// unbounded memory.
func Reusable[T any](seq iter.Seq[T]) iter.Seq[T] {
	var cache []T
	var done bool

	return func(yield func(T) bool) {
		for i := 0; ; i++ {
			if i < len(cache) {
				if !yield(cache[i]) {
					return
				}
				continue
			}
			if done {
				return
			}
			for item := range seq {
				cache = append(cache, item)
				if !yield(item) {
					return
				}
			}
			done = true
			return
		}
	}
}

// Reusable2 caches each key/value pair pulled from seq2 so that the result
// can be iterated multiple times. It shares the same memory trade-offs as
// Reusable.
func Reusable2[K, V any](seq2 iter.Seq2[K, V]) iter.Seq2[K, V] {
	var cacheKeys []K
	var cacheValues []V
	var done bool

	return func(yield func(K, V) bool) {
		for i := 0; ; i++ {
			if i < len(cacheKeys) {
				if !yield(cacheKeys[i], cacheValues[i]) {
					return
				}
				continue
			}
			if done {
				return
			}
			for k, v := range seq2 {
				cacheKeys = append(cacheKeys, k)
				cacheValues = append(cacheValues, v)
				if !yield(k, v) {
					return
				}
			}
			done = true
			return
		}
	}
}
