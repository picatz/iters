package iters_test

import (
	"testing"

	"github.com/picatz/iters"
)

func TestPredicateAlias(t *testing.T) {
	var p iters.Predicate[int] = func(v int) bool { return v%2 == 0 }
	if !p(4) || p(3) {
		t.Fatalf("predicate alias not behaving as expected")
	}
}

func TestPredicate2Alias(t *testing.T) {
	var p2 iters.Predicate2[string, int] = func(k string, v int) bool {
		return k == "target" && v > 0
	}
	if !p2("target", 1) || p2("other", 1) {
		t.Fatalf("predicate2 alias not behaving as expected")
	}
}
