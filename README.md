# iters

[![Go Reference](https://pkg.go.dev/badge/github.com/picatz/iters.svg)](https://pkg.go.dev/github.com/picatz/iters)

A collection of functions for working with [iterators] in Go.

[iterators]: https://pkg.go.dev/iter

## `iters.Map`

The `iters.Map` function allows you to transform each element in a slice using a provided function, 
returning a new iterator with the transformed values. This is useful for applying a function to each 
element in a slice without modifying the original slice.

```go
numbers := []int{1, 2, 3, 4, 5}

stringNumbers := iters.Map(
    slices.Values(numbers),
    func(n int) string {
        return fmt.Sprintf("%d", n)
    },
)

fmt.Println(slices.Values(stringNumbers))
// Output:
// ["1" "2" "3" "4" "5"]
```

```go
numbers := []int{1, 2, 3, 4, 5}

squaredNumbers := iters.Map(
    slices.Values(numbers),
    func(n int) int {
        return n * n
    },
)

fmt.Println(slices.Values(squaredNumbers))
// Output:
// [1 4 9 16 25]
```

```go
type Animal struct {
    Name string
    Legs int
}

animals := []Animal{
    {"cat", 4},
    {"dog", 4},
    {"fish", 0},
    {"bird", 2},
}

animalNames := iters.Map(
    slices.Values(animals),
    func(animal Animal) string {
        return animal.Name
    },
)

result := slices.Collect(animalNames)

fmt.Println(result)
// Output:
// [cat dog fish bird]
```

## `iters.Filter`

The `iters.Filter` function allows you to filter elements in a slice based on a provided predicate function.
It returns a new iterator containing only the elements that satisfy the predicate condition (i.e., the function
returns `true` for that element). This is useful for creating a new slice that only contains elements
that meet certain criteria, without modifying the original slice.

```go
numbers := []int{1, 2, 3, 4, 5}

evenNumbers := iters.Filter(
    slices.Values(numbers),
    func(n int) bool {
        return n%2 == 0
    },
)

fmt.Println(slices.Values(evenNumbers))
// Output:
// [2 4]
```

```go
animals := []Animal{
	{"cat", 4},
	{"dog", 4},
	{"fish", 0},
	{"bird", 2},
}

filteredAnimals := iters.Filter(
	slices.Values(animals),
	func(animal Animal) bool {
		return animal.Legs > 2
	},
)

result := slices.Collect(filteredAnimals)

fmt.Println(result)
// Output:
// [{cat 4} {dog 4}]
```

## `iters.Reduce`

The `iters.Reduce` function allows you to reduce a slice to a single value by applying a provided
function to each element in the slice. The function takes an accumulator value and the current element,
and returns a new accumulator value. This is useful for aggregating values, such as summing numbers,
concatenating strings, or performing any other kind of aggregation where you want to combine all
elements in a slice into a single result.

```go
type Animal struct {
    Name string
    Legs int
}

animals := []Animal{
    {"cat", 4},
    {"dog", 4},
    {"fish", 0},
    {"bird", 2},
}

totalLegs := iters.Reduce(
    slices.Values(animals),
    0,
    func(acc int, animal Animal) int {
        return acc + animal.Legs
    },
)

fmt.Println(totalLegs)
// Output:
// 10
```