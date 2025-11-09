package iters

import (
	"cmp"
	"iter"
	"slices"
)

// Sort collects seq into a slice, sorts it in ascending order using the
// natural ordering for cmp.Ordered types, and returns a sequence that yields
// the sorted values.
func Sort[T cmp.Ordered](seq iter.Seq[T]) iter.Seq[T] {
	items := slices.Collect(seq)

	slices.Sort(items)

	return slices.Values(items)
}

// SortFunc behaves like Sort but orders the elements with cmp, matching
// slices.SortFunc's contract.
func SortFunc[T any](seq iter.Seq[T], cmp func(a T, b T) int) iter.Seq[T] {
	items := slices.Collect(seq)

	slices.SortFunc(items, cmp)

	return slices.Values(items)
}
