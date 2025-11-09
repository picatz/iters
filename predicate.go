package iters

// Predicate describes a boolean test applied to elements of a sequence.
type Predicate[T any] = func(T) bool

// Predicate2 is the keyed counterpart to Predicate.
type Predicate2[K, V any] = func(K, V) bool
