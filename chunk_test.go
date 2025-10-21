package iters_test

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleChunk_numbers() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}

	chunkSize := 3

	chunked := iters.Chunk(
		slices.Values(numbers),
		chunkSize,
	)

	result := slices.Collect(chunked)

	fmt.Println(result)
	// Output:
	// [[1 2 3] [4 5 6] [7]]
}

func ExampleChunk_keyValuePairs() {
	dictionary := map[string]string{
		"apple":  "A fruit",
		"banana": "Another fruit",
		"carrot": "A vegetable",
		"date":   "A sweet fruit",
		"egg":    "A protein source",
	}

	chunkSize := 2

	keys := slices.Collect(maps.Keys(dictionary))
	slices.Sort(keys)
	chunked := iters.Chunk2(
		func(yield func(string, string) bool) {
			for _, key := range keys {
				if !yield(key, dictionary[key]) {
					return
				}
			}
		},
		chunkSize,
	)

	resultKeys := [][]string{}
	resultValues := [][]string{}
	for keys, values := range chunked {
		resultKeys = append(resultKeys, keys)
		resultValues = append(resultValues, values)
	}

	fmt.Printf("%q\n", resultKeys)
	fmt.Printf("%q\n", resultValues)
	// Output:
	// [["apple" "banana"] ["carrot" "date"] ["egg"]]
	// [["A fruit" "Another fruit"] ["A vegetable" "A sweet fruit"] ["A protein source"]]
}

type chunkTableTest[T any] struct {
	name     string
	input    []T
	size     int
	eq       func(T, T) bool
	expected [][]T
}

func (test chunkTableTest[T]) Run(t *testing.T) {
	runChunkTableTest(t, test, test.eq)
}

func runChunkTableTest[T any](t *testing.T, test chunkTableTest[T], eq func(T, T) bool) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Chunk(slices.Values(test.input), test.size))

		if !slices.EqualFunc(got, test.expected, func(a, b []T) bool {
			return slices.EqualFunc(a, b, eq)
		}) {
			t.Errorf("expected output %#+v, got %#+v", test.expected, got)
		}
	})
}

func TestChunk(t *testing.T) {
	tests := []runnableTest{
		chunkTableTest[int]{
			name:  "chunk size 2",
			input: []int{1, 2, 3, 4, 5},
			size:  2,
			eq:    func(a, b int) bool { return a == b },
			expected: [][]int{
				{1, 2},
				{3, 4},
				{5},
			},
		},
		chunkTableTest[string]{
			name:  "chunk size 3",
			input: []string{"a", "b", "c", "d", "e", "f", "g"},
			size:  3,
			eq:    func(a, b string) bool { return a == b },
			expected: [][]string{
				{"a", "b", "c"},
				{"d", "e", "f"},
				{"g"},
			},
		},
		chunkTableTest[int]{
			name:     "chunk size 0",
			input:    []int{1, 2, 3},
			size:     0,
			eq:       func(a, b int) bool { return a == b },
			expected: [][]int{},
		},
		chunkTableTest[int]{
			name:     "chunk size negative",
			input:    []int{1, 2, 3},
			size:     -2,
			eq:       func(a, b int) bool { return a == b },
			expected: [][]int{},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

type chunk2TableTest[K comparable, V any] struct {
	name           string
	input          map[K]V
	size           int
	keyEq          func(K, K) bool
	valueEq        func(V, V) bool
	expectedKeys   [][]K
	expectedValues [][]V
}

func (test chunk2TableTest[K, V]) Run(t *testing.T) {
	runChunk2TableTest(t, test, test.keyEq, test.valueEq)
}

func runChunk2TableTest[K comparable, V any](t *testing.T, test chunk2TableTest[K, V], keyEq func(K, K) bool, valueEq func(V, V) bool) {
	t.Run(test.name, func(t *testing.T) {
		order := make([]K, 0, len(test.input))
		if len(test.expectedKeys) > 0 {
			seen := make(map[K]struct{}, len(test.input))
			for _, chunk := range test.expectedKeys {
				for _, key := range chunk {
					if _, ok := seen[key]; ok {
						continue
					}
					seen[key] = struct{}{}
					order = append(order, key)
				}
			}
			for key := range test.input {
				if _, ok := seen[key]; ok {
					continue
				}
				seen[key] = struct{}{}
				order = append(order, key)
			}
		} else {
			order = slices.Collect(maps.Keys(test.input))
		}
		seq2 := func(yield func(K, V) bool) {
			for _, key := range order {
				if !yield(key, test.input[key]) {
					return
				}
			}
		}

		gotKeys := [][]K{}
		gotValues := [][]V{}
		for ks, vs := range iters.Chunk2(seq2, test.size) {
			gotKeys = append(gotKeys, ks)
			gotValues = append(gotValues, vs)
		}

		if !slices.EqualFunc(gotKeys, test.expectedKeys, func(a, b []K) bool {
			return slices.EqualFunc(a, b, keyEq)
		}) {
			t.Errorf("expected keys %#+v, got %#+v", test.expectedKeys, gotKeys)
		}

		if !slices.EqualFunc(gotValues, test.expectedValues, func(a, b []V) bool {
			return slices.EqualFunc(a, b, valueEq)
		}) {
			t.Errorf("expected values %#+v, got %#+v", test.expectedValues, gotValues)
		}
	})
}

func TestChunk2(t *testing.T) {
	tests := []runnableTest{
		chunk2TableTest[string, int]{
			name: "chunk size 2",
			input: map[string]int{
				"a": 1,
				"b": 2,
				"c": 3,
				"d": 4,
				"e": 5,
			},
			size:    2,
			keyEq:   func(a, b string) bool { return a == b },
			valueEq: func(a, b int) bool { return a == b },
			expectedKeys: [][]string{
				{"a", "b"},
				{"c", "d"},
				{"e"},
			},
			expectedValues: [][]int{
				{1, 2},
				{3, 4},
				{5},
			},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
