package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleAfter_numbers() {
	numbers := []int{1, 2, 3, 4, 5}

	afterTwo := iters.After(slices.Values(numbers), 2)

	fmt.Println(slices.Collect(afterTwo))
	// Output:
	// [3 4 5]
}

func ExampleAfterFunc_threshold() {
	numbers := []int{1, 3, 2, 4, 5}

	afterSmall := iters.AfterFunc(
		slices.Values(numbers),
		func(v int) bool { return v < 3 },
	)

	fmt.Println(slices.Collect(afterSmall))
	// Output:
	// [3 2 4 5]
}

type afterTableTest[T comparable] struct {
	name     string
	input    []T
	n        int
	expected []T
}

func (test afterTableTest[T]) Run(t *testing.T) {
	runAfterTableTest(t, test)
}

func runAfterTableTest[T comparable](t *testing.T, test afterTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.After(slices.Values(test.input), test.n))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("After: expected %v, got %v", test.expected, got)
		}
	})
}

type afterFuncTableTest[T comparable] struct {
	name     string
	input    []T
	pred     func(T) bool
	expected []T
}

func (test afterFuncTableTest[T]) Run(t *testing.T) {
	runAfterFuncTableTest(t, test)
}

func runAfterFuncTableTest[T comparable](t *testing.T, test afterFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.AfterFunc(slices.Values(test.input), test.pred))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("AfterFunc: expected %v, got %v", test.expected, got)
		}
	})
}

func TestAfter(t *testing.T) {
	tests := []runnableTest{
		afterTableTest[int]{
			name:     "skip first two",
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{3, 4, 5},
		},
		afterTableTest[int]{
			name:     "skip beyond length",
			input:    []int{1, 2},
			n:        5,
			expected: nil,
		},
		afterTableTest[int]{
			name:     "non-positive skip",
			input:    []int{1, 2},
			n:        0,
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

func TestAfterFunc(t *testing.T) {
	tests := []runnableTest{
		afterFuncTableTest[int]{
			name:  "skip negatives",
			input: []int{-2, -1, 0, 1, 2},
			pred:  func(v int) bool { return v < 0 },
			expected: []int{
				0, 1, 2,
			},
		},
		afterFuncTableTest[int]{
			name:     "predicate always true",
			input:    []int{1, 2},
			pred:     func(int) bool { return true },
			expected: nil,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
