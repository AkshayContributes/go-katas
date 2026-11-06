// Package reverseslice reverses the order of elements in a slice.
package reverseslice

// Reverse returns a new slice holding the elements of s in reverse order.
// The input slice is not modified. A nil slice reverses to a nil slice.
func Reverse[T any](s []T) []T {
	if s == nil {
		return nil
	}

	out := make([]T, len(s))
	for i, v := range s {
		out[len(s)-1-i] = v
	}

	return out
}
