package iters_test

import (
	"fmt"
	"maps"
	"slices"

	"github.com/picatz/iters"
)

func ExampleConcat_numbers() {
	seq1 := []int{1, 2, 3}
	seq2 := []int{4, 5, 6}

	concatenated := iters.Concat(
		slices.Values(seq1),
		slices.Values(seq2),
	)

	result := slices.Collect(concatenated)

	fmt.Println(result)
	// Output:
	// [1 2 3 4 5 6]
}

func ExampleConcat_keyValuePairs() {
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	map2 := map[string]int{
		"three": 3,
		"four":  4,
	}

	concatenated := iters.Concat2(
		maps.All(map1),
		maps.All(map2),
	)

	result := maps.Collect(concatenated)

	fmt.Println(result["one"], result["two"], result["three"], result["four"])
	// Output:
	// 1 2 3 4
}
