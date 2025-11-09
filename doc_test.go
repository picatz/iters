package iters_test

import (
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func TestPackageSmoke(t *testing.T) {
	seq := iters.Limit(iters.Repeat(1), 3)
	if got := slices.Collect(seq); len(got) != 3 {
		t.Fatalf("unexpected result %v", got)
	}
}
