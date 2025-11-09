package iters

import "iter"

// Number matches any built-in integer or floating-point type.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

// Average returns the arithmetic mean of seq. It returns 0 when seq
// contains no values.
func Average[T Number](seq iter.Seq[T]) float64 {
	var (
		sum   float64
		count int
	)

	for item := range seq {
		sum += float64(item)
		count++
	}

	if count == 0 {
		return 0
	}
	return sum / float64(count)
}

// AverageFunc computes the arithmetic mean of fn(item) for every element
// in seq. It returns 0 when seq is empty.
func AverageFunc[T any](seq iter.Seq[T], fn func(T) float64) float64 {
	var (
		sum   float64
		count int
	)

	for item := range seq {
		sum += fn(item)
		count++
	}

	if count == 0 {
		return 0
	}
	return sum / float64(count)
}
