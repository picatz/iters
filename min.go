package iters

import (
	"cmp"
	"iter"
)

// Min returns the smallest element produced by seq according to Go's
// ordering for cmp.Ordered types. If seq is empty, ok is false.
func Min[T cmp.Ordered](seq iter.Seq[T]) (min T, ok bool) {
	first := true
	for item := range seq {
		if first {
			min = item
			first = false
			ok = true
			continue
		}
		if item < min {
			min = item
		}
	}
	return
}

// MinFunc returns the element in seq that minimizes the provided less
// function. If seq is empty, ok is false.
func MinFunc[T any](seq iter.Seq[T], less func(a, b T) bool) (min T, ok bool) {
	first := true
	for item := range seq {
		if first {
			min = item
			first = false
			ok = true
			continue
		}
		if less(item, min) {
			min = item
		}
	}
	return
}
