package iters

import "iter"

// Unique returns a sequence that yields each distinct element from seq once,
// preserving the first occurrence order. It keeps a set of seen values, so
// T must be comparable.
func Unique[T comparable](seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		seen := make(map[T]struct{})
		for item := range seq {
			if _, exists := seen[item]; !exists {
				seen[item] = struct{}{}
				if !yield(item) {
					return
				}
			}
		}
	}
}

// UniqueFunc behaves like Unique but determines equality with equal,
// allowing use with non-comparable types.
func UniqueFunc[T any](seq iter.Seq[T], equal func(a, b T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		var seen []T
		for item := range seq {
			isUnique := true
			for _, s := range seen {
				if equal(item, s) {
					isUnique = false
					break
				}
			}
			if isUnique {
				seen = append(seen, item)
				if !yield(item) {
					return
				}
			}
		}
	}
}
