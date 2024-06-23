package slices

import (
	tws "github.com/twharmon/slices"
)

// Find finds an item in the given slice that satisfies the given
// test function.
func Find[T any](s []T, test func(item T) bool) T {
	return tws.Find(s, test)
}

// IndexOf finds the index of the first item in the given slice that
// satisfies the given test function.
// func IndexOf[T tws.Ordered](s []T, item T) int {
// 	return tws.IndexOf[T](s, item)
// }

// IndexOfFunc finds the index of the first item in the given slice
// that satisfies the given test function.
func IndexOf[T any](s []T, test func(item T) bool) int {
	return tws.IndexOfFunc(s, test)
}
