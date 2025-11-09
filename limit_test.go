package iters_test

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleLimit_numbers() {
	numbers := []int{1, 2, 3, 4, 5}

	limited := iters.Limit(
		slices.Values(numbers),
		3,
	)

	result := slices.Collect(limited)

	fmt.Println(result)
	// Output:
	// [1 2 3]
}

func ExampleLimit_keyValuePairs() {
	dictionary := map[string]string{
		"apple":  "A fruit",
		"banana": "Another fruit",
		"carrot": "A vegetable",
		"date":   "A sweet fruit",
		"egg":    "A protein source",
	}

	keys := slices.Collect(maps.Keys(dictionary))
	slices.Sort(keys)
	limited := iters.Limit2(
		func(yield func(string, string) bool) {
			for _, key := range keys {
				if !yield(key, dictionary[key]) {
					return
				}
			}
		},
		2,
	)

	resultKeys := []string{}
	resultValues := []string{}
	for key, value := range limited {
		resultKeys = append(resultKeys, key)
		resultValues = append(resultValues, value)
	}

	fmt.Printf("%q\n", resultKeys)
	fmt.Printf("%q\n", resultValues)
	// Output:
	// ["apple" "banana"]
	// ["A fruit" "Another fruit"]
}

type limitTableTest[T comparable] struct {
	name     string
	input    []T
	n        int
	expected []T
}

func (test limitTableTest[T]) Run(t *testing.T) {
	runLimitTableTest(t, test)
}

func runLimitTableTest[T comparable](t *testing.T, test limitTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Limit(slices.Values(test.input), test.n))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Limit: expected %v, got %v", test.expected, got)
		}
	})
}

type limit2TableTest[K cmp.Ordered, V comparable] struct {
	name           string
	input          map[K]V
	n              int
	expectedKeys   []K
	expectedValues []V
}

func (test limit2TableTest[K, V]) Run(t *testing.T) {
	runLimit2TableTest(t, test)
}

func runLimit2TableTest[K cmp.Ordered, V comparable](t *testing.T, test limit2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		keys := slices.Collect(maps.Keys(test.input))
		slices.Sort(keys)
		seq2 := func(yield func(K, V) bool) {
			for _, key := range keys {
				if !yield(key, test.input[key]) {
					return
				}
			}
		}

		var gotKeys []K
		var gotValues []V
		for k, v := range iters.Limit2(seq2, test.n) {
			gotKeys = append(gotKeys, k)
			gotValues = append(gotValues, v)
		}

		if !slices.Equal(gotKeys, test.expectedKeys) {
			t.Fatalf("Limit2 keys: expected %v, got %v", test.expectedKeys, gotKeys)
		}
		if !slices.Equal(gotValues, test.expectedValues) {
			t.Fatalf("Limit2 values: expected %v, got %v", test.expectedValues, gotValues)
		}
	})
}

func TestLimit(t *testing.T) {
	tests := []runnableTest{
		limitTableTest[int]{
			name:     "limit smaller than length",
			input:    []int{1, 2, 3, 4},
			n:        2,
			expected: []int{1, 2},
		},
		limitTableTest[int]{
			name:     "limit zero",
			input:    []int{1, 2},
			n:        0,
			expected: nil,
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

func TestLimit2(t *testing.T) {
	tests := []runnableTest{
		limit2TableTest[string, string]{
			name: "take first two pairs",
			input: map[string]string{
				"a": "apple",
				"b": "banana",
				"c": "cherry",
			},
			n:              2,
			expectedKeys:   []string{"a", "b"},
			expectedValues: []string{"apple", "banana"},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
