package iters_test

import (
	"cmp"
	"context"
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleSplit_map() {
	input := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	output := make(map[int]string)

	maps.Insert(
		output,
		iters.Zip(
			iters.Split(
				context.Background(),
				maps.All(input),
			),
		),
	)
	fmt.Println(maps.Equal(output, input))
	// Output:
	// true
}

type splitTableTest[K cmp.Ordered, V cmp.Ordered] struct {
	name         string
	input        map[K]V
	expectedKeys []K
	expectedVals []V
}

func (test splitTableTest[K, V]) Run(t *testing.T) {
	runSplitTableTest(t, test)
}

func runSplitTableTest[K cmp.Ordered, V cmp.Ordered](t *testing.T, test splitTableTest[K, V]) {
	t.Run(test.name, func(t *testing.T) {
		keysSeq, valuesSeq := iters.Split(context.Background(), maps.All(test.input))
		var gotKeys []K
		var gotVals []V
		for k, v := range iters.Zip(keysSeq, valuesSeq) {
			gotKeys = append(gotKeys, k)
			gotVals = append(gotVals, v)
		}
		if !equalAsMultiset(gotKeys, test.expectedKeys) {
			t.Fatalf("Split keys: expected %v, got %v", test.expectedKeys, gotKeys)
		}
		if !equalAsMultiset(gotVals, test.expectedVals) {
			t.Fatalf("Split values: expected %v, got %v", test.expectedVals, gotVals)
		}
	})
}

func equalAsMultiset[T cmp.Ordered](got, want []T) bool {
	if len(got) != len(want) {
		return false
	}
	gotCopy := slices.Clone(got)
	wantCopy := slices.Clone(want)
	slices.Sort(gotCopy)
	slices.Sort(wantCopy)
	return slices.Equal(gotCopy, wantCopy)
}

func TestSplit(t *testing.T) {
	tests := []runnableTest{
		splitTableTest[int, string]{
			name: "basic map",
			input: map[int]string{
				1: "one",
				2: "two",
			},
			expectedKeys: []int{1, 2},
			expectedVals: []string{"one", "two"},
		},
	}

	for _, test := range tests {
		test.Run(t)
	}
}
