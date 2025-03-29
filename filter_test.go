package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleFilter_numbers() {
	// Example usage of the Filter function to filter a slice of integers.
	numbers := []int{1, 2, 3, 4, 5}

	// Define the filter function to keep only odd numbers.
	oddNumbers := iters.Filter(
		// Convert the slice to an iter.Seq[int]
		slices.Values(numbers),
		func(num int) bool {
			return num%2 != 0 // Keep only odd numbers
		},
	)

	// Collect the filtered results into a slice.
	result := slices.Collect(oddNumbers)

	fmt.Println(result)
	// Output:
	// [1 3 5]
}

func ExampleFilter_structs() {
	// Example usage of the Filter function to filter a slice of structs,
	// which are animals with a name and number of legs in this case.
	type Animal struct {
		Name string
		Legs int
	}

	animals := []Animal{
		{"cat", 4},
		{"dog", 4},
		{"fish", 0},
		{"bird", 2},
	}

	// Define the filter function to keep only animals with more than 2 legs.
	filteredAnimals := iters.Filter(
		// Convert the slice to an iter.Seq[Animal]
		slices.Values(animals),
		func(animal Animal) bool {
			return animal.Legs > 2 // Keep animals with more than 2 legs
		},
	)

	// Collect the filtered results into a slice.
	result := slices.Collect(filteredAnimals)

	fmt.Println(result)
	// Output:
	// [{cat 4} {dog 4}]
}

// filterTableTest is a struct used for testing the Filter function.
//
// It holds the name of the test, input data, a function to apply to
// each element of the input data to determine if it should be included
// in the output, the expected output data, and a function to compare
// the expected output with the actual output.
type filterTableTest[T comparable] struct {
	name     string
	in       []T
	filterFn func(T) bool
	out      []T
	eqFn     func(expected, got []T) bool
}

// Run implements the runnableTest interface for filterTableTest.
func (test filterTableTest[T]) Run(t *testing.T) {
	runFilterTableTest(t, test)
}

// runMapTableTest runs a single test for the Map function.
func runFilterTableTest[T comparable](t *testing.T, test filterTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Filter(slices.Values(test.in), test.filterFn))

		if !test.eqFn(
			test.out,
			got,
		) {
			t.Errorf("expected output %#+v, got %#+v", test.out, got)
		}
	})
}

func TestFilter(t *testing.T) {
	tests := []runnableTest{
		filterTableTest[int]{
			name: "filter out even numbers",
			in: []int{
				1,
				2,
				3,
				4,
				5,
			},
			filterFn: func(i int) bool {
				return i%2 != 0 // is odd
			},
			out: []int{
				1,
				3,
				5,
			},
			eqFn: slices.Equal[[]int],
		},
		filterTableTest[string]{
			name: "filter out strings with length less than 4",
			in: []string{
				"cat",
				"dog",
				"fish",
				"bird",
				"elephant",
			},
			filterFn: func(s string) bool {
				return len(s) >= 4 // keep strings with length 4 or more
			},
			out: []string{
				"fish",
				"bird",
				"elephant",
			},
			eqFn: slices.Equal[[]string],
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
