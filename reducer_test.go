package iters_test

import (
	"testing"

	"github.com/picatz/iters"
)

func TestReducerAlias(t *testing.T) {
	var sum iters.Reducer[int, int] = func(acc, v int) int { return acc + v }
	if got := sum(2, 3); got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
}

func TestReducer2Alias(t *testing.T) {
	var combine iters.Reducer2[int, string, int] = func(acc int, _ string, v int) int {
		return acc + v
	}
	if got := combine(1, "a", 2); got != 3 {
		t.Fatalf("expected 3, got %d", got)
	}
}
