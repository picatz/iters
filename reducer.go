package iters

// Reducer describes the accumulator function accepted by Reduce.
type Reducer[R, T any] = func(R, T) R

// Reducer2 is the keyed counterpart to Reducer.
type Reducer2[R, K, V any] = func(R, K, V) R
