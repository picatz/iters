package iters

import (
	"cmp"
	"iter"
)

// Max returns the largest element produced by seq according to Go's
// ordering for cmp.Ordered types. If seq is empty, ok is false.
func Max[T cmp.Ordered](seq iter.Seq[T]) (max T, ok bool) {
	first := true
	for item := range seq {
		if first {
			max = item
			first = false
			ok = true
			continue
		}
		if item > max {
			max = item
		}
	}
	return
}

// MaxFunc returns the element in seq that maximizes the provided less
// function (a returns true when a < b). If seq is empty, ok is false.
func MaxFunc[T any](seq iter.Seq[T], less func(a, b T) bool) (max T, ok bool) {
	first := true
	for item := range seq {
		if first {
			max = item
			first = false
			ok = true
			continue
		}
		if less(max, item) {
			max = item
		}
	}
	return
}
