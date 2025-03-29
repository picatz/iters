package iters_test

import (
	"testing"
)

// runnableTest is an interface that defines a common method
// for running tests so that we can use different types of tests
// in a simple table-driven test pattern using type parameters.
type runnableTest interface {
	Run(t *testing.T)
}
