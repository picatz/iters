package iters_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/picatz/iters"
)

func ExampleCompact_numbers() {
	seq := iters.Compact(slices.Values([]int{1, 1, 2, 2, 3, 1}))
	fmt.Println(slices.Collect(seq))
	// Output:
	// [1 2 3 1]
}

func ExampleCompactFunc_caseInsensitive() {
	seq := iters.CompactFunc(
		slices.Values([]string{"Go", "go", "iters", "ITERS"}),
		func(a, b string) bool { return strings.EqualFold(a, b) },
	)
	fmt.Println(slices.Collect(seq))
	// Output:
	// [Go iters]
}

func ExampleCompact2_pairs() {
	seq2 := func(yield func(string, int) bool) {
		pairs := []struct {
			key string
			val int
		}{
			{"a", 1},
			{"a", 1},
			{"b", 2},
			{"b", 2},
		}
		for _, p := range pairs {
			if !yield(p.key, p.val) {
				return
			}
		}
	}

	var keys []string
	var values []int
	for k, v := range iters.Compact2(seq2) {
		keys = append(keys, k)
		values = append(values, v)
	}

	fmt.Println(keys)
	fmt.Println(values)
	// Output:
	// [a b]
	// [1 2]
}

func ExampleCompactFunc2_pairs() {
	seq2 := func(yield func(string, string) bool) {
		pairs := []struct {
			key string
			val string
		}{
			{"a", "One"},
			{"a", "one"},
			{"b", "Two"},
		}
		for _, p := range pairs {
			if !yield(p.key, p.val) {
				return
			}
		}
	}

	var keys []string
	var values []string
	for k, v := range iters.CompactFunc2(
		seq2,
		func(k1, v1, k2, v2 string) bool {
			return strings.EqualFold(k1, k2) && strings.EqualFold(v1, v2)
		},
	) {
		keys = append(keys, k)
		values = append(values, v)
	}

	fmt.Println(keys)
	fmt.Println(values)
	// Output:
	// [a b]
	// [One Two]
}

type compactTableTest[T comparable] struct {
	name     string
	input    []T
	expected []T
}

func (test compactTableTest[T]) Run(t *testing.T) {
	runCompactTableTest(t, test)
}

func runCompactTableTest[T comparable](t *testing.T, test compactTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.Compact(slices.Values(test.input)))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Compact: expected %v, got %v", test.expected, got)
		}
	})
}

type compactFuncTableTest[T any] struct {
	name     string
	input    []T
	equal    func(T, T) bool
	expected []T
}

func (test compactFuncTableTest[T]) Run(t *testing.T) {
	runCompactFuncTableTest(t, test)
}

func runCompactFuncTableTest[T any](t *testing.T, test compactFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := slices.Collect(iters.CompactFunc(slices.Values(test.input), test.equal))
		if fmt.Sprint(got) != fmt.Sprint(test.expected) {
			t.Fatalf("CompactFunc: expected %v, got %v", test.expected, got)
		}
	})
}

type compact2TableTest[K comparable, V comparable] struct {
	name  string
	input []struct {
		key K
		val V
	}
	expected []struct {
		key K
		val V
	}
}

func (test compact2TableTest[K, V]) Run(t *testing.T) {
	runCompact2TableTest(t, test)
}

func runCompact2TableTest[K comparable, V comparable](t *testing.T, test compact2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		seq2 := func(yield func(K, V) bool) {
			for _, p := range test.input {
				if !yield(p.key, p.val) {
					return
				}
			}
		}
		var got []struct {
			key K
			val V
		}
		for k, v := range iters.Compact2(seq2) {
			got = append(got, struct {
				key K
				val V
			}{k, v})
		}
		if fmt.Sprint(got) != fmt.Sprint(test.expected) {
			t.Fatalf("Compact2: expected %v, got %v", test.expected, got)
		}
	})
}

type compactFunc2TableTest[K any, V any] struct {
	name  string
	input []struct {
		key K
		val V
	}
	equal    func(K, V, K, V) bool
	expected []struct {
		key K
		val V
	}
}

func (test compactFunc2TableTest[K, V]) Run(t *testing.T) {
	runCompactFunc2TableTest(t, test)
}

func runCompactFunc2TableTest[K any, V any](t *testing.T, test compactFunc2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		seq2 := func(yield func(K, V) bool) {
			for _, p := range test.input {
				if !yield(p.key, p.val) {
					return
				}
			}
		}
		var got []struct {
			key K
			val V
		}
		for k, v := range iters.CompactFunc2(seq2, test.equal) {
			got = append(got, struct {
				key K
				val V
			}{k, v})
		}
		if fmt.Sprint(got) != fmt.Sprint(test.expected) {
			t.Fatalf("CompactFunc2: expected %v, got %v", test.expected, got)
		}
	})
}

func TestCompact(t *testing.T) {
	tests := []runnableTest{
		compactTableTest[int]{name: "remove duplicates", input: []int{1, 1, 2, 2, 3}, expected: []int{1, 2, 3}},
		compactTableTest[int]{name: "single elements", input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestCompactFunc(t *testing.T) {
	tests := []runnableTest{
		compactFuncTableTest[string]{
			name:  "case insensitive",
			input: []string{"Go", "go", "iters"},
			equal: strings.EqualFold,
			expected: []string{
				"Go",
				"iters",
			},
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestCompact2(t *testing.T) {
	tests := []runnableTest{
		compact2TableTest[string, int]{
			name: "pairs",
			input: []struct {
				key string
				val int
			}{
				{"a", 1},
				{"a", 1},
				{"b", 2},
			},
			expected: []struct {
				key string
				val int
			}{
				{"a", 1},
				{"b", 2},
			},
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestCompactFunc2(t *testing.T) {
	tests := []runnableTest{
		compactFunc2TableTest[string, string]{
			name: "case insensitive pairs",
			input: []struct {
				key string
				val string
			}{
				{"a", "One"},
				{"a", "one"},
				{"b", "two"},
			},
			equal: func(k1, v1, k2, v2 string) bool {
				return strings.EqualFold(k1, k2) && strings.EqualFold(v1, v2)
			},
			expected: []struct {
				key string
				val string
			}{
				{"a", "One"},
				{"b", "two"},
			},
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
