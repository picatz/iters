package iters_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleReusable_twice() {
	reusable := iters.Reusable(slices.Values([]int{1, 2, 3}))

	first := slices.Collect(reusable)
	second := slices.Collect(reusable)

	fmt.Println(first, second)
	// Output:
	// [1 2 3] [1 2 3]
}

func ExampleReusable2_pairs() {
	source := map[string]int{"a": 1, "b": 2}
	reusable := iters.Reusable2(maps.All(source))

	first := maps.Collect(reusable)
	second := maps.Collect(reusable)

	fmt.Println(first["a"], second["b"])
	// Output:
	// 1 2
}

type reusableTableTest[T comparable] struct {
	name     string
	input    []T
	expected []T
}

func (test reusableTableTest[T]) Run(t *testing.T) {
	runReusableTableTest(t, test)
}

func runReusableTableTest[T comparable](t *testing.T, test reusableTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		reusable := iters.Reusable(slices.Values(test.input))
		first := slices.Collect(reusable)
		second := slices.Collect(reusable)
		if !slices.Equal(first, test.expected) || !slices.Equal(second, test.expected) {
			t.Fatalf("Reusable: expected %v twice, got %v and %v", test.expected, first, second)
		}
	})
}

type reusable2TableTest[K comparable, V comparable] struct {
	name     string
	input    map[K]V
	expected map[K]V
}

func (test reusable2TableTest[K, V]) Run(t *testing.T) {
	runReusable2TableTest(t, test)
}

func runReusable2TableTest[K comparable, V comparable](t *testing.T, test reusable2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		reusable := iters.Reusable2(maps.All(test.input))
		first := maps.Collect(reusable)
		second := maps.Collect(reusable)
		if !maps.Equal(first, test.expected) || !maps.Equal(second, test.expected) {
			t.Fatalf("Reusable2: expected %v twice, got %v and %v", test.expected, first, second)
		}
	})
}

func TestReusable(t *testing.T) {
	tests := []runnableTest{
		reusableTableTest[int]{name: "slices", input: []int{1, 2}, expected: []int{1, 2}},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestReusable2(t *testing.T) {
	tests := []runnableTest{
		reusable2TableTest[string, int]{
			name: "maps",
			input: map[string]int{
				"a": 1,
				"b": 2,
			},
			expected: map[string]int{
				"a": 1,
				"b": 2,
			},
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
