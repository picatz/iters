package iters

import "iter"

// After returns a sequence that discards the first n elements from seq
// before yielding the remainder. If n is zero or negative the original
// sequence is yielded unchanged. If n is larger than the number of
// elements the resulting sequence is empty.
func After[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		skipped := 0
		for item := range seq {
			if skipped < n {
				skipped++
				continue
			}
			if !yield(item) {
				return
			}
		}
	}
}

// AfterFunc returns a sequence that drops elements from seq while pred
// reports true. The first element for which pred returns false and all
// subsequent elements are yielded. If pred never returns false, nothing
// is produced.
func AfterFunc[T any](seq iter.Seq[T], pred Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		skipping := true
		for item := range seq {
			if skipping {
				if pred(item) {
					continue
				}
				skipping = false
			}
			if !yield(item) {
				return
			}
		}
	}
}
