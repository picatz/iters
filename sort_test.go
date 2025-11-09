package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"cmp"
	"github.com/picatz/iters"
)

func ExampleSort_numbers() {
	sorted := slices.Collect(iters.Sort(slices.Values([]int{3, 1, 2})))
	fmt.Println(sorted)
	// Output:
	// [1 2 3]
}

func ExampleSortFunc_reverse() {
	sorted := slices.Collect(
		iters.SortFunc(
			slices.Values([]int{1, 2, 3}),
			func(a, b int) int { return b - a },
		),
	)
	fmt.Println(sorted)
	// Output:
	// [3 2 1]
}

type sortTableTest[T cmp.Ordered] struct {
	name     string
	input    []T
	expected []T
}

func (test sortTableTest[T]) Run(t *testing.T) {
	runSortTableTest(t, test)
}

func runSortTableTest[T cmp.Ordered](t *testing.T, test sortTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Sort(slices.Values(test.input)))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Sort: expected %v, got %v", test.expected, got)
		}
	})
}

type sortFuncTableTest[T any] struct {
	name     string
	input    []T
	cmp      func(a, b T) int
	expected []T
}

func (test sortFuncTableTest[T]) Run(t *testing.T) {
	runSortFuncTableTest(t, test)
}

func runSortFuncTableTest[T any](t *testing.T, test sortFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.SortFunc(slices.Values(test.input), test.cmp))
		if fmt.Sprint(got) != fmt.Sprint(test.expected) {
			t.Fatalf("SortFunc: expected %v, got %v", test.expected, got)
		}
	})
}

func TestSort(t *testing.T) {
	tests := []runnableTest{
		sortTableTest[int]{name: "ascending", input: []int{3, 1, 2}, expected: []int{1, 2, 3}},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestSortFunc(t *testing.T) {
	tests := []runnableTest{
		sortFuncTableTest[int]{name: "descending", input: []int{1, 2, 3}, cmp: func(a, b int) int { return b - a }, expected: []int{3, 2, 1}},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
