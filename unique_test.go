package iters_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/picatz/iters"
)

func ExampleUnique_numbers() {
	fmt.Println(slices.Collect(iters.Unique(slices.Values([]int{1, 2, 2, 3}))))
	// Output:
	// [1 2 3]
}

func ExampleUniqueFunc_caseInsensitive() {
	words := []string{"Go", "go", "iters"}
	unique := iters.UniqueFunc(
		slices.Values(words),
		func(a, b string) bool { return strings.EqualFold(a, b) },
	)
	fmt.Println(slices.Collect(unique))
	// Output:
	// [Go iters]
}

type uniqueTableTest[T comparable] struct {
	name     string
	input    []T
	expected []T
}

func (test uniqueTableTest[T]) Run(t *testing.T) {
	runUniqueTableTest(t, test)
}

func runUniqueTableTest[T comparable](t *testing.T, test uniqueTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Unique(slices.Values(test.input)))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Unique: expected %v, got %v", test.expected, got)
		}
	})
}

type uniqueFuncTableTest[T any] struct {
	name     string
	input    []T
	equal    func(a, b T) bool
	expected []T
}

func (test uniqueFuncTableTest[T]) Run(t *testing.T) {
	runUniqueFuncTableTest(t, test)
}

func runUniqueFuncTableTest[T any](t *testing.T, test uniqueFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.UniqueFunc(slices.Values(test.input), test.equal))
		if fmt.Sprint(got) != fmt.Sprint(test.expected) {
			t.Fatalf("UniqueFunc: expected %v, got %v", test.expected, got)
		}
	})
}

func TestUnique(t *testing.T) {
	tests := []runnableTest{
		uniqueTableTest[int]{
			name:     "ints",
			input:    []int{1, 2, 2, 3},
			expected: []int{1, 2, 3},
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestUniqueFunc(t *testing.T) {
	tests := []runnableTest{
		uniqueFuncTableTest[string]{
			name:  "case insensitive",
			input: []string{"Go", "go", "iters"},
			equal: func(a, b string) bool { return strings.EqualFold(a, b) },
			expected: []string{
				"Go",
				"iters",
			},
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
