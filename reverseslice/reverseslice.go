// Package reverseslice reverses the order of elements in a slice.
package reverseslice

import "slices"

// Reverse returns a new slice holding the elements of s in reverse order.
// The input slice is not modified.
func ReverseManual[T any](s []T) []T {
	n := len(s)
	reversed := make([]T, 0, n)
	for i := n - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}
	return reversed
}

func ReverseUsingSlice[T any](s []T) []T {
	n := len(s)
	if n == 0 {
		return make([]T, 0)
	}
	reversed := slices.Clone(s)
	slices.Reverse(reversed)
	return reversed
}
