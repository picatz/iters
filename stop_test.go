package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleStop_iteration() {
	// Example usage of the StopIteration function to stop iteration early.
	numbers := []int{1, 2, 3, 4, 5}

	// Define a function that stops iteration when it encounters the number 3.
	stoppedIteration := iters.Stop(
		// Convert the slice to an iter.Seq[int]
		slices.Values(numbers),
		func(num int) bool {
			return num == 3 // Stop iteration when the number is 3
		},
	)

	// Collect the results into a slice.
	result := slices.Collect(stoppedIteration)

	fmt.Println(result)
	// Output:
	// [1 2]
}

func ExampleStop2_pairs() {
	type pair struct {
		key string
		val int
	}
	input := []pair{
		{"a", 1},
		{"stop", 0},
		{"b", 2},
	}

	seq2 := func(yield func(string, int) bool) {
		for _, p := range input {
			if !yield(p.key, p.val) {
				return
			}
		}
	}

	stopped := iters.Stop2(seq2, func(k string, _ int) bool { return k == "stop" })

	for k, v := range stopped {
		fmt.Println(k, v)
	}
	// Output:
	// a 1
}

type stopTableTest[T comparable] struct {
	name     string
	input    []T
	stopFn   func(T) bool
	expected []T
}

func (test stopTableTest[T]) Run(t *testing.T) {
	runStopTableTest(t, test)
}

func runStopTableTest[T comparable](t *testing.T, test stopTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Stop(slices.Values(test.input), test.stopFn))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Stop: expected %v, got %v", test.expected, got)
		}
	})
}

type stop2TableTest[K comparable, V comparable] struct {
	name  string
	input []struct {
		key K
		val V
	}
	stopFn         func(K, V) bool
	expectedKeys   []K
	expectedValues []V
}

func (test stop2TableTest[K, V]) Run(t *testing.T) {
	runStop2TableTest(t, test)
}

func runStop2TableTest[K comparable, V comparable](t *testing.T, test stop2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		seq2 := func(yield func(K, V) bool) {
			for _, pair := range test.input {
				if !yield(pair.key, pair.val) {
					return
				}
			}
		}

		var gotKeys []K
		var gotValues []V
		for k, v := range iters.Stop2(seq2, test.stopFn) {
			gotKeys = append(gotKeys, k)
			gotValues = append(gotValues, v)
		}

		if !slices.Equal(gotKeys, test.expectedKeys) {
			t.Fatalf("Stop2 keys: expected %v, got %v", test.expectedKeys, gotKeys)
		}
		if !slices.Equal(gotValues, test.expectedValues) {
			t.Fatalf("Stop2 values: expected %v, got %v", test.expectedValues, gotValues)
		}
	})
}

func TestStop(t *testing.T) {
	tests := []runnableTest{
		stopTableTest[int]{
			name:     "stop at value",
			input:    []int{1, 2, 3},
			stopFn:   func(v int) bool { return v == 3 },
			expected: []int{1, 2},
		},
		stopTableTest[int]{
			name:     "predicate never true",
			input:    []int{1, 2},
			stopFn:   func(int) bool { return false },
			expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

func TestStop2(t *testing.T) {
	tests := []runnableTest{
		stop2TableTest[string, int]{
			name: "stop when key matches",
			input: []struct {
				key string
				val int
			}{
				{"a", 1},
				{"stop", 0},
				{"b", 2},
			},
			stopFn:         func(k string, _ int) bool { return k == "stop" },
			expectedKeys:   []string{"a"},
			expectedValues: []int{1},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
