// Package recursion shows how to solve a problem using a smaller version of it.
package recursion

// Sum returns the total of values, or zero for an empty slice.
// It takes O(n) time and O(n) call-stack space for n values.
func Sum(values []int) int {
	if len(values) == 0 {
		return 0
	}
	return values[0] + Sum(values[1:])
}
