package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleAverage_numbers() {
	values := []int{2, 4, 6, 8}

	fmt.Println(iters.Average(slices.Values(values)))
	// Output:
	// 5
}

func ExampleAverageFunc_lengths() {
	words := []string{"go", "iters"}

	fmt.Println(
		iters.AverageFunc(
			slices.Values(words),
			func(s string) float64 { return float64(len(s)) },
		),
	)
	// Output:
	// 3.5
}

type averageTableTest[T iters.Number] struct {
	name     string
	input    []T
	expected float64
}

func (test averageTableTest[T]) Run(t *testing.T) {
	runAverageTableTest(t, test)
}

func runAverageTableTest[T iters.Number](t *testing.T, test averageTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.Average(slices.Values(test.input))
		if got != test.expected {
			t.Fatalf("Average: expected %v, got %v", test.expected, got)
		}
	})
}

type averageFuncTableTest[T any] struct {
	name     string
	input    []T
	fn       func(T) float64
	expected float64
}

func (test averageFuncTableTest[T]) Run(t *testing.T) {
	runAverageFuncTableTest(t, test)
}

func runAverageFuncTableTest[T any](t *testing.T, test averageFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.AverageFunc(slices.Values(test.input), test.fn)
		if got != test.expected {
			t.Fatalf("AverageFunc: expected %v, got %v", test.expected, got)
		}
	})
}

func TestAverage(t *testing.T) {
	tests := []runnableTest{
		averageTableTest[int]{
			name:     "even numbers",
			input:    []int{2, 4, 6, 8},
			expected: 5,
		},
		averageTableTest[int]{
			name:     "empty input",
			input:    nil,
			expected: 0,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

func TestAverageFunc(t *testing.T) {
	tests := []runnableTest{
		averageFuncTableTest[string]{
			name:  "lengths",
			input: []string{"go", "iters"},
			fn: func(s string) float64 {
				return float64(len(s))
			},
			expected: 3.5,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
