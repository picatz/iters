package iters

import (
	"iter"
)

// Stop returns a sequence that yields values from seq until stop reports
// true for an element. The matching element is discarded. If stop never
// returns true the entire input is forwarded.
func Stop[T comparable](seq iter.Seq[T], stop func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range seq {
			if stop(item) {
				return
			}
			if !yield(item) {
				return
			}
		}
	}
}

// Stop2 returns a sequence that yields pairs from seq2 until stop reports
// true for a key/value pair, excluding the matching pair.
func Stop2[K, V comparable](seq iter.Seq2[K, V], stop func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if stop(k, v) {
				return
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
