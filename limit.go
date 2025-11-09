package iters

import (
	"iter"
)

// Limit returns a sequence that yields at most n values from seq. When n is
// zero or negative the returned sequence is empty.
func Limit[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		for item := range seq {
			if count >= n {
				return
			}
			if !yield(item) {
				return
			}
			count++
		}
	}
}

// Limit2 returns a sequence that yields at most n key/value pairs from
// seq2.
func Limit2[K, V any](seq2 iter.Seq2[K, V], n int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		count := 0
		for k, v := range seq2 {
			if count >= n {
				return
			}
			if !yield(k, v) {
				return
			}
			count++
		}
	}
}
