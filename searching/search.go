// Package searching contains small, readable implementations for learning.
// For application code, prefer the Go standard library's slices package.
package searching

// LinearSearch returns the first index of target, or -1 if it is absent.
// The values do not need to be sorted, and the slice is not changed.
func LinearSearch(values []int, target int) int {
	for i, value := range values {
		if value == target {
			return i
		}
	}
	return -1
}

// BinarySearch returns the first index of target, or -1 if it is absent.
// Values must already be sorted in ascending order. The slice is not changed.
func BinarySearch(values []int, target int) int {
	low, high := 0, len(values)
	for low < high {
		middle := low + (high-low)/2
		if values[middle] < target {
			low = middle + 1
		} else {
			high = middle
		}
	}
	if low < len(values) && values[low] == target {
		return low
	}
	return -1
}
