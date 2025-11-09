package iters_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleReduce_sum() {
	// Example usage of the Reduce function to sum a slice of integers.
	numbers := []int{1, 2, 3, 4, 5}

	// Define the reduce function to sum the numbers.
	sum := iters.Reduce(
		// Convert the slice to an iter.Seq[int]
		slices.Values(numbers),
		func(acc int, num int) int {
			return acc + num // Sum the numbers
		},
		// Initial value for the sum, which is 0 in this case.
		0,
	)

	fmt.Println(sum)
	// Output: 15
}

func ExampleReduce2_collect() {
	input := map[string]int{
		"a": 1,
		"b": 2,
	}

	sum := iters.Reduce2(
		maps.All(input),
		func(acc int, key string, value int) int {
			return acc + value
		},
		0,
	)

	fmt.Println(sum)
	// Output:
	// 3
}

// reduceTableTest is a struct used for testing the Reduce function.
//
// It holds the name of the test, input data, a function to apply to
// each element of the input data to reduce it to a single value,
// the initial value for the reduction, and the expected output data.
type reduceTableTest[T, R comparable] struct {
	name     string
	in       []T
	reduceFn func(R, T) R
	initial  R
	out      R
}

// Run implements the runnableTest interface for reduceTableTest.
func (test reduceTableTest[T, R]) Run(t *testing.T) {
	runReduceTableTest(t, test)
}

// runReduceTableTest runs a single test for the Reduce function.
func runReduceTableTest[T, R comparable](t *testing.T, test reduceTableTest[T, R]) {
	t.Run(test.name, func(t *testing.T) {
		// Call the Reduce function with the input data and the provided function.
		got := iters.Reduce(
			slices.Values(test.in), // Convert the input slice to an iter.Seq[T]
			test.reduceFn,          // Pass the function to apply to each element
			test.initial,           // Initial value for reduction
		)

		// Compare the expected output with the actual output.
		if fmt.Sprintf("%v", test.out) != fmt.Sprintf("%v", got) {
			t.Errorf("expected output %#+v, got %#+v", test.out, got)
		}
	})
}

func TestReduce(t *testing.T) {
	// Define a slice of tests for the Reduce function.
	tests := []runnableTest{
		reduceTableTest[int, int]{
			name: "sum of numbers",
			in: []int{
				1,
				2,
				3,
				4,
				5,
			},
			reduceFn: func(acc int, num int) int {
				return acc + num // Sum the numbers
			},
			initial: 0,  // Initial value for the sum
			out:     15, // Expected output (1+2+3+4+5)
		},
		reduceTableTest[string, string]{
			name: "concatenate strings",
			in: []string{
				"hello",
				" ",
				"world",
			},
			reduceFn: func(acc string, str string) string {
				return acc + str // Concatenate the strings
			},
			initial: "",            // Initial value for concatenation
			out:     "hello world", // Expected output
		},
	}

	// Run each test in the slice.
	for _, test := range tests {
		test.Run(t)
	}
}

type reduce2TableTest[K comparable, V comparable, R comparable] struct {
	name     string
	input    map[K]V
	fn       func(R, K, V) R
	initial  R
	expected R
}

func (test reduce2TableTest[K, V, R]) Run(t *testing.T) {
	runReduce2TableTest(t, test)
}

func runReduce2TableTest[K comparable, V comparable, R comparable](t *testing.T, test reduce2TableTest[K, V, R]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.Reduce2(maps.All(test.input), test.fn, test.initial)
		if got != test.expected {
			t.Fatalf("Reduce2: expected %v, got %v", test.expected, got)
		}
	})
}

func TestReduce2(t *testing.T) {
	tests := []runnableTest{
		reduce2TableTest[string, int, int]{
			name: "sum values",
			input: map[string]int{
				"a": 1,
				"b": 2,
			},
			fn:       func(acc int, _ string, v int) int { return acc + v },
			initial:  0,
			expected: 3,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
