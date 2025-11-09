package iters

import "iter"

// Before returns a sequence that yields at most n elements from seq.
// If seq is shorter than n the entire input is passed through.
func Before[T any](seq iter.Seq[T], n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		count := 0
		for item := range seq {
			if count >= n {
				return
			}
			if !yield(item) {
				return
			}
			count++
		}
	}
}

// BeforeFunc returns a sequence that yields elements from seq until pred
// first returns true. The matching element is not forwarded. If pred never
// returns true the entire input is yielded.
func BeforeFunc[T any](seq iter.Seq[T], pred Predicate[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range seq {
			if pred(item) {
				return
			}
			if !yield(item) {
				return
			}
		}
	}
}
