package iters_test

import (
	"testing"

	"github.com/picatz/iters"
)

func TestMapperAlias(t *testing.T) {
	var mapper iters.Mapper[int, int] = func(v int) int { return v + 1 }
	if mapper(2) != 3 {
		t.Fatalf("expected mapper to add one")
	}
}

func TestMapper2Alias(t *testing.T) {
	var mapper iters.Mapper2[string, int, string, int] = func(k string, v int) (string, int) {
		return k + "_x", v + 1
	}
	k, v := mapper("a", 1)
	if k != "a_x" || v != 2 {
		t.Fatalf("unexpected result %q %d", k, v)
	}
}
