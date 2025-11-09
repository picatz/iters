package iters_test

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleMax_numbers() {
	max, ok := iters.Max(slices.Values([]int{3, 1, 4}))
	fmt.Println(max, ok)
	// Output:
	// 4 true
}

func ExampleMaxFunc_custom() {
	type Item struct {
		Name  string
		Score int
	}
	items := []Item{{"a", 1}, {"b", 3}}
	max, ok := iters.MaxFunc(
		slices.Values(items),
		func(a, b Item) bool { return a.Score < b.Score },
	)
	fmt.Println(max, ok)
	// Output:
	// {b 3} true
}

type maxTableTest[T cmp.Ordered] struct {
	name     string
	input    []T
	expected T
	ok       bool
}

func (test maxTableTest[T]) Run(t *testing.T) {
	runMaxTableTest(t, test)
}

func runMaxTableTest[T cmp.Ordered](t *testing.T, test maxTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got, ok := iters.Max(slices.Values(test.input))
		if ok != test.ok || (ok && got != test.expected) {
			t.Fatalf("Max: expected (%v,%v), got (%v,%v)", test.expected, test.ok, got, ok)
		}
	})
}

type maxFuncTableTest[T comparable] struct {
	name     string
	input    []T
	less     func(a, b T) bool
	expected T
	ok       bool
}

func (test maxFuncTableTest[T]) Run(t *testing.T) {
	runMaxFuncTableTest(t, test)
}

func runMaxFuncTableTest[T comparable](t *testing.T, test maxFuncTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		got, ok := iters.MaxFunc(slices.Values(test.input), test.less)
		if ok != test.ok {
			t.Fatalf("MaxFunc: expected ok=%v got %v", test.ok, ok)
		}
		if ok && got != test.expected {
			t.Fatalf("MaxFunc: expected %v, got %v", test.expected, got)
		}
	})
}

func TestMax(t *testing.T) {
	tests := []runnableTest{
		maxTableTest[int]{name: "numbers", input: []int{1, 3, 2}, expected: 3, ok: true},
		maxTableTest[int]{name: "empty", input: nil, ok: false},
	}
	for _, test := range tests {
		test.Run(t)
	}
}

func TestMaxFunc(t *testing.T) {
	tests := []runnableTest{
		maxFuncTableTest[int]{name: "custom compare", input: []int{1, 3, 2}, less: func(a, b int) bool { return a < b }, expected: 3, ok: true},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
