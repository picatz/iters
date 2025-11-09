package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleLast_value() {
	fmt.Println(iters.Last(slices.Values([]int{1, 2, 3})))
	// Output:
	// 3 true
}

func ExampleLastFunc_gt() {
	last, ok := iters.LastFunc(
		slices.Values([]int{1, 4, 2, 5}),
		func(v int) bool { return v > 3 },
	)
	fmt.Println(last, ok)
	// Output:
	// 5 true
}

func ExampleLast2_pair() {
	seq2 := func(yield func(string, int) bool) {
		data := []struct {
			key string
			val int
		}{
			{"a", 1},
			{"b", 2},
		}
		for _, pair := range data {
			if !yield(pair.key, pair.val) {
				return
			}
		}
	}
	k, v, ok := iters.Last2(seq2)
	fmt.Println(k, v, ok)
	// Output:
	// b 2 true
}

func ExampleLastFunc2_match() {
	seq2 := func(yield func(string, int) bool) {
		data := []struct {
			key string
			val int
		}{
			{"a", 1},
			{"b", 2},
			{"c", 3},
		}
		for _, pair := range data {
			if !yield(pair.key, pair.val) {
				return
			}
		}
	}
	k, v, ok := iters.LastFunc2(seq2, func(_ string, v int) bool { return v >= 2 })
	fmt.Println(k, v, ok)
	// Output:
	// c 3 true
}

type lastTableTest[T comparable] struct {
	name     string
	input    []T
	expected T
	ok       bool
}

func (test lastTableTest[T]) Run(t *testing.T) {
	runLastTableTest(t, test)
}

func runLastTableTest[T comparable](t *testing.T, test lastTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got, ok := iters.Last(slices.Values(test.input))
		if ok != test.ok || (ok && got != test.expected) {
			t.Fatalf("Last: expected (%v,%v), got (%v,%v)", test.expected, test.ok, got, ok)
		}
	})
}

type lastFuncTableTest[T comparable] struct {
	name     string
	input    []T
	pred     func(T) bool
	expected T
	ok       bool
}

func (test lastFuncTableTest[T]) Run(t *testing.T) {
	runLastFuncTableTest(t, test)
}

func runLastFuncTableTest[T comparable](t *testing.T, test lastFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got, ok := iters.LastFunc(slices.Values(test.input), test.pred)
		if ok != test.ok || (ok && got != test.expected) {
			t.Fatalf("LastFunc: expected (%v,%v), got (%v,%v)", test.expected, test.ok, got, ok)
		}
	})
}

type last2TableTest[K comparable, V comparable] struct {
	name  string
	input []struct {
		key K
		val V
	}
	expectedK K
	expectedV V
	ok        bool
}

func (test last2TableTest[K, V]) Run(t *testing.T) {
	runLast2TableTest(t, test)
}

func runLast2TableTest[K comparable, V comparable](t *testing.T, test last2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		seq2 := func(yield func(K, V) bool) {
			for _, pair := range test.input {
				if !yield(pair.key, pair.val) {
					return
				}
			}
		}
		k, v, ok := iters.Last2(seq2)
		if ok != test.ok || (ok && (k != test.expectedK || v != test.expectedV)) {
			t.Fatalf("Last2: expected (%v,%v,%v), got (%v,%v,%v)", test.expectedK, test.expectedV, test.ok, k, v, ok)
		}
	})
}

type lastFunc2TableTest[K comparable, V comparable] struct {
	name  string
	input []struct {
		key K
		val V
	}
	pred      func(K, V) bool
	expectedK K
	expectedV V
	ok        bool
}

func (test lastFunc2TableTest[K, V]) Run(t *testing.T) {
	runLastFunc2TableTest(t, test)
}

func runLastFunc2TableTest[K comparable, V comparable](t *testing.T, test lastFunc2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		seq2 := func(yield func(K, V) bool) {
			for _, pair := range test.input {
				if !yield(pair.key, pair.val) {
					return
				}
			}
		}
		k, v, ok := iters.LastFunc2(seq2, test.pred)
		if ok != test.ok || (ok && (k != test.expectedK || v != test.expectedV)) {
			t.Fatalf("LastFunc2: expected (%v,%v,%v), got (%v,%v,%v)", test.expectedK, test.expectedV, test.ok, k, v, ok)
		}
	})
}

func TestLast(t *testing.T) {
	tests := []runnableTest{
		lastTableTest[int]{name: "numbers", input: []int{1, 2, 3}, expected: 3, ok: true},
		lastTableTest[int]{name: "empty", input: nil, ok: false},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestLastFunc(t *testing.T) {
	tests := []runnableTest{
		lastFuncTableTest[int]{
			name:     "greater than",
			input:    []int{1, 4, 2},
			pred:     func(v int) bool { return v > 2 },
			expected: 4,
			ok:       true,
		},
		lastFuncTableTest[int]{
			name:  "no match",
			input: []int{1, 2},
			pred:  func(int) bool { return false },
			ok:    false,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestLast2(t *testing.T) {
	tests := []runnableTest{
		last2TableTest[string, int]{
			name: "pairs",
			input: []struct {
				key string
				val int
			}{{"a", 1}, {"b", 2}},
			expectedK: "b",
			expectedV: 2,
			ok:        true,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestLastFunc2(t *testing.T) {
	tests := []runnableTest{
		lastFunc2TableTest[string, int]{
			name: "match last",
			input: []struct {
				key string
				val int
			}{{"a", 1}, {"b", 2}, {"c", 3}},
			pred:      func(_ string, v int) bool { return v >= 2 },
			expectedK: "c",
			expectedV: 3,
			ok:        true,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
