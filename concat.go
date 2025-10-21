package iters

import "iter"

// Concat concatenates multiple sequences into a single sequence. It takes
// a variable number of sequences as input and returns a new sequence that
// yields elements from each input sequence in order.
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

// Concat2 concatenates multiple sequences of pairs of values, into a single
// sequence of pairs of values. It takes a variable number of sequences as input and returns
// a new sequence that yields pairs of values from each input sequence in order.
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
