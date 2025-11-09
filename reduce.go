package iters

import "iter"

// Reduce folds seq into a single value by repeatedly applying fn to the
// accumulator and the next element. The initial value is returned unchanged
// if seq is empty.
func Reduce[T, R any](seq iter.Seq[T], fn func(R, T) R, initial R) R {
	for item := range seq {
		initial = fn(initial, item)
	}
	return initial
}

// Reduce2 folds seq2 into a single value by repeatedly calling fn with the
// accumulator and the next key/value pair.
func Reduce2[K, V, R any](seq2 iter.Seq2[K, V], fn func(R, K, V) R, initial R) R {
	for k, v := range seq2 {
		initial = fn(initial, k, v)
	}
	return initial
}
