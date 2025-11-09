package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleRepeat_value() {
	repeated := iters.Limit(iters.Repeat(7), 3)
	fmt.Println(slices.Collect(repeated))
	// Output:
	// [7 7 7]
}

func ExampleRepeatFunc_counter() {
	counter := 0
	repeated := iters.Limit(
		iters.RepeatFunc(func() int {
			counter++
			return counter
		}),
		3,
	)
	fmt.Println(slices.Collect(repeated))
	// Output:
	// [1 2 3]
}

func ExampleRepeatN_value() {
	fmt.Println(slices.Collect(iters.RepeatN("go", 2)))
	// Output:
	// [go go]
}

type repeatTableTest[T comparable] struct {
	name     string
	value    T
	count    int
	expected []T
}

func (test repeatTableTest[T]) Run(t *testing.T) {
	runRepeatTableTest(t, test)
}

func runRepeatTableTest[T comparable](t *testing.T, test repeatTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Limit(iters.Repeat(test.value), test.count))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Repeat: expected %v, got %v", test.expected, got)
		}
	})
}

type repeatFuncTableTest[T comparable] struct {
	name     string
	fn       func() T
	count    int
	expected []T
}

func (test repeatFuncTableTest[T]) Run(t *testing.T) {
	runRepeatFuncTableTest(t, test)
}

func runRepeatFuncTableTest[T comparable](t *testing.T, test repeatFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Limit(iters.RepeatFunc(test.fn), test.count))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("RepeatFunc: expected %v, got %v", test.expected, got)
		}
	})
}

type repeatNTableTest[T comparable] struct {
	name     string
	value    T
	count    int
	expected []T
}

func (test repeatNTableTest[T]) Run(t *testing.T) {
	runRepeatNTableTest(t, test)
}

func runRepeatNTableTest[T comparable](t *testing.T, test repeatNTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.RepeatN(test.value, test.count))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("RepeatN: expected %v, got %v", test.expected, got)
		}
	})
}

func TestRepeat(t *testing.T) {
	tests := []runnableTest{
		repeatTableTest[int]{name: "three repeats", value: 5, count: 3, expected: []int{5, 5, 5}},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestRepeatFunc(t *testing.T) {
	counter := 0
	tests := []runnableTest{
		repeatFuncTableTest[int]{
			name: "incrementing",
			fn: func() int {
				counter++
				return counter
			},
			count:    3,
			expected: []int{1, 2, 3},
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestRepeatN(t *testing.T) {
	tests := []runnableTest{
		repeatNTableTest[string]{name: "exact count", value: "x", count: 2, expected: []string{"x", "x"}},
		repeatNTableTest[string]{name: "non-positive", value: "x", count: 0, expected: nil},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
