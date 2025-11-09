package iters

import "iter"

// Filter returns a sequence that yields only the elements of seq that make
// fn return true.
func Filter[V any](seq iter.Seq[V], fn Predicate[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range seq {
			if fn(item) && !yield(item) {
				return
			}
		}
	}
}

// Filter2 is the keyed companion to Filter; it forwards only the pairs from
// seq2 that satisfy fn.
func Filter2[K, V any](seq2 iter.Seq2[K, V], fn Predicate2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq2 {
			if fn(k, v) && !yield(k, v) {
				return
			}
		}
	}
}
