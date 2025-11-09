package iters

import "iter"

// First returns the first element produced by seq. If seq yields nothing,
// ok is false and the zero value is returned.
func First[T any](seq iter.Seq[T]) (first T, ok bool) {
	for item := range seq {
		first = item
		ok = true
		return
	}
	return
}

// FirstFunc returns the first element of seq that satisfies pred. If no
// element matches, ok is false.
func FirstFunc[T any](seq iter.Seq[T], pred Predicate[T]) (first T, ok bool) {
	for item := range seq {
		if pred(item) {
			first = item
			ok = true
			return
		}
	}
	return
}

// First2 returns the first key/value pair from seq2. If seq2 yields no
// pairs, ok is false.
func First2[K, V any](seq2 iter.Seq2[K, V]) (firstK K, firstV V, ok bool) {
	for k, v := range seq2 {
		firstK = k
		firstV = v
		ok = true
		return
	}
	return
}

// FirstFunc2 returns the first pair from seq2 that satisfies pred. If no
// pair matches, ok is false.
func FirstFunc2[K, V any](seq2 iter.Seq2[K, V], pred Predicate2[K, V]) (firstK K, firstV V, ok bool) {
	for k, v := range seq2 {
		if pred(k, v) {
			firstK = k
			firstV = v
			ok = true
			return
		}
	}
	return
}
