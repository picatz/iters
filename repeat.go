package iters

import "iter"

// Repeat returns an infinite sequence that keeps yielding value until the
// consumer stops iteration.
func Repeat[T any](value T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(value) {
				return
			}
		}
	}
}

// RepeatFunc returns an infinite sequence that yields the result of fn on
// every iteration.
func RepeatFunc[T any](fn func() T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(fn()) {
				return
			}
		}
	}
}

// RepeatN yields value count times. When count <= 0 nothing is produced.
func RepeatN[T any](value T, count int) iter.Seq[T] {
	return func(yield func(T) bool) {
		for range count {
			if !yield(value) {
				return
			}
		}
	}
}
