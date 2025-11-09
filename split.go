package iters

import (
	"context"
	"iter"
)

// Split duplicates seq2 into two coordinated sequences, one exposing the
// keys and the other the values. Reads stay synchronized so that each key
// is paired with its corresponding value even when the consumers progress
// independently. Iteration stops when ctx is canceled or seq2 finishes.
func Split[K, V any](ctx context.Context, seq2 iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V]) {
	next, stop := iter.Pull2(seq2)

	// Split the keys and values into separate sequences, but only pull from the original
	// sequence once, maintaining synchronization between keys and values, but allowing
	// independent consumption of keys and values, not blocking each other.
	var (
		keyCh   = make(chan K)
		valueCh = make(chan V)
	)
	go func() {
		defer close(keyCh)
		defer close(valueCh)
		defer stop()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				k, v, ok := next()
				if !ok {
					return
				}

				kRead := false
				vRead := false

				for !kRead || !vRead {
					switch {
					case !kRead && !vRead:
						select {
						case <-ctx.Done():
							return
						case keyCh <- k:
							kRead = true
						case valueCh <- v:
							vRead = true
						}
					case !kRead:
						select {
						case <-ctx.Done():
							return
						case keyCh <- k:
							kRead = true
						}
					case !vRead:
						select {
						case <-ctx.Done():
							return
						case valueCh <- v:
							vRead = true
						}
					}
				}
			}
		}
	}()

	return func(yield func(K) bool) {
			for {
				select {
				case <-ctx.Done():
					return
				case k, ok := <-keyCh:
					if !ok {
						return
					}
					if !yield(k) {
						return
					}
				}
			}
		}, func(yield func(V) bool) {
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-valueCh:
					if !ok {
						return
					}
					if !yield(v) {
						return
					}
				}
			}
		}
}
