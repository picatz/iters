package iters

import "iter"

// Zip returns a sequence of pairs formed by reading seq1 and seq2 in lockstep.
// Iteration stops when either input sequence ends.
func Zip[T1, T2 any](seq1 iter.Seq[T1], seq2 iter.Seq[T2]) iter.Seq2[T1, T2] {
	return func(yield func(T1, T2) bool) {
		next1, stop1 := iter.Pull(seq1)
		defer stop1()

		next2, stop2 := iter.Pull(seq2)
		defer stop2()

		var (
			v1, ok1 = next1()
			v2, ok2 = next2()
		)

		for ok1 && ok2 {
			if !yield(v1, v2) {
				return
			}
			v1, ok1 = next1()
			v2, ok2 = next2()
		}
	}
}
