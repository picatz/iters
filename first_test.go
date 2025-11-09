package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleFirst_value() {
	fmt.Println(iters.First(slices.Values([]int{4, 5, 6})))
	// Output:
	// 4 true
}

func ExampleFirstFunc_even() {
	firstEven, ok := iters.FirstFunc(
		slices.Values([]int{1, 3, 4, 5}),
		func(v int) bool { return v%2 == 0 },
	)
	fmt.Println(firstEven, ok)
	// Output:
	// 4 true
}

func ExampleFirst2_pair() {
	seq2 := func(yield func(string, int) bool) {
		if !yield("a", 1) {
			return
		}
	}
	k, v, ok := iters.First2(seq2)
	fmt.Println(k, v, ok)
	// Output:
	// a 1 true
}

func ExampleFirstFunc2_match() {
	seq2 := func(yield func(string, int) bool) {
		pairs := []struct {
			key string
			val int
		}{
			{"a", 1},
			{"b", 3},
			{"c", 4},
		}
		for _, pair := range pairs {
			if !yield(pair.key, pair.val) {
				return
			}
		}
	}
	k, v, ok := iters.FirstFunc2(seq2, func(_ string, val int) bool { return val%2 == 0 })
	fmt.Println(k, v, ok)
	// Output:
	// c 4 true
}

type firstTableTest[T comparable] struct {
	name     string
	input    []T
	expected T
	ok       bool
}

func (test firstTableTest[T]) Run(t *testing.T) {
	runFirstTableTest(t, test)
}

func runFirstTableTest[T comparable](t *testing.T, test firstTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got, ok := iters.First(slices.Values(test.input))
		if ok != test.ok || (ok && got != test.expected) {
			t.Fatalf("First: expected (%v, %v), got (%v, %v)", test.expected, test.ok, got, ok)
		}
	})
}

type firstFuncTableTest[T comparable] struct {
	name     string
	input    []T
	pred     func(T) bool
	expected T
	ok       bool
}

func (test firstFuncTableTest[T]) Run(t *testing.T) {
	runFirstFuncTableTest(t, test)
}

func runFirstFuncTableTest[T comparable](t *testing.T, test firstFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got, ok := iters.FirstFunc(slices.Values(test.input), test.pred)
		if ok != test.ok || (ok && got != test.expected) {
			t.Fatalf("FirstFunc: expected (%v, %v), got (%v, %v)", test.expected, test.ok, got, ok)
		}
	})
}

type first2TableTest[K comparable, V comparable] struct {
	name  string
	input []struct {
		key K
		val V
	}
	expectedK K
	expectedV V
	ok        bool
}

func (test first2TableTest[K, V]) Run(t *testing.T) {
	runFirst2TableTest(t, test)
}

func runFirst2TableTest[K comparable, V comparable](t *testing.T, test first2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		seq2 := func(yield func(K, V) bool) {
			for _, pair := range test.input {
				if !yield(pair.key, pair.val) {
					return
				}
			}
		}
		k, v, ok := iters.First2(seq2)
		if ok != test.ok || (ok && (k != test.expectedK || v != test.expectedV)) {
			t.Fatalf("First2: expected (%v,%v,%v), got (%v,%v,%v)", test.expectedK, test.expectedV, test.ok, k, v, ok)
		}
	})
}

type firstFunc2TableTest[K comparable, V comparable] struct {
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

func (test firstFunc2TableTest[K, V]) Run(t *testing.T) {
	runFirstFunc2TableTest(t, test)
}

func runFirstFunc2TableTest[K comparable, V comparable](t *testing.T, test firstFunc2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		seq2 := func(yield func(K, V) bool) {
			for _, pair := range test.input {
				if !yield(pair.key, pair.val) {
					return
				}
			}
		}
		k, v, ok := iters.FirstFunc2(seq2, test.pred)
		if ok != test.ok || (ok && (k != test.expectedK || v != test.expectedV)) {
			t.Fatalf("FirstFunc2: expected (%v,%v,%v), got (%v,%v,%v)", test.expectedK, test.expectedV, test.ok, k, v, ok)
		}
	})
}

func TestFirst(t *testing.T) {
	tests := []runnableTest{
		firstTableTest[int]{name: "numbers", input: []int{3, 4}, expected: 3, ok: true},
		firstTableTest[int]{name: "empty", input: nil, expected: 0, ok: false},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestFirstFunc(t *testing.T) {
	tests := []runnableTest{
		firstFuncTableTest[int]{name: "find even", input: []int{1, 3, 4}, pred: func(v int) bool { return v%2 == 0 }, expected: 4, ok: true},
		firstFuncTableTest[int]{name: "no match", input: []int{1}, pred: func(int) bool { return false }, ok: false},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestFirst2(t *testing.T) {
	tests := []runnableTest{
		first2TableTest[string, int]{
			name: "pairs",
			input: []struct {
				key string
				val int
			}{{"a", 1}, {"b", 2}},
			expectedK: "a",
			expectedV: 1,
			ok:        true,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestFirstFunc2(t *testing.T) {
	tests := []runnableTest{
		firstFunc2TableTest[string, int]{
			name: "match value",
			input: []struct {
				key string
				val int
			}{{"a", 1}, {"b", 2}},
			pred:      func(_ string, v int) bool { return v > 1 },
			expectedK: "b",
			expectedV: 2,
			ok:        true,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
