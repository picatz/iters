package iters_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func Example_simple_map() {
	for v := range iters.Map(
		slices.Values(
			[]int{1, 2, 3, 4},
		),
		func(i int) string {
			return fmt.Sprintf("%d ", i)
		},
	) {
		fmt.Print(v)
	}

	// Output:
	// 1 2 3 4
}

func Example_simple_map2() {
	// Use Map2 to transform the input map into a new map
	result := maps.Collect(
		iters.Map2(
			maps.All(
				map[string]int{"a": 1, "b": 2, "c": 3},
			),
			func(k string, v int) (string, int) {
				return k + "_new", v * 10
			},
		),
	)

	// For stable output, sort the keys of the resulting map
	// so we can iterate over the result map in a stable order.
	resultKeys := slices.Collect(maps.Keys(result))
	slices.Sort(resultKeys)

	// Collect and print the results in sorted order
	for _, k := range resultKeys {
		// Print each key-value pair in the new map
		fmt.Printf("%s: %d ", k, result[k])
	}

	// Output:
	// a_new: 10 b_new: 20 c_new: 30
}

// mapTableTest is used for testing the Map function.
//
// It holds the name of the test, input data, a function to apply to
// each element of the input data, and the expected output data. The
// type parameters T and R allow for flexibility in the types of
// input and output data, but must satisfy the constraint of being
// "[comparable]" to ensure that we can compare the output with the
// expected output easily.
//
// [comparable]: https://go.dev/ref/spec#Comparison_operators
type mapTableTest[T, R comparable] struct {
	name  string
	in    []T
	mapFn func(T) R
	out   []R
	eqFn  func(expected, got []R) bool
}

// Run implements the runnableTest interface for mapTableTest.
func (test mapTableTest[T, R]) Run(t *testing.T) {
	runMapTableTest(t, test)
}

// runMapTableTest runs a single test for the Map function.
func runMapTableTest[T, R comparable](t *testing.T, test mapTableTest[T, R]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Map(slices.Values(test.in), test.mapFn))

		if !test.eqFn(
			test.out,
			got,
		) {
			t.Errorf("expected output %#+v, got %#+v", test.out, got)
		}
	})
}

// map2TableTest is used for testing the Map2 function.
//
// It holds the name of the test, input data (which is a map for simplicity),
// a function to apply to each key-value pair of the input data, and the expected
// output data.
//
// The type parameters K1, V1, K2, and V2 allow for flexibility in the types of
// input and output data, but they must satisfy the constraint of being
// "[comparable]" to ensure that we can compare the output with the expected
// output easily.
//
// [comparable]: https://go.dev/ref/spec#Comparison_operators
type map2TableTest[K1, V1, K2, V2 comparable] struct {
	name   string
	in     map[K1]V1
	map2Fn func(K1, V1) (K2, V2)
	out    map[K2]V2
	eqFn   func(a, b map[K2]V2) bool
}

// Run implements the runnableTest interface for map2TableTest.
func (test map2TableTest[K1, V1, K2, V2]) Run(t *testing.T) {
	runMap2TableTest(t, test)
}

// runMap2TableTest runs a single test for the Map2 function.
func runMap2TableTest[K1, V1, K2, V2 comparable](t *testing.T, test map2TableTest[K1, V1, K2, V2]) {
	t.Run(test.name, func(t *testing.T) {
		// Convert the input map to an iterator sequence
		seq2 := iters.Map2(
			maps.All(test.in),
			func(k K1, v V1) (K2, V2) {
				return test.map2Fn(k, v)
			},
		)

		// Collect the output from the Map2 function
		got := maps.Collect(seq2)

		if !test.eqFn(
			test.out,
			got,
		) {
			t.Errorf("expected output %#+v, got %#+v", test.out, got)
		}
	})
}

func TestMap(t *testing.T) {
	tests := []runnableTest{
		mapTableTest[int, string]{
			name: "convert each element to string",
			in: []int{
				1,
				2,
				3,
				4,
			},
			mapFn: func(i int) string {
				return fmt.Sprintf("%d", i)
			},
			out: []string{
				"1",
				"2",
				"3",
				"4",
			},
			eqFn: slices.Equal[[]string],
		},
		mapTableTest[int, int]{
			name: "double each element",
			in: []int{
				1,
				2,
				3,
				4,
			},
			mapFn: func(x int) int {
				return x * 2
			},
			out: []int{
				2,
				4,
				6,
				8,
			},
			eqFn: slices.Equal[[]int],
		},
		mapTableTest[string, int]{
			name: "convert each string to its length",
			in: []string{
				"hello",
				"world",
				"go",
				"iters",
			},
			mapFn: func(s string) int {
				return len(s)
			},
			out: []int{
				5,
				5,
				2,
				5,
			},
			eqFn: slices.Equal[[]int],
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

func TestMap2(t *testing.T) {
	tests := []runnableTest{
		map2TableTest[string, int, string, int]{
			name: "convert each key-value pair to a new key-value pair",
			in: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
			},
			map2Fn: func(k string, v int) (string, int) {
				return k + "_new", v * 10
			},
			out: map[string]int{
				"a_new": 10,
				"b_new": 20,
				"c_new": 30,
			},
			eqFn: maps.Equal[map[string]int, map[string]int],
		},
		map2TableTest[int, int, string, string]{
			name: "convert each key-value pair to a new key-value pair",
			in: map[int]int{
				1: 0,
				2: 1,
				3: 2,
			},
			map2Fn: func(k int, v int) (string, string) {
				return fmt.Sprintf("key_%d", k), fmt.Sprintf("value_%d", v)
			},
			out: map[string]string{
				"key_1": "value_0",
				"key_2": "value_1",
				"key_3": "value_2",
			},
			eqFn: maps.Equal[map[string]string, map[string]string],
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
