package iters_test

import (
	"errors"
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleCollectErr_paginated() {
	type page struct {
		items []int
		err   error
	}
	pages := []page{
		{items: []int{1, 2}},
		{items: []int{3}},
	}

	seq := func(yield func(int, error) bool) {
		for _, p := range pages {
			for _, item := range p.items {
				if !yield(item, nil) {
					return
				}
			}
		}
	}

	values, err := iters.CollectErr(seq)
	fmt.Println(values, err)
	// Output:
	// [1 2 3] <nil>
}

func ExampleWalkErr() {
	seq := func(yield func(string, error) bool) {
		for _, word := range []string{"hello", "iter"} {
			if !yield(word, nil) {
				return
			}
		}
	}

	_ = iters.WalkErr(seq, func(s string) bool {
		fmt.Println(s)
		return true
	})
	// Output:
	// hello
	// iter
}

func ExampleUntilErr() {
	seq := func(yield func(int, error) bool) {
		if !yield(1, nil) {
			return
		}
		if !yield(0, errors.New("boom")) {
			return
		}
		if !yield(2, nil) {
			return
		}
	}

	fmt.Println(slices.Collect(iters.UntilErr(seq)))
	// Output:
	// [1]
}

func TestCollectErrStopsOnError(t *testing.T) {
	expectedErr := errors.New("boom")
	seq := func(yield func(int, error) bool) {
		if !yield(1, nil) {
			return
		}
		if !yield(0, expectedErr) {
			return
		}
		if !yield(2, nil) {
			return
		}
	}

	values, err := iters.CollectErr(seq)
	if err != expectedErr {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
	if want := []int{1}; !slices.Equal(values, want) {
		t.Fatalf("expected %v, got %v", want, values)
	}
}

func TestWalkErrStopsWhenFnReturnsFalse(t *testing.T) {
	seq := func(yield func(int, error) bool) {
		for i := 0; i < 3; i++ {
			if !yield(i, nil) {
				return
			}
		}
	}

	var seen []int
	err := iters.WalkErr(seq, func(v int) bool {
		seen = append(seen, v)
		return len(seen) < 2
	})
	if err != nil {
		t.Fatalf("unexpected err %v", err)
	}
	if want := []int{0, 1}; !slices.Equal(seen, want) {
		t.Fatalf("expected %v, got %v", want, seen)
	}
}

func TestWalkErrPropagatesError(t *testing.T) {
	expectedErr := errors.New("boom")
	seq := func(yield func(int, error) bool) {
		if !yield(1, nil) {
			return
		}
		if !yield(0, expectedErr) {
			return
		}
	}

	err := iters.WalkErr(seq, func(int) bool { return true })
	if err != expectedErr {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

func TestUntilErrStopsBeforeError(t *testing.T) {
	seq := func(yield func(int, error) bool) {
		if !yield(1, nil) {
			return
		}
		if !yield(0, errors.New("boom")) {
			return
		}
		if !yield(2, nil) {
			return
		}
	}

	got := slices.Collect(iters.UntilErr(seq))
	if want := []int{1}; !slices.Equal(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
