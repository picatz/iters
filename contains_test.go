package iters_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleContains_numbers() {
	numbers := []int{1, 2, 3, 4, 5}

	hasThree := iters.Contains(
		slices.Values(numbers),
		3,
	)

	hasSix := iters.Contains(
		slices.Values(numbers),
		6,
	)

	fmt.Println(hasThree)
	fmt.Println(hasSix)
	// Output:
	// true
	// false
}

func ExampleContainsFunc_strings() {
	fruit := []string{"apple", "banana", "cherry", "date"}

	hasLongString := iters.ContainsFunc(
		slices.Values(fruit),
		func(s string) bool {
			return len(s) > 5
		},
	)

	hasBString := iters.ContainsFunc(
		slices.Values(fruit),
		func(s string) bool {
			return s[0] == 'b'
		},
	)

	fmt.Println(hasLongString)
	fmt.Println(hasBString)
	// Output:
	// true
	// true
}

func ExampleContains2_keyValuePairs() {
	dictionary := map[string]string{
		"apple":  "A fruit",
		"banana": "Another fruit",
		"carrot": "A vegetable",
	}

	hasAppleDefinition := iters.Contains2(
		maps.All(dictionary),
		"apple",
		"A fruit",
	)

	hasBananaWrongDefinition := iters.Contains2(
		maps.All(dictionary),
		"banana",
		"A yellow fruit",
	)

	fmt.Println(hasAppleDefinition)
	fmt.Println(hasBananaWrongDefinition)
	// Output:
	// true
	// false
}

func ExampleContainsFunc2_keyValuePairs() {
	dictionary := map[string]string{
		"apple":  "A fruit",
		"banana": "Another fruit",
		"carrot": "A vegetable",
	}

	hasDefinitionWithA := iters.ContainsFunc2(
		maps.All(dictionary),
		func(k, v string) bool {
			return v[0] == 'A'
		},
	)

	hasDefinitionWithZ := iters.ContainsFunc2(
		maps.All(dictionary),
		func(k, v string) bool {
			return v[0] == 'Z'
		},
	)

	fmt.Println(hasDefinitionWithA)
	fmt.Println(hasDefinitionWithZ)
	// Output:
	// true
	// false
}

type containsTableTest[V comparable] struct {
	name  string
	in    []V
	value V
	out   bool
}

// Run implements the runnableTest interface for containsTableTest.
func (test containsTableTest[V]) Run(t *testing.T) {
	runContainsTableTest(t, test)
}

// runContainsTableTest runs a single test for the Contains function.
func runContainsTableTest[V comparable](t *testing.T, test containsTableTest[V]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.Contains(
			slices.Values(test.in),
			test.value,
		)

		if got != test.out {
			t.Errorf("expected output %v, got %v", test.out, got)
		}
	})
}

func TestContains(t *testing.T) {
	tests := []runnableTest{
		containsTableTest[int]{
			name: "value exists in sequence",
			in: []int{
				10,
				20,
				30,
				40,
			},
			value: 30,
			out:   true,
		},
		containsTableTest[int]{
			name: "value does not exist in sequence",
			in: []int{
				10,
				20,
				30,
				40,
			},
			value: 50,
			out:   false,
		},
		containsTableTest[string]{
			name: "string exists in sequence",
			in: []string{
				"apple",
				"banana",
				"cherry",
			},
			value: "banana",
			out:   true,
		},
		containsTableTest[string]{
			name: "string does not exist in sequence",
			in: []string{
				"apple",
				"banana",
				"cherry",
			},
			value: "date",
			out:   false,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

type containsFuncTableTest[V any] struct {
	name string
	in   []V
	fn   func(V) bool
	out  bool
}

// Run implements the runnableTest interface for containsFuncTableTest.
func (test containsFuncTableTest[V]) Run(t *testing.T) {
	runContainsFuncTableTest(t, test)
}

// runContainsFuncTableTest runs a single test for the ContainsFunc function.
func runContainsFuncTableTest[V any](t *testing.T, test containsFuncTableTest[V]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.ContainsFunc(
			slices.Values(test.in),
			test.fn,
		)
		if got != test.out {
			t.Errorf("expected output %v, got %v", test.out, got)
		}
	})
}

func TestContainsFunc(t *testing.T) {
	tests := []runnableTest{
		containsFuncTableTest[int]{
			name: "value greater than 25 exists",
			in: []int{
				10,
				20,
				30,
				40,
			},
			fn: func(i int) bool {
				return i > 25
			},
			out: true,
		},
		containsFuncTableTest[int]{
			name: "value greater than 50 does not exist",
			in: []int{
				10,
				20,
				30,
				40,
			},
			fn: func(i int) bool {
				return i > 50
			},
			out: false,
		},
		containsFuncTableTest[string]{
			name: "string with length 6 exists",
			in: []string{
				"apple",
				"banana",
				"cherry",
			},
			fn: func(s string) bool {
				return len(s) == 6
			},
			out: true,
		},
		containsFuncTableTest[string]{
			name: "string with length 10 does not exist",
			in: []string{
				"apple",
				"banana",
				"cherry",
			},
			fn: func(s string) bool {
				return len(s) == 10
			},
			out: false,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

type contains2TableTest[K comparable, V comparable] struct {
	name  string
	in    map[K]V
	key   K
	value V
	out   bool
}

// Run implements the runnableTest interface for contains2TableTest.
func (test contains2TableTest[K, V]) Run(t *testing.T) {
	runContains2TableTest(t, test)
}

// runContains2TableTest runs a single test for the Contains2 function.
func runContains2TableTest[K comparable, V comparable](t *testing.T, test contains2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.Contains2(
			maps.All(test.in),
			test.key,
			test.value,
		)

		if got != test.out {
			t.Errorf("expected output %v, got %v", test.out, got)
		}
	})
}

func TestContains2(t *testing.T) {
	tests := []runnableTest{
		contains2TableTest[string, int]{
			name: "key-value pair exists in map",
			in: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			key:   "two",
			value: 2,
			out:   true,
		},
		contains2TableTest[string, int]{
			name: "key-value pair does not exist in map",
			in: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			key:   "two",
			value: 3,
			out:   false,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

type containsFunc2TableTest[K comparable, V any] struct {
	name string
	in   map[K]V
	fn   func(K, V) bool
	out  bool
}

// Run implements the runnableTest interface for containsFunc2TableTest.
func (test containsFunc2TableTest[K, V]) Run(t *testing.T) {
	runContainsFunc2TableTest(t, test)
}

// runContainsFunc2TableTest runs a single test for the ContainsFunc2 function.
func runContainsFunc2TableTest[K comparable, V any](t *testing.T, test containsFunc2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.ContainsFunc2(
			maps.All(test.in),
			test.fn,
		)

		if got != test.out {
			t.Errorf("expected output %v, got %v", test.out, got)
		}
	})
}

func TestContainsFunc2(t *testing.T) {
	tests := []runnableTest{
		containsFunc2TableTest[string, int]{
			name: "key-value pair with value greater than 2 exists",
			in: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			fn: func(k string, v int) bool {
				return v > 2
			},
			out: true,
		},
		containsFunc2TableTest[string, int]{
			name: "key-value pair with value greater than 5 does not exist",
			in: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			fn: func(k string, v int) bool {
				return v > 5
			},
			out: false,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
