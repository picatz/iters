package iters

import "iter"

// Reduce applies a function to each element in the input sequence and
// reduces it to a single value. It takes an initial value and a
// function that combines the current accumulated value with each
// element in the sequence. This is similar to the [reduce] (or fold)
// function found in many programming languages.
//
// [reduce]: https://en.wikipedia.org/wiki/Reduce_(higher-order_function)
func Reduce[T, R any](seq iter.Seq[T], fn func(R, T) R, initial R) R {
	for item := range seq {
		initial = fn(initial, item)
	}
	return initial
}

// Reduce2 applies a function to pairs of values from a [sequence] based
// on a provided function and reduces it to a single value. It takes an initial
// value and a function that combines the current accumulated value with each
// value pair in the sequence. This is similar to the [reduce] (or fold)
// function found in many programming languages, but for value pairs.
//
// [reduce]: https://en.wikipedia.org/wiki/Reduce_(higher-order_function)
func Reduce2[K, V, R any](seq2 iter.Seq2[K, V], fn func(R, K, V) R, initial R) R {
	for k, v := range seq2 {
		initial = fn(initial, k, v)
	}
	return initial
}
