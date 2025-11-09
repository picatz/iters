package iters

import "iter"

// Equal reports whether seq1 and seq2 yield the same elements in the same
// order. Iteration stops as soon as a mismatch is found.
func Equal[T comparable](seq1, seq2 iter.Seq[T]) bool {
	next1, stop1 := iter.Pull(seq1)
	defer stop1()

	next2, stop2 := iter.Pull(seq2)
	defer stop2()

	for {
		var (
			v1, ok1 = next1()
			v2, ok2 = next2()
		)

		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if v1 != v2 {
			return false
		}
	}
}

// EqualFunc reports whether seq1 and seq2 yield elements that are equal
// according to equal. The sequences must produce the same number of items.
func EqualFunc[T any](seq1, seq2 iter.Seq[T], equal func(T, T) bool) bool {
	next1, stop1 := iter.Pull(seq1)
	defer stop1()

	next2, stop2 := iter.Pull(seq2)
	defer stop2()

	for {
		v1, ok1 := next1()
		v2, ok2 := next2()

		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if !equal(v1, v2) {
			return false
		}
	}
}

// Equal2 reports whether two sequences of key/value pairs yield identical
// pairs in the same order.
func Equal2[K comparable, V comparable](seq1, seq2 iter.Seq2[K, V]) bool {
	next1, stop1 := iter.Pull2(seq1)
	defer stop1()

	next2, stop2 := iter.Pull2(seq2)
	defer stop2()

	for {
		var (
			k1, v1, ok1 = next1()
			k2, v2, ok2 = next2()
		)

		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if k1 != k2 || v1 != v2 {
			return false
		}
	}
}

// EqualFunc2 reports whether seq1 and seq2 produce key/value pairs that are
// equal according to equal. The sequences must yield the same number of pairs.
func EqualFunc2[K any, V any](seq1, seq2 iter.Seq2[K, V], equal func(K, V, K, V) bool) bool {
	next1, stop1 := iter.Pull2(seq1)
	defer stop1()

	next2, stop2 := iter.Pull2(seq2)
	defer stop2()

	for {
		k1, v1, ok1 := next1()
		k2, v2, ok2 := next2()

		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if !equal(k1, v1, k2, v2) {
			return false
		}
	}
}
