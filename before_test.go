package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleBefore_numbers() {
	numbers := []int{1, 2, 3, 4, 5}

	beforeThree := iters.Before(slices.Values(numbers), 3)

	fmt.Println(slices.Collect(beforeThree))
	// Output:
	// [1 2 3]
}

func ExampleBeforeFunc_untilEven() {
	numbers := []int{1, 3, 5, 4, 7}

	beforeEven := iters.BeforeFunc(
		slices.Values(numbers),
		func(v int) bool { return v%2 == 0 },
	)

	fmt.Println(slices.Collect(beforeEven))
	// Output:
	// [1 3 5]
}

type beforeTableTest[T comparable] struct {
	name     string
	input    []T
	n        int
	expected []T
}

func (test beforeTableTest[T]) Run(t *testing.T) {
	runBeforeTableTest(t, test)
}

func runBeforeTableTest[T comparable](t *testing.T, test beforeTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Before(slices.Values(test.input), test.n))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Before: expected %v, got %v", test.expected, got)
		}
	})
}

type beforeFuncTableTest[T comparable] struct {
	name     string
	input    []T
	pred     func(T) bool
	expected []T
}

func (test beforeFuncTableTest[T]) Run(t *testing.T) {
	runBeforeFuncTableTest(t, test)
}

func runBeforeFuncTableTest[T comparable](t *testing.T, test beforeFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.BeforeFunc(slices.Values(test.input), test.pred))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("BeforeFunc: expected %v, got %v", test.expected, got)
		}
	})
}

func TestBefore(t *testing.T) {
	tests := []runnableTest{
		beforeTableTest[int]{
			name:     "take first three",
			input:    []int{1, 2, 3, 4, 5},
			n:        3,
			expected: []int{1, 2, 3},
		},
		beforeTableTest[int]{
			name:     "n larger than len",
			input:    []int{1, 2},
			n:        5,
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

func TestBeforeFunc(t *testing.T) {
	tests := []runnableTest{
		beforeFuncTableTest[int]{
			name:  "stop before even",
			input: []int{1, 3, 5, 4, 6},
			pred:  func(v int) bool { return v%2 == 0 },
			expected: []int{
				1, 3, 5,
			},
		},
		beforeFuncTableTest[int]{
			name:     "predicate never true",
			input:    []int{1, 2},
			pred:     func(int) bool { return false },
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
