package iters_test

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleContext_cancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	seq := iters.Context(
		ctx,
		func(yield func(int) bool) {
			for i := 1; i <= 5; i++ {
				if i == 3 {
					cancel()
				}
				if !yield(i) {
					return
				}
			}
		},
	)

	fmt.Println(slices.Collect(seq))
	// Output:
	// [1 2]
}

func ExampleContext2_cancel() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	seq2 := func(yield func(string, int) bool) {
		pairs := []struct {
			key string
			val int
		}{
			{"a", 1},
			{"b", 2},
			{"c", 3},
		}
		for _, pair := range pairs {
			if pair.key == "b" {
				cancel()
			}
			if !yield(pair.key, pair.val) {
				return
			}
		}
	}

	ctxSeq := iters.Context2(ctx, seq2)

	var keys []string
	for k := range ctxSeq {
		keys = append(keys, k)
	}

	fmt.Println(keys)
	// Output:
	// [a]
}

type contextTableTest[T comparable] struct {
	name      string
	input     []T
	cancelIdx int
	expected  []T
}

func (test contextTableTest[T]) Run(t *testing.T) {
	runContextTableTest(t, test)
}

func runContextTableTest[T comparable](t *testing.T, test contextTableTest[T]) {
	t.Run(test.name, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		seq := func(yield func(T) bool) {
			for i, v := range test.input {
				if test.cancelIdx >= 0 && i == test.cancelIdx {
					cancel()
				}
				if !yield(v) {
					return
				}
			}
		}

		got := slices.Collect(iters.Context(ctx, seq))
		if !slices.Equal(got, test.expected) {
			t.Fatalf("Context: expected %v, got %v", test.expected, got)
		}
	})
}

type context2TableTest[K comparable, V comparable] struct {
	name  string
	input []struct {
		key K
		val V
	}
	cancelIdx int
	expectedK []K
	expectedV []V
}

func (test context2TableTest[K, V]) Run(t *testing.T) {
	runContext2TableTest(t, test)
}

func runContext2TableTest[K comparable, V comparable](t *testing.T, test context2TableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		seq2 := func(yield func(K, V) bool) {
			for i, pair := range test.input {
				if test.cancelIdx >= 0 && i == test.cancelIdx {
					cancel()
				}
				if !yield(pair.key, pair.val) {
					return
				}
			}
		}

		var gotK []K
		var gotV []V
		for k, v := range iters.Context2(ctx, seq2) {
			gotK = append(gotK, k)
			gotV = append(gotV, v)
		}

		if !slices.Equal(gotK, test.expectedK) {
			t.Fatalf("Context2 keys: expected %v, got %v", test.expectedK, gotK)
		}
		if !slices.Equal(gotV, test.expectedV) {
			t.Fatalf("Context2 values: expected %v, got %v", test.expectedV, gotV)
		}
	})
}

func TestContext(t *testing.T) {
	tests := []runnableTest{
		contextTableTest[int]{
			name:      "no cancellation",
			input:     []int{1, 2, 3},
			cancelIdx: -1,
			expected:  []int{1, 2, 3},
		},
		contextTableTest[int]{
			name:      "cancel mid-way",
			input:     []int{1, 2, 3},
			cancelIdx: 1,
			expected:  []int{1},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}

func TestContext2(t *testing.T) {
	tests := []runnableTest{
		context2TableTest[string, int]{
			name: "cancel before last pair",
			input: []struct {
				key string
				val int
			}{
				{"a", 1},
				{"b", 2},
				{"c", 3},
			},
			cancelIdx: 2,
			expectedK: []string{"a", "b"},
			expectedV: []int{1, 2},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
