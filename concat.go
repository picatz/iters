package iters

import "iter"

// Concat returns a sequence that yields every element from seqs in order.
// Iteration stops early if the consumer stops pulling values.
func Concat[T any](seqs ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, seq := range seqs {
			for item := range seq {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Concat2 is the keyed companion to Concat; it streams every pair from seqs2
// in order.
func Concat2[K, V any](seqs2 ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq2 := range seqs2 {
			for k, v := range seq2 {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
