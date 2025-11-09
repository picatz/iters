package iters

import (
	"cmp"
	"iter"
)

// Compare returns -1, 0, or +1 depending on the lexicographic ordering of s1
// and s2, matching slices.Compare semantics. Elements are compared with
// [cmp.Compare] until a difference is found or one sequence ends.
func Compare[T cmp.Ordered](s1, s2 iter.Seq[T]) int {
	it1, stop1 := iter.Pull(s1)
	it2, stop2 := iter.Pull(s2)
	defer stop1()
	defer stop2()

	for {
		var (
			v1, ok1 = it1()
			v2, ok2 = it2()
		)

		if !ok1 && !ok2 {
			return 0 // both sequences are exhausted and equal
		}
		if !ok1 {
			return -1 // s1 is exhausted but s2 is not, so s1 < s2
		}
		if !ok2 {
			return 1 // s2 is exhausted but s1 is not, so s1 > s2
		}

		if cmpResult := cmp.Compare(v1, v2); cmpResult != 0 {
			return cmpResult
		}
		// if v1 == v2, continue to the next elements
	}
}

// CompareFunc behaves like Compare but calls cmp for each pair of elements.
// The first non-zero result is returned, otherwise the shorter sequence sorts
// before the longer one.
func CompareFunc[T1, T2 any](s1 iter.Seq[T1], s2 iter.Seq[T2], cmp func(T1, T2) int) int {
	it1, stop1 := iter.Pull(s1)
	it2, stop2 := iter.Pull(s2)
	defer stop1()
	defer stop2()

	for {
		var (
			v1, ok1 = it1()
			v2, ok2 = it2()
		)

		if !ok1 && !ok2 {
			return 0 // both sequences are exhausted and equal
		}
		if !ok1 {
			return -1 // s1 is exhausted but s2 is not, so s1 < s2
		}
		if !ok2 {
			return 1 // s2 is exhausted but s1 is not, so s1 > s2
		}

		if cmpResult := cmp(v1, v2); cmpResult != 0 {
			return cmpResult
		}
	}
}
