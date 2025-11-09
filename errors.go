package iters

import "iter"

// WalkErr iterates over seq, invoking fn for each value until either fn returns
// false or seq yields a non-nil error. The first error encountered is returned.
func WalkErr[T any](seq iter.Seq2[T, error], fn func(T) bool) error {
	next, stop := iter.Pull2(seq)
	defer stop()

	for {
		v, err, ok := next()
		if !ok {
			return nil
		}
		if err != nil {
			return err
		}
		if !fn(v) {
			return nil
		}
	}
}

// UntilErr converts seq into a plain sequence that yields values until a
// non-nil error appears. The error-causing element is discarded.
func UntilErr[T any](seq iter.Seq2[T, error]) iter.Seq[T] {
	return func(yield func(T) bool) {
		next, stop := iter.Pull2(seq)
		defer stop()

		for {
			v, err, ok := next()
			if !ok || err != nil {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

// CollectErr gathers values from seq until a non-nil error occurs. It returns
// the values seen so far along with the first error, or nil if seq completed.
func CollectErr[T any](seq iter.Seq2[T, error]) ([]T, error) {
	next, stop := iter.Pull2(seq)
	defer stop()

	var out []T
	for {
		v, err, ok := next()
		if !ok {
			return out, nil
		}
		if err != nil {
			return out, err
		}
		out = append(out, v)
	}
}
