package iters_test

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/picatz/iters"
)

func ExampleCompare_numbers() {
	a := slices.Values([]int{1, 2, 3})
	b := slices.Values([]int{1, 2, 4})

	fmt.Println(iters.Compare(a, b))
	// Output:
	// -1
}

func ExampleCompareFunc_ignoreCase() {
	a := slices.Values([]string{"Go", "iters"})
	b := slices.Values([]string{"go", "Iters"})

	fmt.Println(
		iters.CompareFunc(
			a,
			b,
			func(x, y string) int {
				xLower := strings.ToLower(x)
				yLower := strings.ToLower(y)
				switch {
				case xLower < yLower:
					return -1
				case xLower > yLower:
					return 1
				default:
					return 0
				}
			},
		),
	)
	// Output:
	// 0
}

type compareTableTest[T cmp.Ordered] struct {
	name     string
	seq1     []T
	seq2     []T
	expected int
}

func (test compareTableTest[T]) Run(t *testing.T) {
	runCompareTableTest(t, test)
}

func runCompareTableTest[T cmp.Ordered](t *testing.T, test compareTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.Compare(slices.Values(test.seq1), slices.Values(test.seq2))
		if got != test.expected {
			t.Fatalf("Compare: expected %d, got %d", test.expected, got)
		}
	})
}

type compareFuncTableTest[T1 any, T2 any] struct {
	name     string
	seq1     []T1
	seq2     []T2
	cmp      func(T1, T2) int
	expected int
}

func (test compareFuncTableTest[T1, T2]) Run(t *testing.T) {
	runCompareFuncTableTest(t, test)
}

func runCompareFuncTableTest[T1 any, T2 any](t *testing.T, test compareFuncTableTest[T1, T2]) {
	t.Run(test.name, func(t *testing.T) {
		got := iters.CompareFunc(slices.Values(test.seq1), slices.Values(test.seq2), test.cmp)
		if got != test.expected {
			t.Fatalf("CompareFunc: expected %d, got %d", test.expected, got)
		}
	})
}

func TestCompare(t *testing.T) {
	tests := []runnableTest{
		compareTableTest[int]{name: "equal sequences", seq1: []int{1, 2}, seq2: []int{1, 2}, expected: 0},
		compareTableTest[int]{name: "shorter sequence", seq1: []int{1}, seq2: []int{1, 2}, expected: -1},
		compareTableTest[int]{name: "first greater", seq1: []int{2}, seq2: []int{1}, expected: 1},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestCompareFunc(t *testing.T) {
	tests := []runnableTest{
		compareFuncTableTest[string, string]{
			name: "case-insensitive equal",
			seq1: []string{"Go"},
			seq2: []string{"go"},
			cmp: func(a, b string) int {
				return strings.Compare(strings.ToLower(a), strings.ToLower(b))
			},
			expected: 0,
		},
		compareFuncTableTest[string, string]{
			name: "custom order",
			seq1: []string{"a", "c"},
			seq2: []string{"a", "b"},
			cmp: func(a, b string) int {
				switch {
				case a < b:
					return -1
				case a > b:
					return 1
				default:
					return 0
				}
			},
			expected: 1,
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
