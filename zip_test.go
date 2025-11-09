package iters_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/picatz/iters"
)

func ExampleZip_sequences() {
	seq := iters.Zip(
		slices.Values([]int{1, 2}),
		slices.Values([]string{"a", "b"}),
	)
	for n, s := range seq {
		fmt.Println(n, s)
	}
	// Output:
	// 1 a
	// 2 b
}

type zipTableTest[T comparable, U comparable] struct {
	name       string
	left       []T
	right      []U
	expectedTv []T
	expectedUv []U
}

func (test zipTableTest[T, U]) Run(t *testing.T) {
	runZipTableTest(t, test)
}

func runZipTableTest[T comparable, U comparable](t *testing.T, test zipTableTest[T, U]) {
	t.Run(test.name, func(t *testing.T) {
		var gotT []T
		var gotU []U
		for tVal, uVal := range iters.Zip(slices.Values(test.left), slices.Values(test.right)) {
			gotT = append(gotT, tVal)
			gotU = append(gotU, uVal)
		}
		if !slices.Equal(gotT, test.expectedTv) || !slices.Equal(gotU, test.expectedUv) {
			t.Fatalf("Zip: expected %v/%v, got %v/%v", test.expectedTv, test.expectedUv, gotT, gotU)
		}
	})
}

func TestZip(t *testing.T) {
	tests := []runnableTest{
		zipTableTest[int, string]{
			name:       "equal lengths",
			left:       []int{1, 2},
			right:      []string{"a", "b"},
			expectedTv: []int{1, 2},
			expectedUv: []string{"a", "b"},
		},
		zipTableTest[int, string]{
			name:       "truncates shorter sequence",
			left:       []int{1, 2},
			right:      []string{"a"},
			expectedTv: []int{1},
			expectedUv: []string{"a"},
		},
	}
	for _, test := range tests {
		test.Run(t)
	}
}
