package iters_test

import (
	"fmt"
	"iter"
	"math"
	"slices"
	"strings"
	"testing"

	"github.com/picatz/iters"
)

func ExampleEqual_numbers() {
	a := slices.Values([]int{1, 2, 3})
	b := slices.Values([]int{1, 2, 3})

	fmt.Println(iters.Equal(a, b))
	// Output:
	// true
}

func ExampleEqual2_pairs() {
	seq := func(yield func(string, int) bool) {
		pairs := []struct {
			key string
			val int
		}{
			{"a", 1},
			{"b", 2},
		}
		for _, pair := range pairs {
			if !yield(pair.key, pair.val) {
				return
			}
		}
	}

	fmt.Println(iters.Equal2(seq, seq))
	// Output:
	// true
}

func ExampleEqualFunc_numbers() {
	a := slices.Values([]float64{1, 2, math.NaN()})
	b := slices.Values([]float64{1, 2, math.NaN()})

	fmt.Println(
		iters.EqualFunc(
			a,
			b,
			func(x, y float64) bool {
				if math.IsNaN(x) && math.IsNaN(y) {
					return true
				}
				return x == y
			},
		),
	)
	// Output:
	// true
}

func ExampleEqualFunc2_pairs() {
	seq := func(yield func(string, int) bool) {
		pairs := []struct {
			key string
			val int
		}{
			{"a", 1},
			{"b", 2},
		}
		for _, pair := range pairs {
			if !yield(pair.key, pair.val) {
				return
			}
		}
	}

	fmt.Println(
		iters.EqualFunc2(
			seq,
			seq,
			func(k1 string, v1 int, k2 string, v2 int) bool {
				return k1 == k2 && v1 == v2
			},
		),
	)
	// Output:
	// true
}

type equalTableTest[T comparable] struct {
	name     string
	seq1     []T
	seq2     []T
	expected bool
}

func (test equalTableTest[T]) Run(t *testing.T) {
	runEqualTableTest(t, test)
}

func runEqualTableTest[T comparable](t *testing.T, test equalTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		s1 := slices.Values(test.seq1)
		s2 := slices.Values(test.seq2)
		got := iters.Equal(s1, s2)
		if got != test.expected {
			t.Fatalf("Equal: expected %v, got %v", test.expected, got)
		}
	})
}

type equal2TableTest[K comparable, V comparable] struct {
	name string
	seq1 []struct {
		key K
		val V
	}
	seq2 []struct {
		key K
		val V
	}
	expected bool
}

func (test equal2TableTest[K, V]) Run(t *testing.T) {
	runEqual2TableTest(t, test)
}

func runEqual2TableTest[K comparable, V comparable](t *testing.T, test equal2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		toSeq := func(items []struct {
			key K
			val V
		}) iter.Seq2[K, V] {
			return func(yield func(K, V) bool) {
				for _, pair := range items {
					if !yield(pair.key, pair.val) {
						return
					}
				}
			}
		}

		got := iters.Equal2(toSeq(test.seq1), toSeq(test.seq2))
		if got != test.expected {
			t.Fatalf("Equal2: expected %v, got %v", test.expected, got)
		}
	})
}

type equalFuncTableTest[T any] struct {
	name     string
	seq1     []T
	seq2     []T
	equal    func(T, T) bool
	expected bool
}

func (test equalFuncTableTest[T]) Run(t *testing.T) {
	runEqualFuncTableTest(t, test)
}

func runEqualFuncTableTest[T any](t *testing.T, test equalFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		s1 := slices.Values(test.seq1)
		s2 := slices.Values(test.seq2)
		got := iters.EqualFunc(s1, s2, test.equal)
		if got != test.expected {
			t.Fatalf("EqualFunc: expected %v, got %v", test.expected, got)
		}
	})
}

type equalFunc2TableTest[K any, V any] struct {
	name string
	seq1 []struct {
		key K
		val V
	}
	seq2 []struct {
		key K
		val V
	}
	equal    func(K, V, K, V) bool
	expected bool
}

func (test equalFunc2TableTest[K, V]) Run(t *testing.T) {
	runEqualFunc2TableTest(t, test)
}

func runEqualFunc2TableTest[K any, V any](t *testing.T, test equalFunc2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		toSeq := func(items []struct {
			key K
			val V
		}) iter.Seq2[K, V] {
			return func(yield func(K, V) bool) {
				for _, pair := range items {
					if !yield(pair.key, pair.val) {
						return
					}
				}
			}
		}

		got := iters.EqualFunc2(toSeq(test.seq1), toSeq(test.seq2), test.equal)
		if got != test.expected {
			t.Fatalf("EqualFunc2: expected %v, got %v", test.expected, got)
		}
	})
}

func TestEqual(t *testing.T) {
	tests := []runnableTest{
		equalTableTest[int]{
			name:     "identical",
			seq1:     []int{1, 2},
			seq2:     []int{1, 2},
			expected: true,
		},
		equalTableTest[int]{
			name:     "different length",
			seq1:     []int{1, 2},
			seq2:     []int{1},
			expected: false,
		},
		equalTableTest[int]{
			name:     "different values",
			seq1:     []int{1, 2},
			seq2:     []int{1, 3},
			expected: false,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestEqual2(t *testing.T) {
	tests := []runnableTest{
		equal2TableTest[string, int]{
			name: "identical",
			seq1: []struct {
				key string
				val int
			}{
				{"a", 1},
				{"b", 2},
			},
			seq2: []struct {
				key string
				val int
			}{
				{"a", 1},
				{"b", 2},
			},
			expected: true,
		},
		equal2TableTest[string, int]{
			name: "different values",
			seq1: []struct {
				key string
				val int
			}{
				{"a", 1},
			},
			seq2: []struct {
				key string
				val int
			}{
				{"a", 2},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestEqualFunc(t *testing.T) {
	tests := []runnableTest{
		equalFuncTableTest[float64]{
			name: "nan aware",
			seq1: []float64{1, math.NaN()},
			seq2: []float64{1, math.NaN()},
			equal: func(a, b float64) bool {
				if math.IsNaN(a) && math.IsNaN(b) {
					return true
				}
				return a == b
			},
			expected: true,
		},
		equalFuncTableTest[float64]{
			name:     "different length",
			seq1:     []float64{1},
			seq2:     []float64{1, 2},
			equal:    func(a, b float64) bool { return a == b },
			expected: false,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestEqualFunc2(t *testing.T) {
	tests := []runnableTest{
		equalFunc2TableTest[string, string]{
			name: "case insensitive",
			seq1: []struct {
				key string
				val string
			}{
				{"A", "One"},
			},
			seq2: []struct {
				key string
				val string
			}{
				{"a", "one"},
			},
			equal: func(k1, v1, k2, v2 string) bool {
				return strings.EqualFold(k1, k2) && strings.EqualFold(v1, v2)
			},
			expected: true,
		},
		equalFunc2TableTest[string, string]{
			name: "length mismatch",
			seq1: []struct {
				key string
				val string
			}{
				{"a", "one"},
			},
			seq2: []struct {
				key string
				val string
			}{
				{"a", "one"},
				{"b", "two"},
			},
			equal:    func(k1, v1, k2, v2 string) bool { return k1 == k2 && v1 == v2 },
			expected: false,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
