package iters

// Mapper describes the transformation applied by Map.
type Mapper[T, R any] = func(T) R

// Mapper2 describes the transformation applied by Map2.
type Mapper2[K1, V1, K2, V2 any] = func(K1, V1) (K2, V2)
