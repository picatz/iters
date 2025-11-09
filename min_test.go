package iters_test

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleMin_numbers() {
	min, ok := iters.Min(slices.Values([]int{3, 1, 4}))
	fmt.Println(min, ok)
	// Output:
	// 1 true
}

func ExampleMinFunc_custom() {
	type Item struct {
		Name  string
		Score int
	}
	items := []Item{{"a", 1}, {"b", 3}}
	min, ok := iters.MinFunc(
		slices.Values(items),
		func(a, b Item) bool { return a.Score < b.Score },
	)
	fmt.Println(min, ok)
	// Output:
	// {a 1} true
}

type minTableTest[T cmp.Ordered] struct {
	name     string
	input    []T
	expected T
	ok       bool
}

func (test minTableTest[T]) Run(t *testing.T) {
	runMinTableTest(t, test)
}

func runMinTableTest[T cmp.Ordered](t *testing.T, test minTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got, ok := iters.Min(slices.Values(test.input))
		if ok != test.ok || (ok && got != test.expected) {
			t.Fatalf("Min: expected (%v,%v), got (%v,%v)", test.expected, test.ok, got, ok)
		}
	})
}

type minFuncTableTest[T comparable] struct {
	name     string
	input    []T
	less     func(a, b T) bool
	expected T
	ok       bool
}

func (test minFuncTableTest[T]) Run(t *testing.T) {
	runMinFuncTableTest(t, test)
}

func runMinFuncTableTest[T comparable](t *testing.T, test minFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got, ok := iters.MinFunc(slices.Values(test.input), test.less)
		if ok != test.ok {
			t.Fatalf("MinFunc: expected ok=%v got %v", test.ok, ok)
		}
		if ok && got != test.expected {
			t.Fatalf("MinFunc: expected %v, got %v", test.expected, got)
		}
	})
}

func TestMin(t *testing.T) {
	tests := []runnableTest{
		minTableTest[int]{name: "numbers", input: []int{3, 1, 2}, expected: 1, ok: true},
		minTableTest[int]{name: "empty", input: nil, ok: false},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestMinFunc(t *testing.T) {
	tests := []runnableTest{
		minFuncTableTest[int]{name: "custom compare", input: []int{3, 1, 2}, less: func(a, b int) bool { return a < b }, expected: 1, ok: true},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
