package iters

import "iter"

// Compact collapses consecutive duplicate elements in seq, yielding the first
// element of each run.
func Compact[T comparable](seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var (
			first = true
			prev  T
		)
		for item := range seq {
			if first || item != prev {
				if !yield(item) {
					return
				}
				prev = item
				first = false
			}
		}
	}
}

// CompactFunc collapses consecutive elements in seq for which equal reports
// true, yielding only the first element of each run.
func CompactFunc[T any](seq iter.Seq[T], equal func(a, b T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		var (
			first = true
			prev  T
		)
		for item := range seq {
			if first || !equal(item, prev) {
				if !yield(item) {
					return
				}
				prev = item
				first = false
			}
		}
	}
}

// Compact2 collapses consecutive duplicate key/value pairs in seq2, yielding
// the first pair from each run.
func Compact2[K comparable, V comparable](seq2 iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var (
			first = true
			prevK K
			prevV V
		)
		for k, v := range seq2 {
			if first || k != prevK || v != prevV {
				if !yield(k, v) {
					return
				}
				prevK = k
				prevV = v
				first = false
			}
		}
	}
}

// CompactFunc2 collapses consecutive key/value pairs in seq2 for which equal
// reports true, yielding the first pair of each run.
func CompactFunc2[K, V any](seq2 iter.Seq2[K, V], equal func(aK K, aV V, bK K, bV V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var (
			first = true
			prevK K
			prevV V
		)
		for k, v := range seq2 {
			if first || !equal(k, v, prevK, prevV) {
				if !yield(k, v) {
					return
				}
				prevK = k
				prevV = v
				first = false
			}
		}
	}
}
