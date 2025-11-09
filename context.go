package iters

import (
	"context"
	"iter"
)

// Context returns a sequence that yields values from seq until ctx is
// canceled or seq finishes. Cancellation is checked between elements.
func Context[T any](ctx context.Context, seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range seq {
			if ctx.Err() != nil {
				return
			}
			if !yield(item) {
				return
			}
		}
	}
}

// Context2 returns a sequence of key/value pairs from seq2 that stops
// yielding as soon as ctx is canceled.
func Context2[K, V any](ctx context.Context, seq2 iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq2 {
			if ctx.Err() != nil {
				return
			}
			if !yield(k, v) {
				return
			}
		}
	}
}
