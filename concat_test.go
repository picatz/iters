package iters_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/picatz/iters"
	"iter"
)

func ExampleConcat_numbers() {
	seq1 := []int{1, 2, 3}
	seq2 := []int{4, 5, 6}

	concatenated := iters.Concat(
		slices.Values(seq1),
		slices.Values(seq2),
	)

	result := slices.Collect(concatenated)

	fmt.Println(result)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleConcat_keyValuePairs() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map[string]int{
		"three": 3,
		"four":  4,
	}

	concatenated := iters.Concat2(
		maps.All(map1),
		maps.All(map2),
	)

	result := maps.Collect(concatenated)

	fmt.Println(result["one"], result["two"], result["three"], result["four"])
	// Output:
	// 1 2 3 4
}

type concatTableTest[T comparable] struct {
	name     string
	seqs     [][]T
	expected []T
}

func (test concatTableTest[T]) Run(t *testing.T) {
	runConcatTableTest(t, test)
}

func runConcatTableTest[T comparable](t *testing.T, test concatTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		var seqs []iter.Seq[T]
		for _, slice := range test.seqs {
			value := slice
			seqs = append(seqs, slices.Values(value))
		}
		got := slices.Collect(iters.Concat(seqs...))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Concat: expected %v, got %v", test.expected, got)
		}
	})
}

type concat2TableTest[K comparable, V comparable] struct {
	name           string
	seqs           []map[K]V
	expectedKeys   []K
	expectedValues []V
}

func (test concat2TableTest[K, V]) Run(t *testing.T) {
	runConcat2TableTest(t, test)
}

func runConcat2TableTest[K comparable, V comparable](t *testing.T, test concat2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		var seqs []iter.Seq2[K, V]
		for _, m := range test.seqs {
			seqs = append(seqs, maps.All(m))
		}

		var gotKeys []K
		var gotValues []V
		for k, v := range iters.Concat2(seqs...) {
			gotKeys = append(gotKeys, k)
			gotValues = append(gotValues, v)
		}

		if !slices.Equal(gotKeys, test.expectedKeys) {
			t.Fatalf("Concat2 keys: expected %v, got %v", test.expectedKeys, gotKeys)
		}
		if !slices.Equal(gotValues, test.expectedValues) {
			t.Fatalf("Concat2 values: expected %v, got %v", test.expectedValues, gotValues)
		}
	})
}

func TestConcat(t *testing.T) {
	tests := []runnableTest{
		concatTableTest[int]{
			name: "two slices",
			seqs: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: []int{1, 2, 3, 4},
		},
		concatTableTest[int]{
			name: "with empty",
			seqs: [][]int{
				nil,
				{5},
			},
			expected: []int{5},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

func TestConcat2(t *testing.T) {
	tests := []runnableTest{
		concat2TableTest[string, int]{
			name: "merge maps",
			seqs: []map[string]int{
				{"a": 1},
				{"b": 2},
			},
			expectedKeys:   []string{"a", "b"},
			expectedValues: []int{1, 2},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
