package iters

import "iter"

// Last returns the final element produced by seq. If seq yields nothing,
// ok is false.
func Last[T any](seq iter.Seq[T]) (last T, ok bool) {
	for item := range seq {
		last = item
		ok = true
	}
	return
}

// LastFunc returns the final element of seq that satisfies pred. If no
// element matches, ok is false.
func LastFunc[T any](seq iter.Seq[T], pred Predicate[T]) (last T, ok bool) {
	for item := range seq {
		if pred(item) {
			last = item
			ok = true
		}
	}
	return
}

// Last2 returns the final key/value pair produced by seq2. If seq2 yields
// nothing, ok is false.
func Last2[K, V any](seq2 iter.Seq2[K, V]) (lastK K, lastV V, ok bool) {
	for k, v := range seq2 {
		lastK = k
		lastV = v
		ok = true
	}
	return
}

// LastFunc2 returns the final key/value pair from seq2 that satisfies pred.
// If no pair matches, ok is false.
func LastFunc2[K, V any](seq2 iter.Seq2[K, V], pred Predicate2[K, V]) (lastK K, lastV V, ok bool) {
	for k, v := range seq2 {
		if pred(k, v) {
			lastK = k
			lastV = v
			ok = true
		}
	}
	return
}
