package slices

import (
	tws "github.com/twharmon/slices"
)

// Find finds an item in the given slice that satisfies the given
// test function.
func Find[T any](s []T, test func(item T) bool) T {
	return tws.Find[T](s, test)
}
